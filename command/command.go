package command

import "fmt"

type Command struct {
	Target  int    `json:"target"`
	Message string `json:"message"`
}

func NewCommand(target int, message string) *Command {
	return &Command{
		Target:  target,
		Message: message,
	}
}

const BroadcastTarget = 255

func (c *Command) String() string {
	return fmt.Sprintf("{target:%d, message:%s}", c.Target, c.Message)
}

func (c *Command) IsBroadcast() bool {
	return IsBroadcast(c.Target)
}

func IsBroadcast(target int) bool {
	return target == BroadcastTarget
}
