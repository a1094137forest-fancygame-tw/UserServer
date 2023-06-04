package constant

import (
	"encoding/json"

	"gopkg.in/olahol/melody.v1"
)

const (
	SUCCESS = 200
	ERROR   = 500
)

const (
	SUCCESS_MSG = "Ok"
	ERROR_MSG   = "Fail"
)

type ResponseMessage struct {
	ReturnCode int         `json:"return_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func (resp *ResponseMessage) StringToByte() ([]byte, error) {
	msg, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func ResponseWithData(s *melody.Session, returnCode int, msg string, data ...interface{}) {
	resp := ResponseMessage{
		ReturnCode: returnCode,
		Msg:        msg,
		Data:       data,
	}

	jsonMsg, _ := resp.StringToByte()

	_ = s.Write(jsonMsg)
}
