package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/kappa-lab/very-simple-chat/command"
	"github.com/kappa-lab/very-simple-chat/protocol"
)

type ConnectionCtrl interface {
	Id() int
	EventListener() chan string
	HandleEventListener(conn net.Conn)
	ReadMessage(myRoom Room, conn net.Conn)
}

func NewConnectionCtrl(id int) ConnectionCtrl {
	return &connectionCtrl{
		id:            id,
		eventListener: make(chan string, 0),
	}
}

type connectionCtrl struct {
	id            int
	eventListener chan string
}

func (c *connectionCtrl) Id() int                    { return c.id }
func (c *connectionCtrl) EventListener() chan string { return c.eventListener }

func (c *connectionCtrl) HandleEventListener(conn net.Conn) {
	for {
		msg := <-c.eventListener
		fmt.Printf("[ConnCtrl(%d) Event]: %s\n", c.id, msg)
		protocol.Write(conn, []byte(msg))
	}
}

type messageData struct {
	Sender  int    `json:"sender"`
	Message string `json:"message"`
}

type leaveData struct {
	LeaveId int   `json:"leaveId"`
	Member  []int `json:"member"`
}
type rawCmd struct {
	Target  int
	Message string
}

func (r *rawCmd) toCommand() command.Command {
	return command.NewCommand(r.Target, r.Message)
}

func (c *connectionCtrl) ReadMessage(myRoom Room, conn net.Conn) {

	for {
		body, err := protocol.Read(conn)
		if err == io.EOF {
			fmt.Printf("[ConnCtrl(%d) Leave]\n", c.id)
			myRoom.RemoveListener(c.id)
			body, _ := json.Marshal(leaveData{
				LeaveId: c.id,
				Member:  myRoom.GetIds(),
			})
			myRoom.Broadcast(string(body))
			break
		}
		cmd := rawCmd{}
		json.Unmarshal(body, &cmd)

		fmt.Printf("[ConnCtrl(%d) Read]: %s\n", c.id, body)

		m, _ := json.Marshal(messageData{
			Sender:  c.id,
			Message: cmd.Message,
		})

		if command.IsBroadcast(cmd.Target) {
			myRoom.Broadcast(string(m))
		} else {
			myRoom.Unicast(cmd.Target, string(m))
		}
	}
}
