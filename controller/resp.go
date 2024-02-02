package controller

import "fold/protobuf/golang"

type Response struct {
	ErrorCode int         `json:"errorCode,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

var errCodeMap = map[golang.ErrorCode]int{
	golang.ErrorCode_EC_INTERNAL_SERVER_ERROR: 500,
}
