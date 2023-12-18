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

func ParsePodcastItunes(client *http.Client, path string, filters []func(s *source.Source) bool) ([]*source.Source, error) {
	return ParseRss(client, path, filters, func(f *gofeed.Feed, i *gofeed.Item, s *source.Source) {
		if f.ITunesExt != nil {
			s.Performer = f.ITunesExt.Author
		}
		if i.ITunesExt != nil {
			d := strings.Split(strings.TrimSpace(i.ITunesExt.Duration), ":")
			for index, di := range d {
				if v, err := strconv.ParseInt(di, 10, 64); err != nil {
					continue
				} else {
					s.Duration += int64(math.Pow(float64(60), float64(len(d)-1-index))) * v
				}
			}
		}
		if len(i.Authors) > 0 {
			s.Performer = i.Authors[0].Name
		}
		if len(i.Enclosures) > 0 {
			s.FileUrl = i.Enclosures[0].URL
		}
	})
}
