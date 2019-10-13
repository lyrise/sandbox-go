package engine

import (
	"math/rand"

	message "github.com/lyrise/sandbox-go/2019/10/13/messaging/message"
)

type Worker struct {
	Addr   message.Addr
	status *workerStatus
}

type workerStatus struct {
	sendChan    chan interface{}
	receiveChan chan interface{}
	finder      Finder
}

func NewWorker(finder Finder) *Worker {
	var worker Worker
	worker.Addr = message.Addr(rand.Uint64())
	worker.status = new(workerStatus)
	worker.status.sendChan = make(chan interface{}, 10)
	worker.status.receiveChan = make(chan interface{}, 10)
	worker.status.finder = finder

	go eventLoop(&worker)

	return &worker
}

func eventLoop(worker *Worker) {
	tmp := <-worker.status.sendChan
	worker.status.receiveChan <- tmp
}

func (worker *Worker) SendMessage(msg interface{}) {
	worker.status.sendChan <- msg
}

func (worker *Worker) ReceiveMessage() interface{} {
	return <-worker.status.receiveChan
}
