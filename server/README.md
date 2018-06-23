# tcp
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

# http
- SimpleHttpServer
    - curl localhost:9080/abc/xyz
        - Watch server standard output
        - Server echos to client
    - curl -X POST localhost:9080/xyz/abc -d "k1:v1"
        - Watch server standard output
        - Server echos to client
- BetterHttpServer
	- Open browser and type localhost:9081/abc
    - curl -X POST localhost:9081/abc/efg -d "k1:v1"