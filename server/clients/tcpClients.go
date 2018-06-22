package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	singleConnectionWriteServerClient()
	singleConnectionReadServerClient()
}

func singleConnectionWriteServerClient() {
	c, e := net.Dial("tcp", "localhost:8080")
	if e != nil {
		log.Fatal(e)
	}

	bs, err := ioutil.ReadAll(c)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs[:len(bs)]))

	c.Close()
}

func singleConnectionReadServerClient() {
	c, e := net.Dial("tcp", "localhost:8081")
	if e != nil {
		log.Fatal(e)
	}

	fmt.Fprintln(c, "Test message from singleConnectionReadServerClient")

	c.Close()
}
