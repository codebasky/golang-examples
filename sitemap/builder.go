package sitemap

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/codebasky/golang-examples/linkparser"
)

type builder struct {
	home  string
	depth int
}

func New(homePage string, depth int) *builder {
	return &builder{
		home:  homePage,
		depth: depth,
	}
}

func getHTML(link string) ([]byte, error) {
	//fmt.Printf("%s\n", link)
	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseLink(home string, link string) ([]string, error) {
	page, err := getHTML(link)
	if err != nil {
		return nil, err
	}
	href, err := linkparser.Parse(strings.NewReader(string(page)))
	if err != nil {
		//fmt.Println("parse failure")
		return nil, err
	}
	links := []string{}
	for _, ref := range href {
		link := ref.Href
		if strings.Contains(link, "mailto") {
			continue
		}
		if !strings.Contains(link, "http") {
			link = home + link
		}
		links = append(links, link)
	}
	return links, nil
}

func getHomeLinks(home string, link []string) []string {
	vlink := []string{}
	for _, lnk := range link {
		if strings.Contains(lnk, home) {
			vlink = append(vlink, lnk)
		}
	}
	return vlink
}

func buildSite(home string, level int, siteMap map[string]struct{}) error {
	links := []string{home}
	qLen := len(links)
	curr := 0
	for qLen != 0 {
		for i := 0; i < qLen; i++ {
			link := links[i]
			if _, ok := siteMap[link]; !ok {
				siteMap[link] = struct{}{}
				newLinks, err := parseLink(home, link)
				if err != nil {
					fmt.Printf("parset link failed for link: %s error: %s\n", link, err)
					continue
				}
				vlinks := getHomeLinks(home, newLinks)
				links = append(links, vlinks...)
			}
		}
		curr++
		if curr == level {
			return nil
		}
		links = links[qLen:]
		qLen = len(links)
	}
	return nil
}

func (b builder) Build() ([]string, error) {
	siteMap := make(map[string]struct{})
	err := buildSite(b.home, b.depth, siteMap)
	if err != nil {
		return nil, err
	}
	links := []string{}
	for link := range siteMap {
		links = append(links, link)
	}
	return links, nil
}
