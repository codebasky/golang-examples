package sitemap

import (
	"encoding/xml"
	"io"
)

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
}

func Encode(links []string, w io.Writer) {
	url := []Url{}
	for _, link := range links {
		url = append(url, Url{Loc: link})
	}

	data := UrlSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  url,
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	enc.Encode(data)
}
