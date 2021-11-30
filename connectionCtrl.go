package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/kappa-lab/very-simple-chat/command"
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
		conn.Write([]byte(msg))
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

func (c *connectionCtrl) ReadMessage(myRoom Room, conn net.Conn) {
	b := make([]byte, 256)

	for {
		size, err := conn.Read(b)
		if err == io.EOF {
			fmt.Printf("[ConnCtrl(%d) Leave]\n", c.id)
			myRoom.RemoveListener(c.id)
			body, _ := json.Marshal(leaveData{
				LeaveId: c.id,
				Member:  myRoom.GetIds(),
			})
			myRoom.Broadcast(string(body))
			break
		} else if size > 0 {
			target := int(b[0]) // head 1byte is target
			msg := b[1:size]
			fmt.Printf("[ConnCtrl(%d) Read]: target=%d, message=%s\n", c.id, target, msg)

			body, _ := json.Marshal(messageData{
				Sender:  c.id,
				Message: string(msg),
			})

			if command.IsBroadcast(target) {
				myRoom.Broadcast(string(body))
			} else {
				myRoom.Unicast(target, string(body))
			}
		}
	}
}
