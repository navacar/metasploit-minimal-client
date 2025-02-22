package main

import (
	"fmt"
	"log"
	"metasploit-minimal/rpc"
	"os"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSPASS")
	user := "msf"

	if host == "" || pass == "" {
		log.Fatal("Missing required environment variable MSFHOST or MSPASS")
	}

	msf, err := rpc.New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}
	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}
}
