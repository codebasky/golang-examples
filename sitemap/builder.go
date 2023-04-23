package sitemap

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/codebasky/golang-examples/linkparser"
)

func getHTML(link string) ([]byte, error) {
	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s", body)
	return body, nil
}

func parseLink(link string) ([]string, error) {
	page, err := getHTML(link)
	if err != nil {
		return nil, err
	}
	href, err := linkparser.Parse(strings.NewReader(string(page)))
	links := []string{}
	for _, ref := range href {
		links = append(links, ref.Href)
	}
	return links, nil
}

func Build(home string, level int) {
	siteMap := make(map[string]struct{})
	buildSite(home, level, siteMap)
}

func buildSite(home string, level int, siteMap map[string]struct{}) error {
	links := []string{home}
	qLen := len(links)

	for qLen != 0 {
		for i := 0; i < qLen; i++ {
			link := links[i]
			if _, ok := siteMap[link]; !ok {
				siteMap[link] = struct{}{}
				newLinks, err := parseLink(link)
				if err != nil {
					return err
				}
				links = append(links, newLinks...)
			}
		}
		links = links[qLen:]
		qLen = len(links)
	}
	return nil
}
