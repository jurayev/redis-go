package main

import (
	"log"
	"net"
	"fmt"
	str "strings"
	_ "strconv"
	_ "time"
	_ "codecrafters-redis-go/parser"
)

const HOST = "127.0.0.1"
const PORT = "6379"

func main() {
	server_addr := fmt.Sprintf("%s:%s", HOST, PORT)

	conn, err := net.Dial("tcp", server_addr)
	checkErr(err)
	defer conn.Close()
	log.Printf("Connected to a tcp: %s \n", server_addr)

	handleConnection(conn)

	log.Println("Connection closed")
}

func handleConnection(conn net.Conn) {
	// msg := "*1\r\n$4\r\nPING\r\n"
	// msg := "*2\r\n$4\r\nECHO\r\n$4\r\nHello\r\n"
	// msg := "*2\r\n$3\r\nGET\r\n$4\r\nHello2\r\n"
	// msg := "*3\r\n$3\r\nSET\r\n$4\r\nHello\r\n$4\r\nWorld\r\n"
	// msg := "*5\r\n$3\r\nSET\r\n$4\r\nHello1\r\n$4\r\nWorld1\r\n$2\r\nPX\r\n$2\r\n3600\r\n"
	log.Println("Send command? Y/N")
	var err error
	for {
		var input string
		_, err = fmt.Scanln(&input)
		checkErr(err)
		if str.ToLower(input) != "y" {
			break
		}
		msg := "*5\r\n$3\r\nSET\r\n$4\r\nHello1\r\n$4\r\nWorld1\r\n$2\r\nPX\r\n$2\r\n3600\r\n"
		_, err = conn.Write([]byte(msg))
		log.Println("Send:", msg)
		checkErr(err)

		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		checkErr(err)

		log.Println("Received:", string(reply))
		log.Println("Send command? Y/N")
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
