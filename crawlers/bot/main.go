package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/idwall/desafios/crawlers/reddit"
)

var token = os.Getenv("TELEGRAM_BOT_TOKEN")

func main() {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
		return
	}

	updates, err := api.GetUpdatesChan(tgbotapi.UpdateConfig{
		Timeout: 10,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	for u := range updates {
		if u.Message.Command() != "nadaprafazer" {
			continue
		}

		go func(u *tgbotapi.Update) {
			defer func() {
				if r := recover(); r != nil {
					log.Println("Panic recovered: ", r)
				}
			}()

			processNadaPraFazer(api, u)
		}(&u)
	}
}

func processNadaPraFazer(api *tgbotapi.BotAPI, u *tgbotapi.Update) {
	args := u.Message.CommandArguments()
	if args == "" {
		reply(api, u, "Please provide a list of subreddits in the following format: worldnews;videos;brazil")
		return
	}

	subreddits := strings.Split(args, ";")

	if err := reply(api, u, formatFetching(subreddits)); err != nil {
		return
	}

	scraper := reddit.NewScraper()
	for _, s := range subreddits {
		threads, err := scraper.Scrape(s)
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			break
		}

		buf := []string{}
		for _, t := range threads {
			buf = append(buf, formatThread(t))
		}

		reply(api, u, strings.Join(buf, "\n\n"))
	}

	reply(api, u, "Done!")
}

func formatThread(t reddit.Thread) string {
	return fmt.Sprintf(
		`*%v* [%v](%v) in _%v_ ([comments](https://www.reddit.com%v))`,
		t.Score, t.Title, t.Permalink, t.Subreddit, t.CommentsURL)
}

func formatFetching(subreddits []string) string {
	return fmt.Sprintf(
		"Hang on, I'm fetching the threads for you!\nSubreddits: *%v*",
		strings.Join(subreddits, "*, *"))
}

func reply(api *tgbotapi.BotAPI, update *tgbotapi.Update, text string) error {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	message.ParseMode = tgbotapi.ModeMarkdown
	message.DisableWebPagePreview = true

	if _, err := api.Send(message); err != nil {
		log.Println("Error replying: ", err)
		return err
	}

	return nil
}
