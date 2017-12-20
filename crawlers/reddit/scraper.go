package reddit

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func NewScraper() *Scraper {
	return &Scraper{5000, new(http.Client)}
}

type Scraper struct {
	minScore int
	client   *http.Client
}

func (s *Scraper) Scrape(subreddit string) ([]Thread, error) {
	allThreads := []Thread{}
	nextAfter := ""
	for {
		threads, after, err := s.getPage(subreddit, nextAfter)
		if err != nil {
			return nil, err
		}

		allThreads = append(allThreads, threads...)
		nextAfter = after

		if after == "" || len(threads) == 0 {
			break
		}
	}

	return allThreads, nil
}

func (s *Scraper) getPage(subreddit, after string) ([]Thread, string, error) {
	uri := fmt.Sprintf("https://www.reddit.com/r/%v/top/?sort=top&t=day", subreddit)
	if after != "" {
		uri += "&after=" + after
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:57.0) Gecko/20100101 Firefox/57.0")
	res, err := s.client.Do(req)
	if err != nil {
		return nil, "", err
	}

	if res.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("unexpected status code %v (%v)", res.StatusCode, res.Status)
	}

	root, err := html.Parse(res.Body)
	if err != nil {
		return nil, "", err
	}

	siteTable := findNode(root, func(node *html.Node) bool {
		return node.Data == "div" && getAttr(node, "id") == "siteTable"
	})

	threads := []Thread{}
	child := siteTable.FirstChild
	for child != nil {
		if thread, isThread := s.parseThread(child); isThread {
			thread.Subreddit = subreddit
			threads = append(threads, *thread)
		}

		child = child.NextSibling
	}

	nextLink := findNode(siteTable, func(node *html.Node) bool {
		return node.Data == "a" && getAttr(node, "rel") == "nofollow next"
	})

	var afterParam string
	if nextLink != nil {
		href := getAttr(nextLink, "href")
		uri, err := url.Parse(href)
		if err == nil {
			afterParam = uri.Query().Get("after")
		}
	}

	return threads, afterParam, nil
}

func (s *Scraper) parseThread(node *html.Node) (*Thread, bool) {
	attrs := map[string]string{}
	for _, a := range node.Attr {
		attrs[a.Key] = a.Val
	}

	if id, ok := attrs["id"]; !ok || !strings.HasPrefix(id, "thing_") {
		return nil, false
	}

	rawScore := attrs["data-score"]
	score, err := strconv.Atoi(rawScore)
	if err != nil {
		return nil, false
	}

	if score < s.minScore {
		return nil, false
	}

	url := attrs["data-url"]
	commentsURL := attrs["data-permalink"]

	titleNode := findNode(node, isTitle)

	var title string
	if titleNode != nil {
		title = titleNode.FirstChild.Data
	}

	return &Thread{
		Score:       score,
		Title:       title,
		CommentsURL: commentsURL,
		Permalink:   url,
	}, true
}

func isTitle(node *html.Node) bool {
	if node.Data != "a" {
		return false
	}

	for _, class := range strings.Split(getAttr(node, "class"), " ") {
		if class == "title" {
			return true
		}
	}

	return false
}
