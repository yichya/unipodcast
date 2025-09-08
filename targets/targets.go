package targets

import (
	"fmt"

	"github.com/yichya/unipodcast/common/proxy_client"
	"github.com/yichya/unipodcast/pipeline/source"
	"github.com/yichya/unipodcast/targets/echo"
	"github.com/yichya/unipodcast/targets/telegram"
	"github.com/yichya/unipodcast/targets/wxwork_bot"
)

func SendOneFeed(source *source.Source, proxyUrl string, targetType string, targetArgs string) error {
	client := proxy_client.ProxyClient(proxyUrl)
	switch targetType {
	case telegram.BotAudio:
		{
			return telegram.SendBotAudio(client, source, targetArgs)
		}
	case telegram.BotText:
		{
			return telegram.SendBotText(client, source, targetArgs)
		}
	case echo.Stdout:
		{
			return echo.SendStdout(source, targetArgs)
		}
	case wxwork_bot.WxworkBotNews:
		{
			return wxwork_bot.SendWxWorkBotNews(client, source, targetArgs)
		}
	case wxwork_bot.WxWorkBotText:
		{
			return wxwork_bot.SendWxWorkBotText(client, source, targetArgs)
		}
	}
	return fmt.Errorf("unknown targetType %s", targetType)
}
