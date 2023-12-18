package wxwork_bot

import (
	"encoding/json"
	"fmt"
	"github.com/yichya/unipodcast/pipeline/source"
	"io"
	"net/http"
	"strings"
)

const WxworkBotNews = "wxwork_bot_news"

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

type News struct {
	Articles []*Article `json:"articles"`
}

type NewsMessage struct {
	MsgType string `json:"msgtype"`
	News    *News  `json:"news"`
}

func SendWxWorkBotNews(client *http.Client, sourceItem *source.Source, target string) error {
	body := &NewsMessage{
		MsgType: "news",
		News: &News{
			Articles: []*Article{
				{
					Title:       sourceItem.Title,
					Url:         sourceItem.Url,
					Description: sourceItem.Description,
					PicUrl:      sourceItem.PicUrl,
				},
			},
		},
	}
	j, e0 := json.Marshal(body)
	if e0 != nil {
		return e0
	}
	resp, e1 := http.Post(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", target), "application/json", strings.NewReader(string(j)))
	if e1 != nil {
		return e1
	}
	_, e2 := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return e2
}
