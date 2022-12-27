package main

import (
	"fmt"
	"net/http"
	"github.com/go-redis/redis"
	"time"
)

var redisClient *redis.Client
var redisHost = "localhost:6379"
var redisPass = ""
var redisKey = "godemo:distribute_lock"

func logger(message string) {
	dat := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s, %s\n", dat, message)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisHost,
		Password: redisPass,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		resp := fmt.Sprintf("http request failed, redisClient.Ping().Result() error, err: %s\n", err.Error())
		logger(resp)
		w.Write([]byte(resp))
		return
	}

	b, err := redisClient.SetNX(redisKey, "", 1 * time.Second).Result()
	if !b || err != nil {
		resp := fmt.Sprintf("http request succ, lock failed, b: %v, err: %v\n", b, err)
		logger(resp)
		w.Write([]byte(resp))
		return
	} else {
		ret, delErr := redisClient.Del(redisKey).Result()
		resp := fmt.Sprintf("http request succ, lock succ, unlock: %d, b: %v, err: %v, unlockErr: %v\n", ret, b, err, delErr)
		logger(resp)
		w.Write([]byte(resp))
		return
	}
}

func main() {
	// 启动http服务, 监听url请求
	http.HandleFunc("/lock", IndexHandler)
	http.ListenAndServe(":8083", nil)
}