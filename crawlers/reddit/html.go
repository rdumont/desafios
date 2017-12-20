package reddit

import "golang.org/x/net/html"

func findNode(node *html.Node, test func(node *html.Node) bool) *html.Node {
	q := []*html.Node{node}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if test(curr) {
			return curr
		}

		child := curr.FirstChild
		for child != nil {
			q = append(q, child)
			child = child.NextSibling
		}
	}

	return nil
}

func getAttr(node *html.Node, attrKey string) string {
	for _, attr := range node.Attr {
		if attr.Key == attrKey {
			return attr.Val
		}
	}

	return ""
}
