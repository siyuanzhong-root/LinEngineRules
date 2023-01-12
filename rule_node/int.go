package rule_node

import "fmt"

/**
CREATE_USER:SYZ
CREATE_TIME:2023/1/12 10:46
CREATE_BY:GoLand.LinEngineRules
*/

// IntNode int类型结构体
type IntNode struct {
	Value     int64
	TextValue string
}

func (iNode IntNode) GetTextValue() string {
	return iNode.TextValue
}

//func (iNode IntNode) GetValue() int64 {
//	return iNode.Value
//}

func (iNode IntNode) GetType() Type {
	return TypeInt64
}

//func (iNode IntNode) SetValue(str string) {
//	iNode.TextValue = str
//}

func NewIntNode(value int64) ValueNode {
	textValue := fmt.Sprintf("%d", value)
	return IntNode{Value: value, TextValue: textValue}
}
