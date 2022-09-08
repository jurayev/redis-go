package utils
import "log"


func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

type RedisPair struct {
	Value string
	Expiry int
}
