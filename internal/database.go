package internal

import (
	"fmt"
	"os"
	"github.com/gomodule/redigo/redis"
)

func RedisConnection() redis.Conn {
	const IPPort = "127.0.0.1:6379"

	rc, err := redis.Dial("tcp", IPPort)
	if err != nil {
		panic(err)
	}
	return rc
}

func RedisSet(key string, value string, rc redis.Conn) {
	rc.Do("SET", key, value)
}

func RedisSetList(key string, value []string, rc redis.Conn) {
	for _, v := range value {
		fmt.Println(v)
		rc.Do("RPUSH", key, v)
	}
}

func RedisGet(key string, rc redis.Conn) string {
	s, err := redis.String(rc.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func RedisGetList(key string, rc redis.Conn) []string {
	s, err := redis.Strings(rc.Do("LRANGE", key, 0, -1))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

// func main(){
// 	rc := internal.RedisConnection()

// 	defer rc.Close()

// 	key := "KEY"
// 	val := "VALUE"
// 	internal.RedisSet(key, val, rc)
// 	s := internal.RedisGet(key, rc)
// 	fmt.Println(s)

// 	key = "LIST"
// 	vallist := []string{"VALUE1", "VALUE2", "VALUE3"}
// 	internal.RedisSetList(key, vallist, rc)
// 	sl := internal.RedisGetList(key, rc)
// 	fmt.Println(sl)
// 	fmt.Println(sl[0])
// 	fmt.Println(sl[1])
// 	fmt.Println(sl[2])
// }