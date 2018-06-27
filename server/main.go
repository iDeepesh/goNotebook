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

	runTCPHTTP := func() {
		go http.TcpHttpServer()
		go http.BetterTcpHttpServer()
	}

	runHTTP := func() {
		go http.SimpleHTTPServer()
		go http.BetterHTTPServer()
	}

	var wg sync.WaitGroup
	wg.Add(1)

	runTCP()
	runTCPHTTP()
	runHTTP()

	wg.Wait()
}
