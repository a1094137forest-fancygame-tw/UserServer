package constant

import (
	"encoding/json"

	"gopkg.in/olahol/melody.v1"
)

const (
	PING = 1000
)

type RequestMessage struct {
	MessageCode int         `json:"messageCode"`
	Data        interface{} `json:"data"`
}

func (req *RequestMessage) StringToByte() ([]byte, error) {
	msg, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func RequestWithData(s *melody.Session, messageCode int, data interface{}) error {

	req := RequestMessage{
		MessageCode: messageCode,
		Data:        data,
	}

	jsonData, _ := req.StringToByte()
	if err := s.Write(jsonData); err != nil {
		return err
	}

	return nil
}
