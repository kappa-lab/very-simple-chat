package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func startServer() {
	defer fmt.Println("[Server Stop]")

	fmt.Println("[Server Start]")
	ln, _ := net.Listen("tcp", ":8080") // ignore Error Handling

	var room = NewRoom(0)
	for {
		conn, _ := ln.Accept() // ignore Error Handling
		initConnection(conn, room)
	}
}

func initConnection(conn net.Conn, r Room) {
	connCtrl := r.CreateConnectionCtrl()

	type greetingMessage struct {
		ConnId int   `json:"connId"`
		RoomId int   `json:"roomId"`
		Member []int `json:"member"`
	}

	greeting := greetingMessage{
		ConnId: connCtrl.Id(),
		RoomId: r.Id(),
		Member: append(r.GetIds(), connCtrl.Id()),
	}

	msg, _ := json.Marshal(greeting)
	fmt.Printf("[Init Connection] ---> %s\n", msg)

	conn.Write(msg)

	r.Broadcast(string(msg))
	r.AddListener(connCtrl.Id(), connCtrl.EventListener())

	go connCtrl.HandleEventListener(conn)
	go connCtrl.ReadMessage(r, conn)
}
