package targets

import (
	"fmt"
	"github.com/yichya/unipodcast/common/proxy_client"
	"github.com/yichya/unipodcast/pipeline/source"
	"github.com/yichya/unipodcast/targets/echo"
	"github.com/yichya/unipodcast/targets/telegram"
)

func SendOneFeed(source *source.Source, proxyUrl string, targetType string, targetArgs string) error {
	client := proxy_client.ProxyClient(proxyUrl)
	switch targetType {
	case telegram.BotAudio:
		{
			return telegram.SendBotAudio(client, source, targetArgs)
		}
	case echo.Stdout:
		{
			return echo.SendStdout(source, targetArgs)
		}
	}
	return fmt.Errorf("unknown targetType %s", targetType)
}
