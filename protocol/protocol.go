package protocol

type Header struct {
	Action     int
	BodyLength int
}

type Action int

const (
	Join Action = iota + 1
	Leave
	Message
)

func Read(b []byte) {}

func Write(action int, body string) {}
