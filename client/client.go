package client

//its better to have cmd/client and cmd/server  OR  server/cmd and client/cmd ????????????

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCPClient manage connections between client and server
type TCPClient struct {
	ServerAddr string
	Conn       net.Conn
}

func NewClient(serverAddr string) (*TCPClient, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("⛔error in connecting to server: %v", err)
	}

	fmt.Println("✅ connected to server...")
	return &TCPClient{ServerAddr: serverAddr, Conn: conn}, nil
}

func (c *TCPClient) SendMessage() {
	defer c.Conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("📝 enter your message: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("👋 exit the application")
			return
		}

		// send message to server
		fmt.Fprintf(c.Conn, text+"\n")

		// receive respnse from server
		response, _ := bufio.NewReader(c.Conn).ReadString('\n')
		fmt.Println("📤server's response:", response)
	}
}
