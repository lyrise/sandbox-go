package main

import (
	"log"
	"strconv"

	core "github.com/lyrise/sandbox-go/channel-sample/core"
)

func main() {
	ch := make(chan interface{})
	go receiver(ch)

	for i := 0; i < 100; i++ {
		var pz core.Packet
		pz.Value = "count: " + strconv.Itoa(i)

		ch <- pz
	}
}

func receiver(ch <-chan interface{}) {
	for {
		p := <-ch
		switch p := p.(type) {
		case core.Packet:
			log.Printf("%s", p.Value)
		}
	}
}
