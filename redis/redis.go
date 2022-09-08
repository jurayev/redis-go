package redis

import (
	"strconv"
	utils "codecrafters-redis-go/utils"
	"time"
)


type Redis struct {
	Storage map[string]utils.RedisPair 
}

func (r *Redis) Ping() (string, error) {
	return "+PONG\r\n", nil
}

func (r *Redis) Echo(s string) (string, error) {
	return "+" + s + "\r\n", nil
}

func (r *Redis) Get(key string) (string, error) {
	pair, exists := r.Storage[key]
	var r_value string
	if !exists {
		r_value = "-ERR : No such key " + key
	} else if pair.Expiry != -1 && time.Now().UTC().Nanosecond() > pair.Expiry {
		delete(r.Storage, key)
		r_value =  "-ERR : No such key " + key
	} else {
		r_value = "+" + pair.Value
	}
	return r_value + "\r\n", nil
}

func (r *Redis) Set(key string, value string, ttl string) (string, error) {
	var expiry int = -1
	if ttl != "" {
		expiry_time, err := strconv.ParseInt(ttl, 10, 64)
		utils.CheckErr(err)
		if err != nil {
			return "-ERR : Expiry key is not a valid integer type", nil
		}
		expiry = time.Now().UTC().Add(time.Second * time.Duration(expiry_time)).Nanosecond()
		
	}
		
	r.Storage[key] = utils.RedisPair{Value: value, Expiry: expiry}
	return "+OK" + "\r\n", nil
}
