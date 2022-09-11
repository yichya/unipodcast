package source

import (
	"time"
)

type Source struct {
	Id        string
	Title     string
	Duration  int64
	Performer string
	PubDate   *time.Time
	FileUrl   string
}

type Pagination struct {
	Offset int
	Limit  int
	Desc   bool
}

func (p *Pagination) Do(data []*Source) []*Source {
	if len(data) == 0 {
		return nil
	}
	var s = make([]*Source, len(data))
	copy(s, data)
	if p.Desc {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	if p.Offset+p.Limit < len(s) {
		return s[p.Offset : p.Offset+p.Limit]
	}
	return s[p.Offset:]
}
