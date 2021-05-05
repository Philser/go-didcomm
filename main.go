package main

import (
	"bufio"
	"fmt"
	"go-didcomm/service"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Start as service or run client (s/c)?")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Encountered error: " + err.Error())
		}

		input := strings.Trim(text, "\n")
		if input == "c" {
			log.Fatal("client is not implemented yet")
		} else if input == "s" {
			service.Start()
		} else {
			continue
		}
	}
}
