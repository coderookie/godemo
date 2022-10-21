package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// 所有client发消息都写入该全局channel
var messageChan = make(chan string, 1000)

// 每一个客户端维护自己的channel, 里面装的都是别人发送给它的信息
type clientChan chan string

// 新进来的client的channel
var clientInChan = make(chan clientChan, 100)

// 退出的client的channel
var clientOutChan = make(chan clientChan, 100)

// 所有在线client数据结构
type OnlineClientStruct struct {
	// 修改全局map的锁
	lock	sync.Mutex

	// 所有在线的client的map
	onlineClientMap map[clientChan]bool
}
var onlineClient = &OnlineClientStruct{}

// 初始化全局map
func init() {
	onlineClient.onlineClientMap = make(map[clientChan]bool)
}

// 获取当前时间
func getFormatDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 处理新的客户端
func HandleConn(conn net.Conn) {
	defer conn.Close()

	// 自己的channel
	client := make(clientChan)

	// 自己channel有消息, 返回客户端
	go handleClientMessage(conn, client)

	// 将自己加入到新进来的client中
	clientInChan <- client

	// 用户地址
	address := conn.RemoteAddr().String()
	message := fmt.Sprintf("%s %s: 进来了", getFormatDate(), address)
	messageChan <- message

	// 读取用户输入流, 写入全局消息channel
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChan <- fmt.Sprintf("%s %s: %s", getFormatDate(), address, input.Text())
	}

	// 用户不输入了, 关闭了(input.Scan() == false)
	clientOutChan <- client
	message = fmt.Sprintf("%s %s: 退出了", getFormatDate(), address)
}

// 每个client进来, 如果自己维护的channel有数据, 则返回给客户端, 每个client都在服务端有自己的协程控制
func handleClientMessage(conn net.Conn, ch clientChan) {
	for {
		select {
		case message := <-ch:
			conn.Write([]byte(message))
		}
	}
}

// 监听多个channel
func broadcaster() {
	for {
		select {
		case c := <-clientInChan:
			// 修改全局map
			onlineClient.lock.Lock()
			onlineClient.onlineClientMap[c] = true
			onlineClient.lock.Unlock()
		case c := <-clientOutChan:
			onlineClient.lock.Lock()
			delete(onlineClient.onlineClientMap, c)
			onlineClient.lock.Unlock()
			close(c)
		case message := <-messageChan:
			for k, _ := range onlineClient.onlineClientMap {
				k <- message
			}
			// 输出到server的屏幕上去
			fmt.Println(message)
		}
	}
}

func main() {
	// 启动tcp服务, 监听8082端口
	listener, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		fmt.Printf("chat server, net.Listen error, error[%s]\n", err.Error())
		os.Exit(1)
	}

	// 服务启动后, 开启协程, 监听多个channel
	go broadcaster()

	// 服务端监听所有的请求, 非阻塞
	for {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			fmt.Printf("chat server, listener.Accept error, error[%s]\n", err.Error())
			continue
		}

		// 每个tcp连接开启协程
		go HandleConn(conn)
	}
}