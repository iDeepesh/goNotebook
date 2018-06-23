package main

import (
	"sync"

	"github.com/iDeepesh/goNotebook/server/http"
	"github.com/iDeepesh/goNotebook/server/tcp"
)

func main() {
	runTCP := func() {
		go tcp.SingleConnectionWriteServer()
		go tcp.SingleConnectionReadServer()
		go tcp.MultipleConnectionServer()
	}

	runTcpHTTP := func() {
		go http.SimpleHttpServer()
		go http.BetterHttpServer()
	}

	var wg sync.WaitGroup
	wg.Add(1)

	runTCP()
	runTcpHTTP()

	wg.Wait()
}
