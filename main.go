package main

import "github.com/andrewarrow/grepslak/client"
import "os"

func main() {
	if len(os.Args) == 1 {
		client.ListTeams()
		return
	}
	if len(os.Args) == 2 {
		client.ListRooms(os.Args[1])
	}
}
