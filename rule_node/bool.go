package rule_node

import "fmt"

/**
CREATE_USER:SYZ
CREATE_TIME:2023/1/12 10:46
CREATE_BY:GoLand.LinEngineRules
*/

// BoolNode 布尔类型结构体
type BoolNode struct {
	textValue string
	True      bool
}

func (bNode BoolNode) GetTextValue() string {
	return bNode.textValue
}

func (bNode BoolNode) GetValue() bool {
	return bNode.True
}

func (bNode BoolNode) GetType() Type {
	return TypeBool
}

func NewBoolNode(b bool) ValueNode {
	return BoolNode{
		True:      b,
		textValue: fmt.Sprintf("%t", b),
	}
}
