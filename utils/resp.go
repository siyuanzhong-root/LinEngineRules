package utils

import (
	"LinEngineRules/constants"
	"LinEngineRules/types"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"time"
)

type errType struct {
	Msg string `json:"msg"`
}

type dataCntType struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

func RespOK(resp *restful.Response) {
	//resp.WriteHeader(http.StatusOK)
	_ = resp.WriteAsJson(&types.ResultVO{
		Code:      0,
		Timestamp: time.Now().UnixNano() / 1e6,
		Msg:       constants.ResponseStatusOK,
	})
}

func RespERR(resp *restful.Response) {
	//resp.WriteHeader(http.StatusOK)
	_ = resp.WriteAsJson(&types.ResultVO{
		Code:      400,
		Timestamp: time.Now().UnixNano() / 1e6,
		Msg:       constants.ResponseStatusError,
	})
}

func RespErrWithData(resp *restful.Response, msg string) {
	_ = resp.WriteAsJson(&types.ResultVO{
		Code:      http.StatusBadRequest,
		Timestamp: time.Now().UnixNano() / 1e6,
		Msg:       msg,
	})
}

func RespOKWithData(resp *restful.Response, data interface{}) {
	_ = resp.WriteAsJson(&types.ResultVO{
		Code:      0,
		Msg:       constants.ResponseStatusOK,
		Timestamp: time.Now().UnixNano() / 1e6,
		Data:      data,
	})
}

func RespWithDataAndCnt(resp *restful.Response, data interface{}, cnt int64) {
	_ = resp.WriteAsJson(&types.ResultVO{
		Code:      0,
		Msg:       constants.ResponseStatusOK,
		Timestamp: time.Now().UnixNano() / 1e6,
		Data: &dataCntType{
			Data:  data,
			Count: cnt,
		},
	})
}
