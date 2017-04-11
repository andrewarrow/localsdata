package client

import "github.com/nlopes/slack"
import "fmt"
import "os"
import "strings"
import "time"
import "github.com/mvdan/xurls"

var links = make(map[string]int)

func Clean() {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	for _, t := range teams {
		CleanDir(t)
	}
}

func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}
func rightPad(s string, padStr string, pLen int) string {
	return s + strings.Repeat(padStr, pLen)
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func CheckForHit(j int, thing slack.CtxMessage, channel string) {
	items := xurls.Strict.FindAllString(thing.Text, -1)
	for _, item := range items {
		if channel[0] == 'U' {
			fmt.Printf("%s|%s\n", leftPad2Len(thing.Username, " ", 15), item)
		} else {
			fmt.Printf("%s%s|%s\n", leftPad2Len(channel, " ", 15), leftPad2Len(thing.Username, " ", 15), item)
		}
	}
}

func Search() {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, team := range teams {
		fmt.Println(team)
		api := slack.New(tokens[i])

		sp := slack.SearchParameters{}
		sp.Sort = "timestamp"
		sp.SortDirection = "desc"
		sp.Highlight = false
		sp.Count = 20
		sp.Page = 1
		list, err := api.SearchMessages("from:me http", sp)
		if err != nil {
			fmt.Println(err)
		}
		for j, r := range list.Matches {
			cn := r.Channel.Name
			CheckForHit(j, r.Previous, cn)
			CheckForHit(j, r.Previous2, cn)
			fake := slack.CtxMessage{}
			fake.Username = r.Username
			fake.Text = r.Text
			CheckForHit(j, fake, cn)
			CheckForHit(j, r.Next, cn)
			CheckForHit(j, r.Next2, cn)
		}
	}
}

func Say(team, room, what string) {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		api := slack.New(tokens[i])
		//_ := api.Say(room, hp)
    fmt.Println(api)
  }
}

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

		j := 0
		for {
			fmt.Println("syncing ", j)
			j += 1000
			hp := slack.HistoryParameters{Oldest: tss, Latest: "", Count: 1000, Inclusive: false, Unreads: false}
			list, _ := api.GetGroupHistory(room, hp)
			stamps := make([]string, 0)
			for _, r := range list.Messages {
				SaveMsg(team, room, r.Msg)
				//fmt.Println(r.Msg.Timestamp)
				//fmt.Println(r.Msg.Text)
				//fmt.Println(r.Msg.Attachments)
				stamps = append(stamps, r.Msg.Timestamp)
			}
			if len(stamps) == 0 {
				break
			}
			tss = stamps[0]
			//fmt.Println("-----")
			//time.Sleep(time.Second)
		}

		gfp := slack.GetFilesParameters{Channel: room, Count: 100}
		list, _, _ := api.GetFiles(gfp)
		for _, r := range list {
			SaveFile(team, room, r.URLPrivate, tokens[i], int64(r.Timestamp))

			//fmt.Println(r.Name, int64(r.Timestamp), r.URLPrivateDownload)
		}
	}

	for k, v := range links {
		if v > 2 {
			fmt.Println(k, v)
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
			u, _ := api.GetUserInfo(r.User)
			fmt.Println(r.ID, u.Name)
		}
	}
}

func ListTeams() {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	for _, team := range teams {
		fmt.Println(team)
	}
}
