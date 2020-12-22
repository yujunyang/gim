package common

import (
	"gim/im/client"
	"gim/im/transport"
)

/*
客户端table
*/
type ClientTable map[string]*client.Client

/**
输入消息通道
*/
type InMessage chan transport.IMRequest

/**
输出消息通道
*/
type OutMessage chan transport.IMResponse
