package client

import "github.com/nlopes/slack"
import "fmt"
import "os"
import "strconv"

func ListRooms(team string) {
}

func ListTeams() {
	for _, ti := range GetTeams() {
		fmt.Println(ti.ID, ti.Name, ti.Domain)
	}
}

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
}
