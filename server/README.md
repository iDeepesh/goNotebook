# Servers
- Listen
- Accept connection
- Data exchange
- Close connection

# tcp
- https://golang.org/pkg/net
    - TCP server
    - Client connection via telnet
    - Scanner for reading the input stream
    - Write into output stream
- SingleConnectionWriteServer
    - telnet localhost 8080
- SingleConnectionReadServer
    - telnet localhost 8081
    - Type anything and watch the server standard output
    - Open another session
        - telnet localhost 8080
        - Type something and watch the server standard output. Nothing will show
        - Close first telnet sesstion
        - Now the input from this session should output in server standard output
- MultipleConnectionServer
    - telnet localhost 8082
    - Type anything and server echos
    - Open another session at the same time
    - Type anything and server echos
    - Any session closes after 30 seconds on non-activity

# tcp http
- https://golang.org/pkg/net
    - TCP server
    - HTTP is on top of TCP so accepts HTTP connections as well
    - Custom code to handle HTTP payload
    - Client connection via telnet/curl
    - Scanner for reading the input stream
    - Write into output stream
- TcpHttpServer
    - curl localhost:9080/abc/xyz
        - Watch server standard output
        - Server echos to client
    - telnet localhost 9080
        - GET / HTTP/1.1
        - empty line
    - curl -X POST localhost:9080/xyz/abc -d "k1:v1"
        - Watch server standard output
        - Server echos to client
- BetterTcpHttpServer
	- Open browser and type localhost:9081/abc
    - curl localhost:9081/abc
    - telnet localhost 9081
        - GET /abc HTTP/1.1
        - empty line
    - curl -X POST localhost:9081/abc/efg -d "k1:v1"

# http
- https://golang.org/pkg/net/http
    - Includes all the handling for HTTP protocol
    - Practical usable package
    - https://golang.org/pkg/net/http/#Handler
        - Any objet that serves http request should implement this interface
    - https://golang.org/pkg/net/http/#ListenAndServe
        - Wraps the basic server steps of listen, accept and close for http compliant requests
        - Pass custom handler or nil for DefaultServeMux
    - https://golang.org/pkg/net/http/#HandleFunc
        - DefaultServeMux
        - Routing to map paths with handlers with DefaultServeMux
- SimpleHTTPServer
    - curl 'localhost:7080?k=v&k1=v1' -v
    - curl -X POST localhost:7080 -d 'k1=v1' -d 'k2=v2'
- BetterHTTPServer
    - curl 'localhost:7081?k=v&k1=v1' -v
    - curl -X POST localhost:7081 -d 'k1=v1' -d 'k2=v2' 

# file
- ServeFilesInManyWays
    - Uses IO stream to serve files
    - http://localhost:6080
- ServeWithFileServer
    - Uses FileServer to serve files in folders
    - http://localhost:6081/fs
- StaticFileServer
    - Quickly creates a static file server
    - Special behavior when index.html is present in a folder
    - http://localhost:6082/static