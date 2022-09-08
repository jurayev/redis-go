package parser

import (
	err "errors"
	str "strings"
	"strconv"
)

const ARRAY byte = '*'

func ParseArray(input []byte) ([]string, error) {
	// "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"
	// returns ["ECHO", "hey"]
	// 1st element is the command name
	// rest elements are data
	var data []string

	if input[0] != ARRAY {
		return []string{}, err.New("invalid encoding type. Not an array")
	}

	splitted := str.Split(string(input), "\r\n")
	
	total, _ := strconv.ParseInt(splitted[0][1:], 10, 64)

	data = append(data, splitted[2])

	for i := 4; int64(i) <= total*2; i = i + 2 {
		data = append(data, splitted[i])
	}

	return data, nil
}