package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var f func(*html.Node, bool)
	output := []Link{}
	f = func(n *html.Node, collect bool) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					output = append(output, Link{Href: a.Val})
					collect = true
					break
				}
			}
		}
		if collect && n.Type == html.TextNode {
			output[len(output)-1].Text = output[len(output)-1].Text + strings.Trim(n.Data, "\t \n")
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, collect)
		}
	}
	f(doc, false)
	return output, nil
}
