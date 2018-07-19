package main

import (
	"sync"

	"github.com/iDeepesh/goNotebook/server/file"
	"github.com/iDeepesh/goNotebook/server/http"
	"github.com/iDeepesh/goNotebook/server/state"
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

	runFile := func() {
		go file.ServeFilesInManyWays()
		go file.ServeWithFileServer()
		go file.StaticFileServer()
	}

	var wg sync.WaitGroup
	wg.Add(1)

	runTCP()
	runTCPHTTP()
	runHTTP()
	runFile()
	go state.ServerToParseInputs()

	wg.Wait()
}
