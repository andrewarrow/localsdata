package client

import "runtime"
import "fmt"
import "os"
import "net/http"
import "path/filepath"
import "github.com/nlopes/slack"
import "strings"
import "io/ioutil"
import "github.com/mvdan/xurls"
import "time"
import "strconv"

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

func SaveFile(team, room, url, token string, ts int64) {
	dir := SetupDirs(team, room)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	res, _ := client.Do(req)
	data, _ := ioutil.ReadAll(res.Body)

	parts := strings.Split(url, ".")

        t := time.Unix(int64(ts), 0)
        dstr := t.Format("2006-01-02")

	base := filepath.Join(dir, "files", dstr)
	os.Mkdir(base, 0700)

	lpath := filepath.Join(dir, "files", dstr, fmt.Sprintf("%d.%s", ts, parts[len(parts)-1]))
	//fmt.Println(lpath)
	file, _ := os.OpenFile(lpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	file.Write(data)
	file.Close()
}

func CleanDir(team string) {
	base := filepath.Join(UserHomeDir(), ".grepslak")
	fmt.Println("removing ", team)
	os.RemoveAll(filepath.Join(base, team))
}

func SetupDirs(team, room string) string {
	base := filepath.Join(UserHomeDir(), ".grepslak")
	os.Mkdir(base, 0700)
	dir := filepath.Join(base, team, room)
	//fmt.Println(dir)
	os.MkdirAll(dir, 0700)

	os.MkdirAll(filepath.Join(dir, "msg"), 0700)
	os.MkdirAll(filepath.Join(dir, "attachments"), 0700)
	os.MkdirAll(filepath.Join(dir, "files"), 0700)

	return dir
}

func LookForLinks(text string) {
	items := xurls.Strict.FindAllString(text, -1)
	for _, url := range items {
		val, ok := links[url]
		if ok {
			links[url] = val + 1
		} else {
			links[url] = 1
		}
	}
}

func SaveMsg(team, room string, msg slack.Msg) {
	dir := SetupDirs(team, room)

        ts, _ := strconv.ParseFloat(msg.Timestamp, 64)
        t := time.Unix(int64(ts), 0)
        dstr := t.Format("2006-01-02")

	base := filepath.Join(dir, "msg", dstr)
	os.Mkdir(base, 0700)

	lpath := filepath.Join(dir, "msg", dstr, msg.Timestamp)
	//fmt.Println(lpath)
	file, _ := os.OpenFile(lpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	file.WriteString(msg.Text)
	file.Close()
	LookForLinks(msg.Text)
	for _, a := range msg.Attachments {
		lpath := filepath.Join(dir, "attachments", msg.Timestamp)
		file, _ = os.OpenFile(lpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		file.WriteString(a.Title + "|" + a.TitleLink + "\n" + a.Text)
		file.Close()
	}
}
