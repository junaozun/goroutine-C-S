package main

import (
	"fmt"
	"net"
	"strings"
)

func main()  {
	//指定服务器通信协议、IP地址、Port端口，创建一个用于监听的socket---listener
	listener ,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println("net.Listener err:",err)
		return
	}
	defer listener.Close()//关闭socket
	fmt.Println("服务器已启动...")

	//循环监听客户端连接请求
	for {
		conn ,err := listener.Accept()
		if err != nil{
			fmt.Println("listener.Accept err:",err)
			return
		}

		go HandleConnect(conn)
	}

}

func HandleConnect(conn net.Conn)  {
	defer conn.Close()
	//获取连接客户端的Addr
	addr := conn.RemoteAddr()
	fmt.Println(addr,"客户端连接成功")

    //循环读取客户端发送的数据
    buf := make([]byte,4096)
	for  {
		n ,err := conn.Read(buf)
		if err != nil{
			fmt.Println("conn.Read err:",err)
			return
		}

		if n == 0{
			fmt.Println("服务器检测到客户端已关闭，断开连接")
			return
		}
		if "exit\n"==string(buf[:n]){
			fmt.Println("服务器收到客户端退出请求，断开连接")
			return
		}
		fmt.Printf("服务器读到%s客户端发来的数据%s：",addr,string(buf[:n]))

		//小写转大写，回发给客户端
		S := strings.ToUpper(string(buf[:n]))
		conn.Write([]byte(S))
	}


}