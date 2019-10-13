package engine

import (
	message "github.com/lyrise/sandbox-go/2019/10/13/messaging/message"
)

type Finder interface {
	Find(message.Addr) (Worker, error)
}
