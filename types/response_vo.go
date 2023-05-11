package types

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/18 17:36
CREATE_BY:GoLand.LinEngineRules
*/

// ResultVO http返回结构体
type ResultVO struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}
