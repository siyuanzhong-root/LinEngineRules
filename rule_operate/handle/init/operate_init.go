package init

import (
	"encoding/json"
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/constants"
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/handle/types"
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/handle/util"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/1 17:33
CREATE_BY:GoLand.LinEngineRules
*/

func init() {
	// 判断引擎所需事件源和规则配置文件路径是否已经指定好
	if constants.MqttBroker == "" || constants.ExprRulePath == "" {
		log.Fatalln("linEngineRules init failed,please operate env 'MQTT_BROKER' or assigner 'expr_rule.json'")
	}
	// 判断指定目录内规则配置文件是否存在
	exists, err := PathExists(path.Join(constants.ExprRulePath, constants.ExprRuleFileName))
	if err != nil {
		log.Println("query file exist has error", err.Error())
		return
	}
	if exists == false {
		_, err := os.Create(path.Join(constants.ExprRulePath, constants.ExprRuleFileName))
		if err != nil {
			return
		}
	}
	ReadFileInfo()

}

// ReadFileInfo 初始化读取规则信息
func ReadFileInfo() {
	bytes, err := ioutil.ReadFile(path.Join(constants.ExprRulePath, constants.ExprRuleFileName))
	if err != nil {
		log.Println("read expr_rule.json error")
		return
	}
	if bytes != nil {
		err := json.Unmarshal(bytes, &constants.RuleOperateInfo)
		if err != nil {
			return
		}
	}
	for _, ruleOperate := range constants.RuleOperateInfo {
		for _, detail := range ruleOperate.RuleDetail {
			var wg sync.WaitGroup
			var wgNum int
			var ruleMap = make(map[string]interface{})
			var cacheMap = make(map[string]interface{})
			ruleMap["rule_id"] = ruleOperate.RuleID
			cacheMap["rule_id"] = ruleOperate.RuleID
			for _, rule := range constants.RuleOperateInfo {
				wgNum = len(rule.RuleDetail)
			}
			wg.Add(wgNum)
			go func(ruleDetail types.RuleDetail) {
				source := util.HandleStrToArrayBy0(ruleDetail.RuleSource)[0]
				deviceSN := util.HandleStrToArrayBy0(ruleDetail.RuleSource)[1]
				attribute := ruleDetail.RuleAttr
				for _, attr := range attribute {
					ruleMap[source+"-"+deviceSN+"-"+attr] = ""
					cacheMap[attr] = ""
				}
				constants.RuleExchangeArray = append(constants.RuleExchangeArray, ruleMap)
				constants.CacheDataOperate = append(constants.CacheDataOperate, cacheMap)
				wg.Done()
			}(detail)
			wg.Wait()
		}
		constants.ExprStrArray = append(constants.ExprStrArray, ruleOperate.RuleExpr)
	}

	//constants.CacheDataOperate = append(constants.CacheDataOperate, ruleMap)
}

// PathExists 判断一个文件或文件夹是否存在 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
