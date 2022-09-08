package parser

import (
	e "errors"
	_ "strings"
	utils "codecrafters-redis-go/utils"
	"strconv"
	"bufio"
	"fmt"
)

const ARRAY byte = '*'

func ParseArray(byteStream *bufio.Reader) ([]string, error) {
	// "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"
	// returns ["ECHO", "hey"]
	// 1st element is the command name
	// rest elements are data
	bytes, err := readUntilCRLF(byteStream)
	if err != nil {
		return []string{}, err
	}

	if bytes[0] != ARRAY {
		return []string{}, e.New("invalid encoding type. Not an array")
	}
	fmt.Println("Bytes", bytes)

	//splitted := str.Split(string(input), "\r\n")
	
	//total, _ := strconv.ParseInt(splitted[0][1:], 10, 64)

	//data = append(data, splitted[2])
	items_count, err := strconv.Atoi(string(bytes[1]))
	fmt.Println("Items", items_count)
	utils.CheckErr(err)

	var data []string
	for i := 0; i < items_count; i++ {
		readUntilCRLF(byteStream) // simply skip '$4' bytes size data slice
		nextBytes, err := readUntilCRLF(byteStream)
		utils.CheckErr(err)
		data = append(data, string(nextBytes))
		fmt.Println("Str", string(nextBytes))
	}

	return data, nil
}

func readUntilCRLF(byteStream *bufio.Reader) ([]byte, error) {
	readBytes := []byte{}
	for {
		b, err := byteStream.ReadBytes('\n')
		if err != nil {
			return nil, err
		}
		readBytes = append(readBytes, b...)
		if len(readBytes) >= 2 && readBytes[len(readBytes)-2] == '\r' {
			break
		}
	}
	return readBytes[:len(readBytes)-2], nil
}