package command

import "fmt"

type Command interface {
	Target() int
	Message() string
}

func NewCommand(target int, message string) Command {
	return &command{
		target:  target,
		message: message,
	}
}

const BroadcastTarget = 255

type command struct {
	target  int
	message string
}

func (c *command) Target() int {
	return c.target
}

func (c *command) Message() string {
	return c.message
}

func (c *command) String() string {
	return fmt.Sprintf("{target:%d, message:%s}", c.target, c.message)
}

func (c *command) IsBroadcast() bool {
	return IsBroadcast(c.target)
}

func IsBroadcast(target int) bool {
	return target == BroadcastTarget
}
