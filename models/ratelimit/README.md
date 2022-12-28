# 限流策略
1、实现一个http服务器, 监听请求, 使用ab压测来验证不同策略的限流情况 <br />
2、ab -c 100 -n 100000 'http://localhost:8083/ratelimit/counter', 压测计数器限流 <br />
3、ab -c 100 -n 100000 'http://localhost:8083/ratelimit/bucket_token', 压测令牌桶限流 <br />