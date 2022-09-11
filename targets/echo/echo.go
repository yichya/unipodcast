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
		a := strings.Split(strings.TrimSpace(x), ":")
		switch a[0] {
		case "space":
			{
				fmt.Print(` `)
			}
		case "comma":
			{
				fmt.Print(`,`)
			}
		case "quote":
			{
				fmt.Print(`"`)
			}
		case "colon":
			{
				fmt.Print(`:`)
			}
		case "print":
			{
				if len(a) > 1 {
					fmt.Print(a[1])
				}
			}
		case "id":
			{
				fmt.Print(sourceItem.Id)
			}
		case "title":
			{
				fmt.Print(sourceItem.Title)
			}
		case "performer":
			{
				fmt.Print(sourceItem.Performer)
			}
		case "pub_date":
			{
				if sourceItem.PubDate != nil {
					fmt.Print(sourceItem.PubDate.Format(constants.DefaultTimeFormat))
				} else {
					fmt.Print(constants.DefaultTimeFormat)
				}
			}
		case "url":
			{
				fmt.Print(sourceItem.FileUrl)
			}
		case "endl":
			{
				fmt.Println()
			}
		}
	}
	fmt.Println()
	return nil
}
