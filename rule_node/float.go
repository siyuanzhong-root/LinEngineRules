package rule_node

import (
	"fmt"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/1/12 10:46
CREATE_BY:GoLand.LinEngineRules
*/

// FloatNode Float类型结构体
type FloatNode struct {
	Value     float64
	TextValue string
}

func (iNode FloatNode) GetTextValue() string {
	return iNode.TextValue
}

//func (iNode IntNode) GetValue() int64 {
//	return iNode.Value
//}

func (iNode FloatNode) GetType() Type {
	return TypeFloat
}

//func (iNode IntNode) SetValue(str string) {
//	iNode.TextValue = str
//}

func NewFloatNode(value float64) ValueNode {
	textValue := fmt.Sprintf("%v", value)
	return FloatNode{Value: value, TextValue: textValue}
}
