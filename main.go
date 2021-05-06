package main

import (
	"fmt"
	"go-didcomm/client"
	"go-didcomm/db"
	"go-didcomm/server"
	"time"
)

func logHeader() {
	fmt.Println("\x1B[2J\x1B[1;1H")
	fmt.Println("--------------------------------------")
	fmt.Println("| DIDCOMM - GO                       |")
	fmt.Printf("|    SERVER: listening on PORT %s  |\n", server.PORT)
	fmt.Println("|    CLIENT: waiting for user inputs |")
	fmt.Println("--------------------------------------\n\n")
}

func printMessages() {

}

func main() {
	db.GetDb()

	go server.Start()

	time.Sleep(10 * time.Millisecond)

	for {
		logHeader()
		printMessages()
		client.Start()
	}
}
