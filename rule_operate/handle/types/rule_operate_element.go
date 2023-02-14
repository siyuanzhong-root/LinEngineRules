package types

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/1 17:31
CREATE_BY:GoLand.LinEngineRules
*/

// RuleOperateElement 交互模型
type RuleOperateElement struct {
	RuleDetail []RuleDetail `json:"rule_detail,omitempty"`
	RuleExpr   string       `json:"rule_expr,omitempty"`
	RuleHandle []string     `json:"rule_handle,omitempty"`
	RuleID     string       `json:"rule_id,omitempty"`
	RuleName   string       `json:"rule_name,omitempty"`
	RuleStatus string       `json:"rule_status,omitempty"`
}

type RuleDetail struct {
	RuleAttr   []string `json:"rule_attr,omitempty"`
	RuleSource string   `json:"rule_source,omitempty"`
}
