package client

import "github.com/nlopes/slack"
import "fmt"
import "os"
import "strings"

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
