package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"os"
	"golang.org/x/time/rate"
	"github.com/go-redis/redis"
)

// redis配置
var redisClient *redis.Client
var redisHost = "localhost:6379"
var redisPass = ""
var redisKeyHttpCounter = "godemo:ratelimit:http:counter:"
var redisKeyDistribute = "godemo:ratelimit:distribute"

// 每秒只允许200个请求
var ratelimit = 200

// 令牌桶
var rateLimiter *rate.Limiter

// 记录日志
func logger(ratelimit int, message string) {
	dat := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Printf("date[%s] ratelimit[%d] message[%s]\n", dat, ratelimit, message)
}

// 计数器限流
func CounterHandler(w http.ResponseWriter, r *http.Request) {
	// 当前放行的http请求的计数器的redis key
	suffix := time.Now().Format("20060102150405")
	redisKey := fmt.Sprintf("%s%s", redisKeyHttpCounter, suffix)

	// 分布式锁的redis key
	b, err := redisClient.SetNX(redisKeyDistribute, "", 1 * time.Second).Result()
	if err != nil {
		// redis操作失败, 不放行, 直接返回
		resp := fmt.Sprintf("redis operation error, err: %v", err)
		logger(0, resp)
		w.Write([]byte(resp))
		return
	} else if !b {
		// 没拿到分布式锁, 不放行, 直接返回
		resp := fmt.Sprintf("lock failed, b: %v", b)
		logger(0, resp)
		w.Write([]byte(resp))
		return
	} else {
		// 拿到分布式锁, 获取当前秒的http计数
		isExists, _ := redisClient.Exists(redisKey).Result()
		if isExists == 1 {
			// 存在, 则查看结果是否超过ratelimit
			counterStr, _ := redisClient.Get(redisKey).Result()
			counter, _ := strconv.Atoi(counterStr)
			resp := fmt.Sprintf("counter: %d", counter)
			if counter > ratelimit {
				// 当前秒超过限制, 不放行
				redisClient.Del(redisKeyDistribute).Result()
				logger(0, resp)
				w.Write([]byte(fmt.Sprintf("deny, %s", resp)))
				return
			} else {
				// 当前秒没超过限制, 放行, 计数器加1
				redisClient.Incr(redisKey).Result()
				redisClient.Del(redisKeyDistribute).Result()
				logger(1, resp)
				w.Write([]byte(fmt.Sprintf("allow, %s", resp)))
				return
			}
		} else {
			// 不存在, 则设置计数器为1, 过期时间10分钟
			redisClient.Set(redisKey, 1, 10 * time.Minute).Result()
			redisClient.Del(redisKeyDistribute).Result()
			resp := fmt.Sprintf("allow, counter: %d", 1)
			logger(1, resp)
			w.Write([]byte(resp))
			return
		}
	}
}

// 令牌桶算法
func BucketTokenHandler(w http.ResponseWriter, r *http.Request) {
	if rateLimiter.Allow() {
		// 桶里还有数, 放行. 则从桶里获取一个令牌
		resp := fmt.Sprintf("allow, size: %d", rateLimiter.Burst())
		logger(1, resp)
		w.Write([]byte(resp))
		return
	} else {
		// 桶里没有数, 不放行. 等待以后每秒自从往桶里放令牌
		resp := fmt.Sprintf("deny, size: %d", rateLimiter.Burst())
		logger(0, resp)
		w.Write([]byte(resp))
		return
	}

}

func main() {
	// 初始化redis链接
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisHost,
		Password: redisPass,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Printf("http request failed, redisClient.Ping().Result() error, err: %s\n", err.Error())
		os.Exit(1)
	}

	// 初始化令牌桶. 第一个参数: 每秒往桶中放多少令牌, 第二个参数: 桶的容量, go官方的实现中, 初始化后桶中的令牌是满的
	rateLimiter = rate.NewLimiter(200, ratelimit)

	// 启动http服务, 监听url请求
	http.HandleFunc("/ratelimit/counter", CounterHandler)
	http.HandleFunc("/ratelimit/bucket_token", BucketTokenHandler)
	http.ListenAndServe(":8083", nil)
}