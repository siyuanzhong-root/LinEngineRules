package types

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/4 11:37
CREATE_BY:GoLand.LinEngineRules
*/

type Task struct {
	RuleID          string   `json:"rule_id"`
	TopicList       []string `json:"topic_list"`
	DeviceList      []string `json:"device_list"`
	RuleDataExample chan ExprExample
}
