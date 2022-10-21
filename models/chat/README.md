# 聊天室
## server端
1、启动tcp服务 <br />
2、维护一个client自己接收消息channel <br />
3、维护一个了所有在线client的map, key: 每一个client的channel, value: true <br />
4、维护一个所有用户发消息的channel, 目的是监听该channel, 如果channel有消息, 将该消息写入到所有在线client维护的channel中 <br />
5、维护一个新进client的channel <br />
6、维护一个退出client的channel <br />
7、服务端开启一个协程, 来监听client的进入的channel、client的退出channel、所有用户发消息的channel <br />
    1）当有client进来, 发全局消息告知大家有人进来了, 同时维护全局map写数据(加锁) <br />
    2）当有client退出, 销毁tcp客户端连接, 销毁自己的channel, 发全局消息告诉大家有人退出了, 同时更新全局map <br />
    3）有全局消息写入, 则循环所有在线的map, 将该消息写入到所有client维护的channel中 <br />
8、客户端连接过来的时候, 服务端开启协程用来处理客户端的读写 <br />
    1）自己维护的channel有数据的时候, 服务端通过tcp的连接将数据发送给客户端 <br />
9、cd models/chat/server; go run main.go <br />

## client端
1、启动tcp客户端, 连接tcp的服务端 <br />
2、启动协程, 从tcp连接获取服务端发送的数据 <br />
2、往服务端发送数据, 监控现象 <br />
3、cd models/chat/client; go run main.go <br />