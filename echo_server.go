/*
write an echo TCP server
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

		conn.Write([]byte(string("Welcome to benny's echo service\nWe echo!\n")))
		conn.Write([]byte(string("Enter your string!\n")))
		conn.Write([]byte(string("Enter STOP to end the session\n")))
		conn.Write([]byte(string("Enter PANIC to end the program in panic\n")))
		// handle the request received in a goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println(conn.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		if temp == "PANIC" {
			panic(fmt.Sprint("Panic!"))
		}

		fmt.Println("got ", temp," replying")

		conn.Write([]byte(string("benny says: " + temp+"\n")))
	}
	conn.Close()
}