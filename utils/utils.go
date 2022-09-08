package utils
import "log"


func CheckErr(err error) {
	if err != nil {
		log.Println("WARNING :", err)
	}
}

type RedisPair struct {
	Value string
	Expiry int
}
