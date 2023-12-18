package rss

import (
	"github.com/mmcdole/gofeed"
	"github.com/yichya/unipodcast/pipeline/source"
	"net/http"
	"strings"
)

func ParseRss(client *http.Client, path string, filters []func(s *source.Source) bool, extraProperties func(f *gofeed.Feed, i *gofeed.Item, s *source.Source)) ([]*source.Source, error) {
	p := gofeed.NewParser()
	p.Client = client
	feed, err := p.ParseURL(path)
	if err != nil {
		return nil, err
	}
	var resp []*source.Source
	for _, x := range feed.Items {
		sourceItem := &source.Source{
			Id:          x.GUID,
			Title:       strings.TrimSpace(x.Title),
			PubDate:     x.PublishedParsed,
			Description: x.Description,
		}
		if extraProperties != nil {
			extraProperties(feed, x, sourceItem)
		}
		for _, f := range filters {
			if f != nil {
				if f(sourceItem) {
					resp = append(resp, sourceItem)
				}
			}
		}
	}
	return resp, nil
}
