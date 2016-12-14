package main

import "github.com/andrewarrow/localsdata/client"
import "os"

func main() {
	if len(os.Args) == 1 {
		client.ListTeams()
		return
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "clean" {
			client.Clean()
			return
		}
		client.ListRooms(os.Args[1])
		return
	}
	if len(os.Args) == 3 {
		client.SaveHistory(os.Args[1], os.Args[2])
		return
	}
}
