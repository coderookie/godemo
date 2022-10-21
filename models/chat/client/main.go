package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 发送数据给服务端
func write(conn net.Conn) {
	// 从命令行读取数据
	reader := bufio.NewReader(os.Stdin)

	for {
		// 换行打散
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("chat client, reader.ReadString error, error[%s]\n", err.Error())
			continue
		}

		_, writeErr := conn.Write([]byte(input))
		if writeErr != nil {
			fmt.Printf("chat client, conn.Write error, error[%s]\n", writeErr.Error())
			break
		}
	}
}

// 从服务端读取数据
func read(conn net.Conn) {
	for {
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil {
			fmt.Printf("chat client, conn.Read error, error[%s]\n", err.Error())
			break
		}

		message := string(b[:n])
		fmt.Println(message)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		fmt.Printf("chat client, net.Dial error, error[%s]\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// 将服务端数据返回到客户端中
	go read(conn)

	// 客户端的数据写入到服务单
	write(conn)
}