package transport

import (
	"encoding/json"
	"gim/im/client"
)

/**
输入对象
*/

type IMRequest struct {
	Client   *client.Client
	Commoand string
	Data     map[string]map[string]string
}

/**
转成 JSON 数据
*/
func (this *IMRequest) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

/**
解析 JSON 数据
*/
func (this *IMRequest) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}

/**
解析 JSON 数据
*/
func DecodeIMRquest(data []byte) (*IMRequest, error) {
	req := new(IMRequest)

	err := req.Decode(data)

	if err != nil {
		return nil, err
	}
	return req, nil
}
