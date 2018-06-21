package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

//SingleConnectionWriteServer - creates a simple TCP server that can be connected with telnet
func SingleConnectionWriteServer() {
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal(e)
	}

	var sessionID int

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		sessionID++

		io.WriteString(c, "Hello from simple TCP server\n")
		fmt.Fprintln(c, "Connection identifier for this connection is", sessionID)
		fmt.Fprintln(c, "Closing this connection now. Open another telnet sessoin for new connection")
		c.Close()
	}
}

//SingleConnectionReadServer - Reads input sent over tcp/http
func SingleConnectionReadServer() {
	l, e := net.Listen("tcp", ":8081")
	if e != nil {
		log.Fatal(e)
	}

	defer l.Close()

	for {
		c, e := l.Accept()
		if e != nil {
			log.Println(e)
			continue
		}

		s := bufio.NewScanner(c)
		for s.Scan() {
			log.Println(s.Text())
		}

		c.Close()
	}
}

//MultipleConnectionServer - creates a server that can accept multiple connections concurrently
func MultipleConnectionServer() {
	l, e := net.Listen("tcp", ":8082")
	if e != nil {
		log.Fatal(e)
	}

	defer l.Close()

	handler := func(c net.Conn) {
		defer c.Close()
		c.SetReadDeadline(time.Now().Add(10 * time.Second))

		sc := bufio.NewScanner(c)

		io.WriteString(c, "Hello from scalable TCP server!!\n")
		for sc.Scan() {
			ln := sc.Text()
			if strings.Compare(ln, "exit") == 0 {
				fmt.Fprintln(c, "Bye bye....")
				break
			}
			fmt.Fprintln(c, "You entered -", ln)
		}

		fmt.Fprintln(c, "Closing connection now")
	}

	for {
		c, e := l.Accept()
		if e != nil {
			log.Println(e)
			continue
		}

		go handler(c)
	}
}
