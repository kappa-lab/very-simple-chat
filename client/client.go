package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/kappa-lab/very-simple-chat/command"
	"github.com/kappa-lab/very-simple-chat/protocol"
)

func main() {
	fmt.Println("[Client Start]")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer func() {
		conn.Close()
		fmt.Println("[Client End]")
	}()

	go func() {
		time.Sleep(1000 * time.Millisecond)
		sendGreeting(conn)
	}()

	go parseInput(conn)

	for {
		body, err := protocol.Read(conn)
		if err == nil {
			fmt.Printf("[Receive]: %s\n", body)
		} else if err == io.EOF {
			break
		}
	}
}

func sendGreeting(conn net.Conn) {
	raw, _ := json.Marshal(command.NewCommand(command.BroadcastTarget, "Hello everybody"))
	fmt.Printf("[SendGreeting]: %s\n", raw)
	protocol.Write(conn, raw)
}

func parseInput(conn net.Conn) {
	for {
		/** use double quote
		 * unicast
		 * {"target":1, "message":"hello 1"}
		 *
		 * broadcast (255 as broadcast)
		 * {"target":255, "message":"hello everybody"}
		 */
		reader := bufio.NewReader(os.Stdin)
		dec := json.NewDecoder(reader)
		var cmd command.Command

		if err := dec.Decode(&cmd); err != nil {
			fmt.Println("[Invalid Command]:", err)
			continue
		}

		json, _ := json.Marshal(cmd)
		fmt.Printf("[Input]:%s\n", json)
		protocol.Write(conn, json)
	}
}
