package main

import (
	"fmt"
	"net"
	"log"
)

const HOST = "127.0.0.1"
const PORT = "6379"


func main() {
	server_addr := fmt.Sprintf("%s:%s", HOST, PORT)
	fmt.Printf("Starting a tcp server at %s \n", server_addr)

	server, err := net.Listen("tcp", server_addr)
	checkErr(err)
	defer server.Close()

	for {
		conn, err := server.Accept()
		checkErr(err)
		fmt.Printf("New incoming connection %s \n", conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	for {
		reply := make([]byte, 1024)
		var err error 
		_, err = conn.Read(reply)
		checkErr(err)
		if err != nil {
			break
		}

		fmt.Println("Talking to: ", addr)
		fmt.Println("Received: ", string(reply))

		msg := "+PONG"
		_, err = conn.Write([]byte(msg))
		checkErr(err)
		fmt.Println("Sent back: ", msg)
	}
	fmt.Println("Connection closed: ", addr)
}

func checkErr(err error) {
    if err != nil {
        log.Println(err)
    }
}
