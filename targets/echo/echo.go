package echo

import (
	"errors"
	"fmt"
	"github.com/yichya/unipodcast/common/constants"
	"github.com/yichya/unipodcast/pipeline/source"
	"strings"
)

const Stdout = "echo_stdout"

func SendStdout(sourceItem *source.Source, target string) error {
	if sourceItem == nil {
		return errors.New("nil sourceItem")
	}
	fields := strings.Split(target, ",")
	for _, x := range fields {
		switch strings.TrimSpace(x) {
		case "id":
			{
				fmt.Printf("%s ", sourceItem.Id)
			}
		case "title":
			{
				fmt.Printf("%s ", sourceItem.Title)
			}
		case "pub_date":
			{
				if sourceItem.PubDate != nil {
					fmt.Printf("%s ", sourceItem.PubDate.Format(constants.DefaultTimeFormat))
				} else {
					fmt.Printf("%s ", constants.DefaultTimeFormat)
				}
			}
		}
	}
	fmt.Println()
	return nil
}
