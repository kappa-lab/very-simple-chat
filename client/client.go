package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/kappa-lab/very-simple-chat/command"
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

	b := make([]byte, 256)
	for {
		size, err := conn.Read(b)
		if err == nil {
			fmt.Printf("[Receive]: %s\n", b[:size])
		} else if err == io.EOF {
			break
		}
	}
}

func sendGreeting(conn net.Conn) {
	msg := "Hello everybody"
	target := command.BroadcastTarget
	conn.Write(createSendData(target, msg))
}

func createSendData(target int, message string) []byte {
	fmt.Printf("[Send]: target=%d, mesaage=%s\n", target, message)
	var buf bytes.Buffer
	msg := []byte(message)
	buf.WriteByte(byte(target))
	buf.Write(msg)
	return buf.Bytes()
}

func parseInput(conn net.Conn) {
	for {
		/** use double quote
		 * unicast
		 * {"target":1, "message":"hello 1"}
		 *
		 * broadcast (255 as broadcast)
		 * {"target":255, "message":"hello evrybody"}
		 */
		reader := bufio.NewReader(os.Stdin)
		dec := json.NewDecoder(reader)
		var raw rawCmd

		if err := dec.Decode(&raw); err != nil {
			fmt.Println("[Invalid Command]:", err)
			continue
		}
		var cmd = raw.toCommand()

		fmt.Printf("[input]:%s\n", cmd)
		conn.Write(createSendData(cmd.Target(), cmd.Message()))
	}
}

type rawCmd struct {
	Target  int
	Message string
}

func (r *rawCmd) toCommand() command.Command {
	return command.NewCommand(r.Target, r.Message)
}
