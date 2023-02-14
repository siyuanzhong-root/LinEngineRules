package main

import (
	"fmt"
	"github.com/siyuanzhong-root/LinEngineRules/rule_expr"
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/constants"
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/handle/execute"
	_ "github.com/siyuanzhong-root/LinEngineRules/rule_operate/handle/init"
	"log"
	"time"
)

var exprExampleList = []struct {
	control map[string]interface{}
	expr    string
}{
	{
		map[string]interface{}{"pType": 36.5, "pid": 317, "pet": 252, "rule": 1},
		`pType > 36.12 && pid > 310 && pet >251`,
	},

	{
		map[string]interface{}{"pType": 66, "pid": 317},
		`pType < 90 && pid == 310`,
	},

	{
		map[string]interface{}{"version": "11.0.1", "pid": 317},
		`ver_compare(version, ">", "10.1.1")`,
	},

	{
		map[string]interface{}{"version": "11.0.1", "pid": 317},
		`ver_compare(version, "<", "10.1.1")`,
	},

	{
		map[string]interface{}{"pid": 317},
		`in_array(pid, []int{317, 318, 319})`,
	},

	{
		map[string]interface{}{"pid": "317"},
		`in_array(pid, []string{"317", "318", "319"})`,
	},
}

func main() {
	fmt.Println(constants.RuleExchangeArray)
	fmt.Println(constants.ExprStrArray)
	execute.DriverEngine()
	time.Sleep(time.Minute * 2)
	for _, example := range exprExampleList {
		log.Println("执行了一次,规则数据是：", example.control, "表达式是：", example.expr)
		engine, err := rule_expr.NewEngine(example.expr)
		if err != nil {
			fmt.Println("rules has error", err.Error())
		}
		result, err := engine.RunRule(example.control)
		if err != nil {
			fmt.Println("error has occurred", err.Error())
		}
		log.Println("规则执行结果", result)
	}

}
