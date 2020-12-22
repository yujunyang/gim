package client

import (
	"bufio"
	"gim/im/common"
	"gim/im/model"
	"net"
)

/*
	客户端结构体
*/
type Client struct {
	key    string            // 客户端；连接唯一标识
	conn   net.Conn          // 连接
	in     common.InMessage  // 输入
	out    common.OutMessage // 输出
	quit   chan *Client
	reader *bufio.Reader // 读取
	writer *bufio.Writer // 输出
	Login  *model.Login  // 登陆信息
}
