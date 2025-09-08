package wxwork_bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/yichya/unipodcast/common/htmlutil"
	"github.com/yichya/unipodcast/pipeline/source"
)

const WxWorkBotText = "wxwork_bot_text"

type TextMessageContent struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type TextMessage struct {
	MsgType string              `json:"msgtype"`
	Text    *TextMessageContent `json:"text"`
}

func SendWxWorkBotText(client *http.Client, sourceItem *source.Source, target string) error {
	contentCut := htmlutil.ExtractText(sourceItem.Description)
	content := fmt.Sprintf("%s\n\n%s", sourceItem.Title, contentCut)
	if len(content) > 2000 {
		var truncated []byte
		for i := range content {
			if len(truncated)+utf8.RuneLen(rune(content[i])) > 1997 {
				content = string(truncated) + "..."
				break
			}
			truncated = append(truncated, content[i])
		}
	}
	body := &TextMessage{
		MsgType: "text",
		Text: &TextMessageContent{
			Content: content,
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
