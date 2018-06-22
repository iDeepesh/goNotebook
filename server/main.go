package main

import (
	"sync"

	"github.com/iDeepesh/goNotebook/server/tcp"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go tcp.SingleConnectionWriteServer()
	go tcp.SingleConnectionReadServer()
	go tcp.MultipleConnectionServer()

	wg.Wait()
}
