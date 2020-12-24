package server

import (
	"gim/config"
	"gim/im/client"
	"gim/im/common"
	"gim/im/model"
	"log"
	"net"
)

/**
服务端server结构体
*/
type Server struct {
	listener    net.Listener        // 服务端监听器
	clients     common.ClientTable  // 客户端列表
	joinsniffer chan net.Conn       // 连接接入
	quitsniff   chan *client.Client // 连接退出
	insniffer   common.InMessage    // 接受消息
}

var ClientMaps common.ClientTable

/**
IM 服务启动
*/
func starIMServer(config config.IMConfig) {
	log.Fatal("IM服务端启动")
	// 初始化服务端
	sever := &Server{
		clients:     make(common.ClientTable, model.Config.MaxClients),
		joinsniffer: make(chan net.Conn),
		quitsniff:   make(chan *client.Client),
		insniffer:   make(common.InMessage),
	}

	// 启动监听
	sever.listen()

}

/**
监听
*/
func (this *Server) listen() {
	go func() {
		for {
			select {
			// 接入连接
			//case conn := <-this.joinsniffer:
			//	{
			//	}
			//case

			}

		}
	}()
}

/**
建立连接
*/
func (this *Server) joinHandler(conn net.Conn) {
	// 客户端 key
	key := "123"
	// 创建一个客户端
	client := client.CreateClient(key, conn)
	// 给客户端指定key
	this.clients[key] = client
	log.Printf("新客户端Key:[%s] online:%d", client.Key, len(ClientMaps))

	// 开启协程不断接受消息
	go func() {

	}()

}
