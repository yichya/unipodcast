package telegram

import (
	"fmt"
	"github.com/yichya/unipodcast/pipeline/source"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const BotAudio = "telegram_bot_audio"

func SendBotAudio(client *http.Client, sourceItem *source.Source, target string) error {
	if client == nil {
		client = http.DefaultClient
	}
	tgBotToken := os.Getenv("TG_BOT_TOKEN")
	api, err := tgbotapi.NewBotAPIWithClient(tgBotToken, tgbotapi.APIEndpoint, client)
	if err != nil {
		log.Printf("bot %v init failed %v", tgBotToken, err)
		return err
	}
	p, err := url.Parse(sourceItem.FileUrl)
	if err != nil {
		log.Printf("parse url %v init failed %v", sourceItem.FileUrl, err)
		return err
	}
	f, err := http.Get(sourceItem.FileUrl)
	if err != nil {
		log.Printf("fetch %v failed %v", sourceItem.FileUrl, err)
		return err
	}
	chatId, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		log.Printf("parse target %v failed %v", target, err)
		return err
	}
	fns := strings.Split(p.Path, "/")
	audio := tgbotapi.NewAudio(chatId, tgbotapi.FileReader{
		Name:   fns[len(fns)-1],
		Reader: f.Body,
	})
	audio.Title = sourceItem.Title
	audio.Duration = int(sourceItem.Duration)
	audio.Performer = sourceItem.Performer
	tags := []string{sourceItem.Performer}
	audio.Caption = fmt.Sprintf("%s #%s", audio.Title, strings.Join(tags, " #"))
	_, err = api.Send(audio)
	return err
}
