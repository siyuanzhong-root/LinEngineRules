package constants

import (
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/handle/types"
	"os"
	"time"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/1 17:43
CREATE_BY:GoLand.LinEngineRules
*/
var (
	// MqttBroker mqtt地址
	MqttBroker = os.Getenv("MQTT_BROKER")
	// ExprRulePath 规则配置文件路径
	ExprRulePath = os.Getenv("EXPR_RULE_PATH")
	// ExprRuleFileName 规则配置文件名称
	ExprRuleFileName = "expr_rule.json"
	// RuleOperateInfo 规则配置信息
	RuleOperateInfo []types.RuleOperateElement
	// MQTTQos MQTT常量
	MQTTQos = byte(1)
	// MQTTListenALLTopic MQTT监听所有Topic
	MQTTListenALLTopic = "/v1/ai/data"
	// CacheDataOperate 缓存消息体消息
	CacheDataOperate []map[string]interface{}
	// RuleExchangeArray 缓存转换体消息
	RuleExchangeArray []map[string]interface{}
	// TimeSchedule 数据定时缓存任务
	TimeSchedule = time.NewTicker(time.Minute * 1)
	// ExprStrArray 规则字符串
	ExprStrArray []string
)
