package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {

	//指定服务器IP+port，创建通信socket--conn
	conn,err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println("net.Dial err:",err)
		return
	}
	defer conn.Close()

	//获取用户键盘输入，将数据发送给服务器
	go func() {
		buf := make([]byte,4096)//将键盘输入存入buf缓冲区
		for {
			n ,err := os.Stdin.Read(buf)//读取键盘输入
			if err != nil{
				fmt.Println("os.Stdin.Read err:",err)
				continue
			}
			//将键盘读取的数据发送给服务器
			conn.Write(buf[:n])
		}

	}()

	//读服务器回发的大写数据
	buffer :=make([]byte,4096)
	for  {
		n,err := conn.Read(buffer)//从conn连接中读取数据，将读取出来的数据存入buffer中
		if err != nil{
			fmt.Println("conn.Read err:",err)
			return
		}
		fmt.Println("读到服务器回发的数据：",string(buffer[:n]))

	}


}
