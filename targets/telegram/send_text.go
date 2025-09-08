package telegram

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yichya/unipodcast/common/htmlutil"
	"github.com/yichya/unipodcast/pipeline/source"
)

const BotText = "telegram_bot_text"

func SendBotText(client *http.Client, sourceItem *source.Source, target string) error {
	if client == nil {
		client = http.DefaultClient
	}
	tgBotToken := os.Getenv("TG_BOT_TOKEN")
	api, err := tgbotapi.NewBotAPIWithClient(tgBotToken, tgbotapi.APIEndpoint, client)
	if err != nil {
		log.Printf("bot %v init failed %v", tgBotToken, err)
		return err
	}
	chatId, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		log.Printf("parse target %v failed %v", target, err)
		return err
	}

	textMessage := tgbotapi.NewMessage(chatId, fmt.Sprintf("<a href='%s'><strong>%s</strong></a>\n\n%s", sourceItem.Url, sourceItem.Title, htmlutil.ExtractText(sourceItem.Description)))
	textMessage.ParseMode = tgbotapi.ModeHTML
	_, err = api.Send(textMessage)
	return err
}
