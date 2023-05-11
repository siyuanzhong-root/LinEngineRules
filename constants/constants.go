package constants

import (
	"LinEngineRules/model"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/1 17:43
CREATE_BY:GoLand.LinEngineRules
*/
var (
	// RuleDataSourceArray 数据源信息
	RuleDataSourceArray []model.RuleDataSource
	// ControlDeviceSendJson 控制报文格式
	ControlDeviceSendJson = `{
	  "mid":1000000000020034,
	  "type":"CMD_SERVICE",
	  "deviceId":"%v",
	  "timestamp":%d,
	  "expire":-1,
	  "param":{
		"nodeId": "%v",
		"cmd":"%v",
		"paras":{
			"body":%v
		}
	  }
	}`
	// ControlCommand 控制常量
	ControlCommand = "analog_Get"
	// LoraDeviceControlTopic lora设备控制主题
	LoraDeviceControlTopic = "/v1/loraapp/service/command"
	// BleDeviceControlTopic ble设备控制主题
	BleDeviceControlTopic = "/v1/ble/service/command"
)

const (
	// MQTTQosCount MQTT常量
	MQTTQosCount = byte(1)
	// ResponseStatusOK 请求成功返回响应值
	ResponseStatusOK = "ok"
	// ResponseStatusError 请求失败响应结果
	ResponseStatusError = "error"
	// IsAsyncRuleSymbol 是异步规则的标识
	IsAsyncRuleSymbol = 1
	// NotAsyncRuleSymbol 是异步规则的标识
	NotAsyncRuleSymbol = 0
)

// 接口API
const (
	// DetailTag 规则接口tag
	DetailTag = "规则详情信息"
	// DataSourceTag 数据源接口Tag
	DataSourceTag = "数据源详细信息"
	// ExprRuleRecordTag 规则记录信息
	ExprRuleRecordTag = "规则记录信息"
)

const (
	// DirectControlType 直接控制
	DirectControlType = "control"
	// DelayControlType 延时控制
	DelayControlType = "delay"
	// StatusEnable 状态启用
	StatusEnable = "1"
	// StatusDisable 状态启用
	StatusDisable = "0"
	// ExecuteFailedResult 条件触发规则失败
	ExecuteFailedResult = 0
	// ExecuteSuccessResult 条件触发规则成功
	ExecuteSuccessResult = 1
	// ControlFailedResult 执行规则失败
	ControlFailedResult = 0
	// ControlSuccessResult 执行规则成功
	ControlSuccessResult = 1
	// SingleRuleExpr 单条规则数
	SingleRuleExpr = 1
	// MultiRuleExpr 单条规则数
	MultiRuleExpr = 3
	// BasicRuleStarter 数据基础出发点
	BasicRuleStarter = 0
)
