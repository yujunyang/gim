package client

import (
	"bufio"
	"gim/im/common"
	"gim/im/model"
	"gim/im/transport"
	"log"
	"net"
)

/*
	客户端结构体
*/
type Client struct {
	Key    string            // 客户端；连接唯一标识
	Conn   net.Conn          // 连接
	In     common.InMessage  // 输入
	Out    common.OutMessage // 输出
	Quit   chan *Client
	reader *bufio.Reader // 读取
	writer *bufio.Writer // 输出
	Login  *model.Login  // 登陆信息
}

/**
创建客户端
*/
func CreateClient(key string, conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	client := &Client{
		Key:    key,
		Conn:   conn,
		In:     make(common.InMessage),
		Out:    make(common.OutMessage),
		Quit:   make(chan *Client),
		reader: reader,
		writer: writer,
	}

	return client
}

/**
自动读入或者写出消息
*/
func (this *Client) Listen() {
	go this.read()
	go this.write()
}

/**
读消息
*/
func (this *Client) read() {
	for {
		if line, _, err := this.reader.ReadLine(); err == nil {
			req, err := transport.DecodeIMRquest(line)
			if err == nil {
				req.Client = this
				this.In <- *req
			} else {
				log.Printf("解析JSON错误：%s", line)
			}
		} else {
			this.Quiting()
			return
		}
	}
}

/**
写消息
*/
func (this *Client) write() {
	for resp := range this.Out {
		if _, err := this.writer.WriteString(string(resp.Encode()) + "\n"); err != nil {
			this.Quiting()
			return
		}
		if err := this.writer.Flush(); err != nil {
			this.Quiting()
			return
		}
	}
}

/**
退出
*/
func (this *Client) Quiting() {
	this.Quit <- this
}

/**
关闭连接
*/
func (this *Client) Close() {
	this.Conn.Close()
}
