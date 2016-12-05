package client

import "github.com/nlopes/slack"
import "fmt"
import "os"
import "strings"
import "time"

func SaveHistory(team, room string) {
	ts := time.Now().Unix() - int64(31536000*5)
	tss := fmt.Sprintf("%d", ts)

	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		api := slack.New(tokens[i])
		hp := slack.HistoryParameters{Oldest: tss, Latest: "", Count: 10, Inclusive: false, Unreads: false}
		list, _ := api.GetGroupHistory(room, hp)
		for _, r := range list.Messages {
			fmt.Println(r.Msg.Text)
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
	}
}

func ListTeams() {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	for _, team := range teams {
		fmt.Println(team)
	}
}

/*
func GetTeams() []*slack.TeamInfo {
	list := make([]*slack.TeamInfo, 0)

	slack_teams, _ := strconv.ParseInt(os.Getenv("SLACK_TEAMS"), 10, 64)
	i := int64(0)
	for {
		key := fmt.Sprintf("SLACK_TOKEN_%d", i)
		api := slack.New(os.Getenv(key))
		r, err := api.GetTeamInfo()
		if err == nil {
			list = append(list, r)
		}
		i++
		if i >= slack_teams {
			break
		}
	}

	return list
}*/
