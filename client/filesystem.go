package client

import "runtime"
import "os"
import "github.com/nlopes/slack"

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func SaveMsg(team, room string, msg slack.Msg) {
	lpath := UserHomeDir() + "/.grepslak/" + team + "/" + room + "/msg/" + msg.Timestamp
	file, _ := os.OpenFile(lpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
	file.WriteString(msg.Text)
	file.Close()
	for _, a := range msg.Attachments {
		lpath = UserHomeDir() + "/.grepslak/" + team + "/" + room + "/attachments/" + msg.Timestamp
		file, _ = os.OpenFile(lpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
		file.WriteString(a.Title + "|" + a.TitleLink + "\n" + a.Text)
		file.Close()
	}
}
