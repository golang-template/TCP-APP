package main

import (
	"TCP-server/client"
	"fmt"
	"log"
)

func main() {
	fmt.Println("🚀client is running")

	cli, err := client.NewClient("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	cli.SendMessage()
}
