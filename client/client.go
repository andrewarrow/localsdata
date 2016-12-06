package client

import "github.com/nlopes/slack"
import "fmt"
import "os"
import "strings"
import "time"

func SaveHistory(team, room string) {
	//ts := time.Now().Unix() - int64(31536000*5)
	//tss := fmt.Sprintf("%d", ts)

	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		api := slack.New(tokens[i])

		gfp := slack.GetFilesParameters{Channel: room, Count: 100}
		list, p, err := api.GetFiles(gfp)
		for _, r := range list {
			SaveFile(team, room, r.URLPrivate, tokens[i], int64(r.Timestamp))
			fmt.Println(int64(r.Timestamp), r.URLPrivateDownload)
		}
		fmt.Println(p, err)
	}
}

func SaveHistory2(team, room string) {
	ts := time.Now().Unix() - int64(31536000*5)
	tss := fmt.Sprintf("%d", ts)

	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		api := slack.New(tokens[i])

		for {
			hp := slack.HistoryParameters{Oldest: tss, Latest: "", Count: 100, Inclusive: false, Unreads: false}
			list, _ := api.GetGroupHistory(room, hp)
			stamps := make([]string, 0)
			for _, r := range list.Messages {
				SaveMsg(team, room, r.Msg)
				fmt.Println(r.Msg.Timestamp)
				//fmt.Println(r.Msg.Text)
				fmt.Println(r.Msg.Attachments)
				stamps = append(stamps, r.Msg.Timestamp)
			}
			tss = stamps[0]
			fmt.Println("-----")
			time.Sleep(time.Second)
		}
	}
}

func ListRooms(team string) {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		api := slack.New(tokens[i])
		list, _ := api.GetChannels(false)
		for _, r := range list {
			fmt.Println(r.ID, r.Name)
		}
		list2, _ := api.GetGroups(false)
		for _, r := range list2 {
			fmt.Println(r.ID, r.Name)
		}
		list3, _ := api.GetIMChannels()
		for _, r := range list3 {
			fmt.Println(r.ID, r.User)
		}
	}
}

func ListTeams() {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	for _, team := range teams {
		fmt.Println(team)
	}
}
