package client

import (
	"bufio"
	"fmt"
	"go-didcomm/db"
	"log"
	"os"
	"strings"
)

func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var actionType string
		var receiver string
		var message string
		var inputError error

		fmt.Println("What do you want to do?")
		fmt.Println("  1. Read old messages")
		fmt.Println("  2. Send message")
		actionType, inputError = reader.ReadString('\n')
		if inputError != nil {
			log.Fatal(inputError)
		}

		switch strings.Trim(actionType, "\n") {
		case "1":
			messages := db.GetMessages()

			for i := 0; i < len(messages); i++ {
				db.LogMessage(messages[i])
			}

			fmt.Println("\n")
		case "2":
			fmt.Println("please insert the did of the message receiver:")
			for {
				receiver, inputError = reader.ReadString('\n')
				if inputError != nil {
					log.Fatal(inputError)
				}

				if !strings.HasPrefix(receiver, "did:") {
					fmt.Println("Not a valid did!")
					continue
				}
				break
			}

			fmt.Println("	- Please enter the message")
			message, inputError = reader.ReadString('\n')
			if inputError != nil {
				log.Fatal(inputError)
			}
			fmt.Println("receiver: " + receiver)
			fmt.Println("message: " + message)
		default:
			continue
		}

		fmt.Println("press any key to continue...")
		reader.ReadString('\n')
		break
	}
}
