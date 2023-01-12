package rule_node

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/1/12 10:48
CREATE_BY:GoLand.LinEngineRules
*/

/**
定义基础接口结构和方法,统一不同类型需要实现的内容
*/

// Type 定义规则参数类型
type Type int

const (
	TypeInt64 Type = iota
	TypeStr
	TypeBool
	TypeBad
)

// ValueNode 定义valueNode相关方法
type ValueNode interface {
	GetType() Type
	GetTextValue() string
	//SetValue(string)
}

// Lit2ValueNode 定义生成valueNode方法
func Lit2ValueNode(lit *ast.BasicLit) ValueNode {
	switch lit.Kind {
	case token.INT:
		value, err := strconv.ParseInt(lit.Value, 10, 64)
		if err != nil {
			return NewBadNode(err.Error())
		}
		return NewIntNode(value)
	case token.STRING:
		value, err := strconv.Unquote(lit.Value)
		if err != nil {
			return NewBadNode(err.Error())
		}
		return NewStrNode(value)
	}

	return NewBadNode(fmt.Sprintf("%s is not support type", lit.Kind))
}
