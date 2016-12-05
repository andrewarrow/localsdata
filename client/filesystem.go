package client

import "runtime"
import "os"

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

func SaveInProgress(path string) {
	lpath := UserHomeDir() + "/foo"
	file, _ := os.OpenFile(lpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
	file.WriteString(path)
	file.Close()
}

func RemoveInProgress() {
	path := UserHomeDir() + "/foo"
	os.Remove(path)
}
