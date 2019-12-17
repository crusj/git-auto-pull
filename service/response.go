package service

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(data interface{}, msg ...string) []byte {
	return NewResponse(fasthttp.StatusOK, data, msg...)
}
func Error(data interface{}, msg ...string) []byte {
	return NewResponse(fasthttp.StatusBadRequest, data, msg...)
}

func NewResponse(code int, data interface{}, msg ...string) []byte {
	res, err := json.Marshal(
		&Response{
			Code: code,
			Data: data,
			Msg:  append(msg, "")[0],
		})
	if err != nil {
		panic(err)
	} else {
		return res
	}
}
