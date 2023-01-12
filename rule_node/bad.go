package rule_node

/**
CREATE_USER:SYZ
CREATE_TIME:2023/1/12 10:47
CREATE_BY:GoLand.LinEngineRules
*/

// BadNode 错误消息结构体
type BadNode struct {
	ErrorMessage string
}

func (bNode BadNode) GetTextValue() string {
	return bNode.ErrorMessage
}

func (bNode BadNode) GetType() Type {
	return TypeBad
}

func NewBadNode(str string) ValueNode {
	return BadNode{
		ErrorMessage: str,
	}
}
