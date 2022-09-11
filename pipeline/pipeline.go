package pipeline

import (
	"github.com/yichya/unipodcast/pipeline/filters"
	"github.com/yichya/unipodcast/pipeline/source"
	"github.com/yichya/unipodcast/sources"
	"github.com/yichya/unipodcast/targets"
	"strings"
)

func SendItems(sourceType, sourceUrl, sourceProxy string, sourceFilters []string, offset, limit int, desc bool, target, targetProxy string) error {
	feed, err := sources.ParseFeed(sourceType, sourceUrl, sourceProxy, filters.FromStringArray(sourceFilters))
	if err != nil {
		panic(err)
	}
	p := source.Pagination{
		Offset: offset,
		Limit:  limit,
		Desc:   desc,
	}
	s := p.Do(feed)

	targetSplit := strings.Split(target, ";")
	for _, t := range targetSplit {
		targetItem := strings.Split(t, "=")
		targetType := strings.TrimSpace(targetItem[0])
		targetArgs := ""
		if len(targetItem) > 1 {
			targetArgs = strings.TrimSpace(targetItem[1])
		}
		for _, x := range s {
			if err = targets.SendOneFeed(x, targetProxy, targetType, targetArgs); err != nil {
				return err
			}
		}
	}

	return nil
}
