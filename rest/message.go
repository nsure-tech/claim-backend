package rest

import (
	"github.com/ethereum/go-ethereum/signer/core"
	"net/http"
)

type ResultMessage struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type MetamaskRequest struct {
	UserId string         `json:"user_id"`
	SigHex string         `json:"sig_hex"`
	Msg    core.TypedData `json:"msg"`
}

func newMessageError(code int, error error) *ResultMessage {
	return &ResultMessage{
		Code: code,
		Msg:  error.Error(),
	}
}

func newMessageBadRequest(error error) *ResultMessage {
	return &ResultMessage{
		Code: http.StatusBadRequest,
		Msg:  error.Error(),
	}
}

func newMessageInternalServerError(error error) *ResultMessage {
	return &ResultMessage{
		Code: http.StatusInternalServerError,
		Msg:  error.Error(),
	}
}
func newMessageOK() *ResultMessage {
	return &ResultMessage{
		Code: 0,
		Msg:  "OK",
	}
}

func newMessageDataOK(val interface{}) *ResultMessage {
	return &ResultMessage{
		Code: 0,
		Msg:  "OK",
		Data: val,
	}
}

func newMessageFail() *ResultMessage {
	return &ResultMessage{
		Code: -1,
		Msg:  "Fail",
	}
}

func newResultMessage(code int, msg, data string) *ResultMessage {
	return &ResultMessage{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
