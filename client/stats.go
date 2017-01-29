package client

import "github.com/nlopes/slack"
import "fmt"
import "os"
import "strings"
import "time"

func Stats(team string) {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		fmt.Println(team)
		api := slack.New(tokens[i])

		list3, _ := api.GetIMChannels()
		for _, r := range list3 {
			u, _ := api.GetUserInfo(r.User)
			fmt.Println(r.ID, u.Name)
			ts := time.Now().Unix() - int64(31536000*5)
			tss := fmt.Sprintf("%d", ts)
			j := 0
			count := 0
			for {
				fmt.Println("syncing ", j)
				j += 1000
				hp := slack.HistoryParameters{Oldest: tss, Latest: "", Count: 1000, Inclusive: false, Unreads: false}
				list, _ := api.GetIMHistory(r.ID, hp)

				if list == nil {
					break
				}

				stamps := make([]string, 0)
				for _, rr := range list.Messages {
					//fmt.Println(r.Msg.Timestamp)
					//fmt.Println(r.Msg.Text)
					//fmt.Println(r.Msg.Attachments)
					stamps = append(stamps, rr.Msg.Timestamp)
					count += 1
				}
				if len(stamps) == 0 {
					break
				}
				tss = stamps[0]
			}
			fmt.Println(count)
		}
	}
}
