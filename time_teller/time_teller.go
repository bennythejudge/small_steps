/*
A service that tells time

Remember to test using

nc localhost 3333 (don't use curl)
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// listen for incoming connections
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
	}
	// close the listener when the app closes
	defer l.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	// wait indefinitely listeninig for incoming connections
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error executing l.Accept(): ", err.Error())
			os.Exit(1)
		}

		// handle the request received in a goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println(conn.RemoteAddr().String())
	conn.Write([]byte(fmt.Sprintf("%d\n", time.Now().UnixNano() / int64(time.Millisecond))))
	conn.Close()
}