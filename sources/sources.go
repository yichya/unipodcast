package sources

import (
	"fmt"
	"github.com/yichya/unipodcast/common/proxy_client"
	"github.com/yichya/unipodcast/pipeline/source"
	"github.com/yichya/unipodcast/sources/rss"
)

func ParseFeed(sourceType, sourceUrl, sourceProxy string, filters []func(i *source.Source) bool) ([]*source.Source, error) {
	client := proxy_client.ProxyClient(sourceProxy)
	switch sourceType {
	case rss.PodcastItunes:
		{
			return rss.ParseRssFeed(client, sourceUrl, filters)
		}
	}
	return nil, fmt.Errorf("unknown sourceType %s", sourceType)
}
