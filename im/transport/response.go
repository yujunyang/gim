package transport

import "encoding/json"

/**
输出对象
*/
type IMResponse struct {
	Status int         // 状态 0 成功，非 0 错误
	Msg    string      // 消息
	Data   interface{} // 数据
	Refer  string      // 来源
}

/*
错误消息构造方法
*/
func NewIMResponseSimple(status int, msg string, refer string) *IMResponse {
	return &IMResponse{status, msg, nil, refer}
}

/*
成功消息构造方法
*/
func NewIMResponseData(data interface{}, refer string) *IMResponse {
	return &IMResponse{0, "", data, refer}
}

/*
将返回消息转成JSON
*/
func (this *IMResponse) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

/*
将JSON转成返回消息
*/
func (this *IMResponse) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}
