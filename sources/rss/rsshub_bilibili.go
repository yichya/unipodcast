package rss

import (
	"github.com/mmcdole/gofeed"
	"github.com/yichya/unipodcast/pipeline/source"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

const RsshubBilibili = "rss_rsshub_bilibili"

func ParseRsshubBilibili(client *http.Client, path string, filters []func(s *source.Source) bool) ([]*source.Source, error) {
	return ParseRss(client, path, filters, func(f *gofeed.Feed, i *gofeed.Item, s *source.Source) {
		h, err := html.Parse(strings.NewReader(i.Description))
		if err != nil {
			return
		}
		if h.FirstChild != nil {
			if h.FirstChild.FirstChild != nil {
				if h.FirstChild.FirstChild.NextSibling != nil {
					if h.FirstChild.FirstChild.NextSibling.FirstChild != nil {
						s.Description = h.FirstChild.FirstChild.NextSibling.FirstChild.Data
						if h.FirstChild.FirstChild.NextSibling.FirstChild.NextSibling != nil {
							if h.FirstChild.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling != nil {
								if h.FirstChild.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling != nil {
									if len(h.FirstChild.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.Attr) > 0 {
										s.Url = h.FirstChild.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.Attr[0].Val
									}
								}
							}
						}
					}
					if h.FirstChild.FirstChild.NextSibling.LastChild != nil {
						if len(h.FirstChild.FirstChild.NextSibling.LastChild.Attr) > 0 {
							s.PicUrl = h.FirstChild.FirstChild.NextSibling.LastChild.Attr[0].Val
						}
					}
				}
			}
		}
	})
}
