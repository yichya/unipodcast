package filters

import (
	"github.com/yichya/unipodcast/common/constants"
	"github.com/yichya/unipodcast/pipeline/source"
	"strings"
	"time"
)

const (
	after = "after"
)

func NopFilter(i *source.Source) bool {
	return true
}

func FromString(s string) func(i *source.Source) bool {
	switch {
	case strings.HasPrefix(s, after):
		{
			ss := strings.Split(s, "=")
			t, err := time.Parse(constants.DefaultTimeFormat, ss[1])
			if err != nil {
				return NopFilter
			}
			return func(i *source.Source) bool {
				if i.PubDate != nil {
					return i.PubDate.After(t)
				}
				return false
			}
		}
	}
	return NopFilter
}

func FromStringArray(s []string) []func(i *source.Source) bool {
	var resp []func(i *source.Source) bool
	for _, x := range s {
		resp = append(resp, FromString(x))
	}
	return resp
}
