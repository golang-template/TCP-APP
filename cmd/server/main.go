package main

import (
	"TCP-server/server"
	"fmt"
	"log"
)

func main() {
	fmt.Println("🚀connecting to server...")

	srv := server.NewServer("127.0.0.1:8080")
	log.Fatal(srv.Start())
}
