package rss

import (
	"github.com/mmcdole/gofeed"
	"github.com/yichya/unipodcast/pipeline/source"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const PodcastItunes = "rss_podcast_itunes"

func ParseRssFeed(client *http.Client, path string, filters []func(s *source.Source) bool) ([]*source.Source, error) {
	p := gofeed.NewParser()
	p.Client = client
	feed, err := p.ParseURL(path)
	if err != nil {
		return nil, err
	}
	var resp []*source.Source
	for _, x := range feed.Items {
		sourceItem := &source.Source{
			Id:      x.GUID,
			Title:   x.Title,
			PubDate: x.PublishedParsed,
		}
		if feed.ITunesExt != nil {
			sourceItem.Performer = feed.ITunesExt.Author
		}
		if x.ITunesExt != nil {
			s := strings.Split(strings.TrimSpace(x.ITunesExt.Duration), ":")
			for index, di := range s {
				if v, err := strconv.ParseInt(di, 10, 64); err != nil {
					continue
				} else {
					sourceItem.Duration += int64(math.Pow(float64(60), float64(len(s)-1-index))) * v
				}
			}
		}
		if len(x.Authors) > 0 {
			sourceItem.Performer = x.Authors[0].Name
		}
		if len(x.Enclosures) > 0 {
			sourceItem.FileUrl = x.Enclosures[0].URL
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
