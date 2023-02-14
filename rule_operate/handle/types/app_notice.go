package types

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/2 14:11
CREATE_BY:GoLand.LinEngineRules
*/

type AppNoticeMsg struct {
	TimesTamp int                    `json:"timestamp"`
	DevSN     string                 `json:"devSN"`
	Param     map[string]interface{} `json:"param"`
}

type ExprExample struct {
	Control map[string]interface{} `json:"control,omitempty"`
	Expr    string                 `json:"expr,omitempty"`
	RuleID  string                 `json:"rule_id,omitempty"`
}
