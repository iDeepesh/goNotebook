package http

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
	"time"
)

//SimpleHttpServer - Simple http server
func SimpleHttpServer() {
	handler := func(c net.Conn) {
		s := bufio.NewScanner(c)
		var method string
		var uri string
		var i int

		for s.Scan() {
			t := s.Text()
			fmt.Println(t)

			if i == 0 {
				sa := strings.Split(t, " ")
				method = sa[0]
				uri = sa[1]
			}

			if t == "" {
				break
			}

			i++
		}

		fmt.Println("The http method for request is:", method)
		fmt.Println("The URI for the request is:", uri)

		fmt.Fprintln(c, "Here is a response for", method, "request")

		c.Close()
	}

	l, e := net.Listen("tcp", ":9080")
	if e != nil {
		log.Fatal(e)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handler(c)
	}
}

//BetterHttpServer - server returns html respone to display in browser
func BetterHttpServer() {
	l, e := net.Listen("tcp", ":9081")
	if e != nil {
		log.Fatal(e)
	}

	for {
		c, er := l.Accept()
		if er != nil {
			log.Println(er)
			break
		}

		go httpServerHandler(c)
	}
}

func httpServerHandler(c net.Conn) {
	type data struct {
		Time   time.Time
		Method string
		Uri    string
	}
	sc := bufio.NewScanner(c)
	var i int
	var d data
	for sc.Scan() {
		s := sc.Text()
		fmt.Println(s)
		if i == 0 {
			sa := strings.Split(s, " ")
			d.Method = sa[0]
			d.Uri = sa[1]
			d.Time = time.Now()
		}
		i++
		if s == "" {
			break
		}
	}

	t, err := template.ParseFiles("http/httpServerResponse.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	var bodyBytes bytes.Buffer
	t.ExecuteTemplate(&bodyBytes, "httpServerResponse.gohtml", d)
	body := bodyBytes.String()

	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)
}
