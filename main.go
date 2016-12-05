package main

import "github.com/andrewarrow/grepslak/client"
import "os"

func main() {
	//command := os.Args[0]
	if len(os.Args) == 1 {
		client.ListTeams()
	}
}
