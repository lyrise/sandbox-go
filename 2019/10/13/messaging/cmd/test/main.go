package main

import (
	"fmt"

	engine "github.com/lyrise/sandbox-go/2019/10/13/messaging/engine"
	message "github.com/lyrise/sandbox-go/2019/10/13/messaging/message"
)

type Status struct {
	nodes map[message.Addr]engine.Worker
}

func NewStatus() *Status {
	var s Status
	return &s
}

func (s *Status) Find(addr message.Addr) (engine.Worker, error) {
	return s.nodes[addr], nil
}

func main() {
	s := NewStatus()
	n := engine.NewWorker(s)
	n.SendMessage("test2")
	fmt.Println(n.ReceiveMessage())
}
