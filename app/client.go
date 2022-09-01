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
	
	conn, err := net.Dial("tcp", server_addr)
	checkErr(err)
	defer conn.Close()
	fmt.Printf("Connected to a tcp: %s \n", server_addr)

	handleConnection(conn)
	
	fmt.Println("Connection closed")
}

func handleConnection(conn net.Conn) {
	fmt.Println("Send command? Y/N")
	var err error
	for {
		var input string
		_, err = fmt.Scanln(&input)
		checkErr(err)
		if input != "y" {
			break
		}

		msg := "+PING"
		_, err = conn.Write([]byte(msg))
		fmt.Println("Send:", msg)
		checkErr(err)
	
		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		checkErr(err)

		fmt.Println("Received:", string(reply))
		fmt.Println("Send command? Y/N")
	}
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
