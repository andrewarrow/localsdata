package client

import "runtime"
import "fmt"
import "os"
import "path/filepath"
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
	// /Users/aa/.grepslak/0x7a69/G0ETQAEBF
	base := filepath.Join(UserHomeDir(), ".grepslak")
	os.Mkdir(base, 0777)
	dir := filepath.Join(base, team, room)
	fmt.Println(dir)
	os.MkdirAll(dir, 0777)

	os.MkdirAll(filepath.Join(dir, "msg"), 0777)
	os.MkdirAll(filepath.Join(dir, "attachments"), 0777)
	os.MkdirAll(filepath.Join(dir, "files"), 0777)

	lpath := filepath.Join(dir, "msg", msg.Timestamp)
	fmt.Println(lpath)
	file, _ := os.OpenFile(lpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
	file.WriteString(msg.Text)
	file.Close()
	for _, a := range msg.Attachments {
		lpath := filepath.Join(dir, "attachments", msg.Timestamp)
		file, _ = os.OpenFile(lpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
		file.WriteString(a.Title + "|" + a.TitleLink + "\n" + a.Text)
		file.Close()
	}
}
