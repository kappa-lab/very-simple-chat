package protocol

import (
	"bytes"
	"net"
)

/**
* Protocol Structure
* |----Header(1byte)-----|-----Body(max255byte)-----|
*       BodyLength                  Body
 */
func Read(conn net.Conn) ([]byte, error) {
	head := make([]byte, 1)
	_, err := conn.Read(head)
	if err != nil {
		return nil, err
	}

	body := make([]byte, head[0])
	_, err = conn.Read(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Write(conn net.Conn, body []byte) (int, error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(len(body)))
	buf.Write(body)
	return conn.Write(buf.Bytes())
}
