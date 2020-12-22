package config

/**
IM 配置
*/
type IMConfig struct {
	IMPort     int `json:"im_port"`     //服务端长连接监听端口
	HttpPort   int `json:"http_port"`   //服务端短连接监听端口(登录接口)
	MaxClients int `json:"max_clients"` //服务端长连接最大连接数
}

/**
mysql 配置
*/
type MysqlConfig struct {
}

/**
redis 配置
*/
type redisConfig struct {
}
