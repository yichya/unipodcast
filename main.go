package main

import (
	"flag"
	"fmt"
	"github.com/yichya/unipodcast/pipeline"
	"github.com/yichya/unipodcast/sources/rss"
	"strings"
)

const (
	SourceType   = rss.PodcastItunes
	SourceUrl    = "http://127.0.0.1/example.xml"
	SourceProxy  = ""
	SourceFilter = "after=2022-08-26 00:00:00"
	Offset       = 0
	Limit        = 10
	Desc         = false
	Targets      = "echo_stdout=id,title,pub_date"
	TargetProxy  = "socks5://10.32.15.1:1082"
)

func main() {
	empty := ""
	sourceType := flag.String("sourceType", SourceType, "source type")
	sourceUrl := flag.String("sourceUrl", SourceUrl, "source URL")
	sourceProxy := flag.String("sourceProxy", SourceProxy, "source Proxy")
	sourceFilter := flag.String("sourceFilter", SourceFilter, "source Filter")
	targets := flag.String("targets", Targets, "targets")
	targetProxy := flag.String("targetProxy", TargetProxy, "target Proxy")
	offset := flag.Int("offset", Offset, "offset")
	limit := flag.Int("limit", Limit, "limit")
	desc := flag.Bool("desc", Desc, "desc")

	flag.Parse()
	for index, x := range []*string{sourceType, sourceUrl, targets} {
		if x == nil {
			fmt.Printf("string %d nil\n", index)
			flag.Usage()
			return
		}
		if *x == "" {
			fmt.Printf("string %d empty\n", index)
			flag.Usage()
			return
		}
	}
	for index, x := range []*int{offset, limit} {
		if x == nil {
			fmt.Printf("int %d nil\n", index)
			flag.Usage()
			return
		}
	}

	if sourceProxy == nil {
		sourceProxy = &empty
	}
	if targetProxy == nil {
		targetProxy = &empty
	}

	var filters []string
	if sourceFilter != nil {
		filters = strings.Split(strings.TrimSpace(*sourceFilter), ";")
	}
	if err := pipeline.SendItems(*sourceType, *sourceUrl, *sourceProxy, filters, *offset, *limit, *desc, *targets, *targetProxy); err != nil {
		panic(err)
	}
}
