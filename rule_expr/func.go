package rule_expr

import (
	"fmt"
	"github.com/siyuanzhong-root/LinEngineRules/rule_method"
	node "github.com/siyuanzhong-root/LinEngineRules/rule_node"
	"go/ast"
	"go/token"
	"strings"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/1 10:41
CREATE_BY:GoLand.LinEngineRules
*/

type BinaryBoolExpr struct{}

type BinaryStrExpr struct{}

type BinaryIntExpr struct{}

type BinaryFloatExpr struct{}

// CallExpr 内置方法比较
type CallExpr struct {
	fn   string // one of "in_array", "ver_compare"
	args []ast.Expr
}

// Invoke bool类型变量扩展方法
func (b BinaryBoolExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
	//判断x,y是否实现了node.BoolNode类型
	xb, xok := x.(node.BoolNode)
	yb, yok := y.(node.BoolNode)
	//如果有未实现的变量则返回错误信息和字段值
	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}
	// 判断token符号
	switch op {
	//如果是&&
	case token.LAND:
		//返回布尔比较结果
		return node.NewBoolNode(xb.True && yb.True)
		//如果是||
	case token.LOR:
		//比较比较结果
		return node.NewBoolNode(xb.True || yb.True)
	}
	return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
}

// Invoke string类型变量扩展方法
func (b BinaryStrExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
	xs, xok := x.(node.StrNode)
	ys, yok := y.(node.StrNode)
	//如果有未实现的变量则返回错误信息和字段值
	if !xok || !yok {
		return node.NewBadNode("x: " + x.GetTextValue() + "y: " + y.GetTextValue())
	}

	switch op {
	//判断xs,ys是否相等
	case token.EQL: // ==
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) == 0)
	case token.LSS: // <
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) == -1)
	case token.GTR: // >
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) == +1)
	case token.GEQ: // >=
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) >= 0)
	case token.LEQ: // <=
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) <= 0)
	}
	return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
}

// Invoke float类型变量扩展方法
func (b BinaryFloatExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
	xs, xok := x.(node.FloatNode)
	ys, yok := y.(node.FloatNode)
	//判断是否为intNode节点类型
	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}

	switch op {
	case token.EQL: // ==
		return node.BoolNode{
			True: xs.Value == ys.Value,
		}
	case token.LSS: // <
		return node.BoolNode{
			True: xs.Value < ys.Value,
		}
	case token.GTR: // >
		return node.NewBoolNode(xs.Value > ys.Value)
	case token.GEQ: // >=
		return node.NewBoolNode(xs.Value >= ys.Value)
	case token.LEQ: // <=
		return node.NewBoolNode(xs.Value <= ys.Value)
	}
	return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
}

// Invoke int类型变量扩展方法
func (b BinaryIntExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
	xs, xok := x.(node.IntNode)
	ys, yok := y.(node.IntNode)
	//判断是否为intNode节点类型
	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}

	switch op {
	case token.EQL: // ==
		return node.BoolNode{
			True: xs.Value == ys.Value,
		}
	case token.LSS: // <
		return node.BoolNode{
			True: xs.Value < ys.Value,
		}
	case token.GTR: // >
		return node.NewBoolNode(xs.Value > ys.Value)
	case token.GEQ: // >=
		return node.NewBoolNode(xs.Value >= ys.Value)
	case token.LEQ: // <=
		return node.NewBoolNode(xs.Value <= ys.Value)
	}
	return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
}

//Invoke 判断是否使用内置方法
func (c CallExpr) Invoke(mem map[string]node.ValueNode) node.ValueNode {
	switch c.fn {
	case "in_array":
		parm := Eval(mem, c.args[0])
		if parm.GetType() == node.TypeBad {
			return parm
		}
		vRange, ok := c.args[1].(*ast.CompositeLit)
		if !ok {
			return node.NewBadNode("func in_array 2ed params is not a composite lit")
		}
		eltNodes := make([]node.ValueNode, 0, len(vRange.Elts))
		for _, p := range vRange.Elts {
			elt := Eval(mem, p)
			eltNodes = append(eltNodes, elt)
		}

		has := false
		for _, v := range eltNodes {
			if v.GetType() == parm.GetType() && v.GetTextValue() == parm.GetTextValue() {
				has = true
			}
		}

		return node.NewBoolNode(has)
	case "ver_compare":
		if len(c.args) != 3 {
			return node.NewBadNode("func ver_compare doesn't have enough params")
		}

		args := make([]string, 0, 3)
		for _, v := range c.args {
			arg := Eval(mem, v)
			if arg.GetType() != node.TypeStr {
				return node.NewBadNode("func ver_compare params type error")
			}
			args = append(args, arg.GetTextValue())
		}

		ret := rule_method.VersionCompare(args[0], args[1], args[2])
		return node.NewBoolNode(ret)
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
