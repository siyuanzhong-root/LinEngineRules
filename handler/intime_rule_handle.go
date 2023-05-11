package handler

import (
	"LinEngineRules/constants"
	"LinEngineRules/model"
	"LinEngineRules/types"
	"LinEngineRules/utils"
	"fmt"
	"log"
	"strings"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/20 10:57
CREATE_BY:GoLand.LinEngineRules
*/

// InTimeHandle 同步处理器控制
func InTimeHandle(detail model.RuleDetail, msg types.AppNoticeMsg, topic string) {
	ruleArray := utils.HandleReturnFullResult(detail.SourceDeviceAttr)
	for i := range ruleArray {
		if !dataIfInRule(ruleArray[i], msg, topic) {
			return
		}
	}
	var flag bool
	if len(ruleArray) == constants.SingleRuleExpr {
		log.Println("进入单条件方法", ruleArray[0], msg, topic)
		flag = inTimeExecute(ruleArray[0], msg, topic)
	}
	if len(ruleArray) >= constants.MultiRuleExpr {
		for index := range ruleArray {
			if index%2 == 0 && index < len(ruleArray)-2 {
				preRule := ruleArray[index]
				nextRule := ruleArray[index+2]
				preRuleResult := inTimeExecute(preRule, msg, topic)
				nextRuleResult := inTimeExecute(nextRule, msg, topic)
				if index == 0 {
					flag = utils.DetermineLogicResult(preRuleResult, ruleArray[index+1], nextRuleResult)
				}
				if index != 0 {
					flag = utils.DetermineLogicResult(flag, ruleArray[index+1], nextRuleResult)
				}
			}
		}
	}
	log.Println("同步规条件解析结果是", flag)
	updateRecordAndResult(flag, detail, msg)
}

// inTimeExecute 同步处理器处理规则信息
func inTimeExecute(sDA string, msg types.AppNoticeMsg, topic string) bool {
	sourceDeviceAttr := strings.Split(sDA, "-")
	//规则数据源信息
	ruleTopic := sourceDeviceAttr[0]
	//规则设备SN信息
	ruleDevice := sourceDeviceAttr[1]
	//规则设备属性信息
	ruleAttr := sourceDeviceAttr[2]
	//规则判断标识信息
	ruleExpr := sourceDeviceAttr[3]
	//规则数据值信息
	ruleData := sourceDeviceAttr[4]
	//数据运行结果表达
	var flag = false
	if topic == ruleTopic {
		if msg.DevSN == ruleDevice {
			reportData, ok := msg.Param[ruleAttr]
			if ok {
				reData := fmt.Sprint(reportData)
				flag = utils.DetermineResultByStr(ruleData, ruleExpr, reData)
			}
		}
	}
	return flag
}

// dataIfInRule 判断消息数据是否匹配规则
func dataIfInRule(sDA string, msg types.AppNoticeMsg, topic string) bool {
	sourceDeviceAttr := strings.Split(sDA, "-")
	//规则数据源信息
	ruleTopic := sourceDeviceAttr[0]
	//规则设备SN信息
	ruleDevice := sourceDeviceAttr[1]
	//规则设备属性信息
	ruleAttr := sourceDeviceAttr[2]
	//数据运行结果表达
	if topic == ruleTopic {
		if msg.DevSN == ruleDevice {
			_, ok := msg.Param[ruleAttr]
			if ok {
				return true
			}
		}
	}
	return false
}
