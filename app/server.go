package main

import (
	redis "codecrafters-redis-go/redis"
	parser "codecrafters-redis-go/parser"
	utils "codecrafters-redis-go/utils"
	"log"
	"net"
	"fmt"
	str "strings"
	"bufio"
)

const HOST string = "127.0.0.1"
const PORT string = "6379"

func main() {
	server_addr := fmt.Sprintf("%s:%s", HOST, PORT)
	log.Printf("Starting a tcp server at %s \n", server_addr)

	server, err := net.Listen("tcp", server_addr)
	utils.CheckErr(err)
	defer server.Close()

	redis := redis.Redis{
		Storage: map[string]utils.RedisPair{},
	}
	for {
		conn, err := server.Accept()
		utils.CheckErr(err)
		log.Printf("New incoming connection %s \n", conn.RemoteAddr().String())
		go handleConnection(conn, &redis)
	}
}

func handleConnection(conn net.Conn, redis *redis.Redis) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	for {
		//var reply []byte = []byte{}
		//_, err := conn.Read(reply)
		//var b byte = '\n'
		//reply, _ := readUntilCRLF(bufio.NewReader(conn))
		// utils.CheckErr(err)
		// if err != nil {
		// 	continue
		// }

		log.Println("Talking to: ", addr)
		//log.Printf("Received: '%s'", string(reply))

		data, err := parser.ParseArray(bufio.NewReader(conn))
		utils.CheckErr(err)
		if err != nil && err.Error() == "EOF" {
			break
		} else if err != nil {
			continue
		}
		log.Printf("Received: %s", data)

		command := str.ToLower(data[0])
		var msg string 
		switch command {
		case "ping":
			msg, _ = redis.Ping()
		case "echo":
			msg, _ = redis.Echo(str.Join(data[1:], " "))
		case "set":
			expiry := ""
			if len(data) > 3 {
				expiry = data[4]
			}
			msg, _ = redis.Set(data[1], data[2], expiry)
		case "get":
			msg, _ = redis.Get(data[1])
		default:
			log.Println("ERROR - Uknown command:", string(command))
			msg = "-ERR : Unknow command" + string(command)
		}

		_, err = conn.Write([]byte(msg))
		utils.CheckErr(err)
		if err != nil {
			break
		}
		log.Println("Sent back: ", msg)
	}
	log.Println("Connection closed: ", addr)
}


