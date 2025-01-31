package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type TCPServer struct {
	Address string
}

func NewServer(address string) *TCPServer {
	return &TCPServer{Address: address}
}

func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return fmt.Errorf("⛔error in running server: %v", err)
	}
	defer listener.Close()

	fmt.Println("✅running in server", s.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("⛔error in joining client", err)
			continue
		}
		fmt.Println("🔗new client joined", conn.RemoteAddr())

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("🔌client disconnected", conn.RemoteAddr())
			return
		}

		message = strings.TrimSpace(message)
		fmt.Println("📩 received message :", message)

		response := fmt.Sprintf("📢 server: you told « %s »\n", message)
		conn.Write([]byte(response))
	}
}
