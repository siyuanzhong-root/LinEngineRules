package rule_expr

import (
	"errors"
	. "fmt"
	node "github.com/siyuanzhong-root/LinEngineRules/rule_node"
	"go/ast"
	"go/parser"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/1/31 16:02
CREATE_BY:GoLand.LinEngineRules
*/

// int string
// > < >= <= && ||
// in_arr(1, []int{1,2,3,4}), ver_compare(x, ">", "10.1.1") with no nested

// LogicEngine 逻辑引擎
type LogicEngine struct {
	ruleAst ast.Expr
}

// NewEngine 新增引擎
func NewEngine(expr string) (*LogicEngine, error) {
	engine := &LogicEngine{}
	result, err := engine.UpdateAst(expr)
	if err != nil || !result {
		return nil, err
	}
	return engine, nil
}

// UpdateAst 将规则转换str为语法树
func (engine *LogicEngine) UpdateAst(expr string) (bool, error) {
	//确保只能从NewEngine方法进入
	if engine == nil {
		panic("please init the engine first")
	}
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		Println(err)
		return false, err
	}
	engine.ruleAst = exprAst
	return true, nil
}

// RunRule 运行规则
func (engine *LogicEngine) RunRule(controlMap map[string]interface{}) (bool, error) {
	if engine == nil || engine.ruleAst == nil {
		return false, errors.New("rule expr is empty, please init it")
	}

	nodeMap := parseControlMap(controlMap)
	value := Eval(nodeMap, engine.ruleAst)
	bValue, ok := value.(node.BoolNode)
	if !ok {
		return false, errors.New(value.GetTextValue())
	}
	return bValue.True, nil
}

// parseControlMap 入参map转换类
func parseControlMap(controlMap map[string]interface{}) map[string]node.ValueNode {
	//新建map接受入参
	nodeMap := make(map[string]node.ValueNode, len(controlMap))
	for key, control := range controlMap {
		switch control.(type) {
		case int:
			node := node.NewIntNode(int64(control.(int)))
			nodeMap[key] = node
		case int64:
			node := node.NewIntNode(control.(int64))
			nodeMap[key] = node

		case float64:
			// value from json will be always float64
			node := node.NewFloatNode(control.(float64))
			nodeMap[key] = node

		case string:
			node := node.NewStrNode(control.(string))
			nodeMap[key] = node
		}
	}
	return nodeMap
}

func Eval(mem map[string]node.ValueNode, expr ast.Expr) (y node.ValueNode) {
	switch x := expr.(type) {
	//1个ast.BasicLit节点表示一个基本类型的文字，实现了ast.Expr接口
	//判断如果是BasicList类型,则获取该类型值
	case *ast.BasicLit:
		return node.Lit2ValueNode(x)
		//如果还是结构树则遍历
		//递归跳出树结构
	case *ast.BinaryExpr:
		a := Eval(mem, x.X)
		b := Eval(mem, x.Y)
		op := x.Op

		//如果a,b是空
		if a == nil || b == nil {
			return node.NewBadNode(Sprintf("%+v, %+v is nil", a, b))
		}

		switch a.GetType() {
		case node.TypeInt64:
			return BinaryIntExpr{}.Invoke(a, b, op)
		case node.TypeBool:
			return BinaryBoolExpr{}.Invoke(a, b, op)
		case node.TypeFloat:
			return BinaryFloatExpr{}.Invoke(a, b, op)
		case node.TypeStr:
			return BinaryStrExpr{}.Invoke(a, b, op)
		case node.TypeBad:
			return node.NewBadNode("a:" + a.GetTextValue() + "b:" + b.GetTextValue())
		}
		return node.NewBadNode(Sprintf("%d op is not suppoort", op))
	case *ast.CallExpr:
		name := x.Fun.(*ast.Ident).Name
		return CallExpr{name, x.Args}.Invoke(mem)
	case *ast.ParenExpr:
		return Eval(mem, x.X)
	case *ast.Ident:
		return mem[x.Name]
	default:
		return node.NewBadNode(Sprintf("%x type is not suppoort", x))
	}

	panic("internal error")
}
