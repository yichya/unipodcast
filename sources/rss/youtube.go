package rss

import (
	"github.com/mmcdole/gofeed"
	"github.com/yichya/unipodcast/pipeline/source"
	"net/http"
)

const Youtube = "rss_youtube"

func ParseYoutube(client *http.Client, path string, filters []func(s *source.Source) bool) ([]*source.Source, error) {
	return ParseRss(client, path, filters, func(f *gofeed.Feed, i *gofeed.Item, s *source.Source) {
		s.Url = i.Link
		if len(i.Authors) > 0 {
			s.Performer = i.Authors[0].Name
		}
		if media, ok1 := i.Extensions["media"]; ok1 {
			if group, ok2 := media["group"]; ok2 && len(group) > 0 {
				if title, ok3 := group[0].Children["title"]; ok3 && len(title) > 0 {
					s.Title = title[0].Value
				}
				if description, ok3 := group[0].Children["description"]; ok3 && len(description) > 0 {
					s.Description = description[0].Value
				}
				if thumbnail, ok3 := group[0].Children["thumbnail"]; ok3 && len(thumbnail) > 0 {
					if picUrl, ok4 := thumbnail[0].Attrs["url"]; ok4 {
						s.PicUrl = picUrl
					}
				}
			}
		}
	})
}
