package handler

import (
	"LinEngineRules/constants"
	"LinEngineRules/model"
	"LinEngineRules/types"
	"LinEngineRules/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/21 14:27
CREATE_BY:GoLand.LinEngineRules
*/

// AsyncRuleHandle 异步处理器控制
func AsyncRuleHandle(detail model.RuleDetail, msg types.AppNoticeMsg, topic string) {
	ruleArray := utils.HandleReturnFullResult(detail.SourceDeviceAttr)
	var ifMarchDevice = false
	for i := range ruleArray {
		if strings.Contains(ruleArray[i], "-") && dataIfInRule(ruleArray[i], msg, topic) {
			ifMarchDevice = true
		}
	}
	if !ifMarchDevice {
		return
	}
	var flag bool
	if len(ruleArray) == constants.SingleRuleExpr {
		flag = asyncExecute(detail, ruleArray[0], msg, topic)
	}
	if len(ruleArray) >= constants.MultiRuleExpr {
		for index := range ruleArray {
			if index%2 == 0 && index < len(ruleArray)-2 {
				preRule := ruleArray[index]
				nextRule := ruleArray[index+2]
				preRuleResult := asyncExecute(detail, preRule, msg, topic)
				nextRuleResult := asyncExecute(detail, nextRule, msg, topic)
				if index == 0 {
					flag = utils.DetermineLogicResult(preRuleResult, ruleArray[index+1], nextRuleResult)
				}
				if index != 0 {
					flag = utils.DetermineLogicResult(flag, ruleArray[index+1], nextRuleResult)
				}
			}
		}
	}
	log.Println("异步规条件解析结果是", flag)
	updateRecordAndResult(flag, detail, msg)
}

// getAllDevSNByRule 获取异步规则执行结果
func getAllDevSNByRule(ruleDetail model.RuleDetail) bool {
	allDetail := utils.SpiltByDefineStr(ruleDetail.SourceDeviceAttr)
	var flag = false
	for _, detail := range allDetail {
		sourceDeviceAttr := strings.Split(detail, "-")
		//规则设备SN信息
		ruleDevice := sourceDeviceAttr[1]
		//规则设备属性信息
		ruleAttr := sourceDeviceAttr[2]
		//规则判断标识信息
		ruleExpr := sourceDeviceAttr[3]
		//规则数据值信息
		ruleData := sourceDeviceAttr[4]
		var async model.AsynchronousHistoryData
		async.DevSN = ruleDevice
		devData := async.QueryNewerDataByDevSN()
		var tmpStruct = make(map[string]interface{})
		if devData.ReceiveData != "" {
			err := json.Unmarshal([]byte(devData.ReceiveData), &tmpStruct)
			if err != nil {
				log.Println("异步结构体解析", devData.ReceiveData, "出错", err)
				return flag
			}
			if reportData, ok := tmpStruct[ruleAttr]; ok != false {
				flag = utils.DetermineResultByStr(ruleData, ruleExpr, fmt.Sprint(reportData))
				if flag == false {
					return flag
				}
			}
		}
	}
	return flag
}

// asyncExecute 同步处理器处理规则信息
func asyncExecute(ruleDetail model.RuleDetail, sDA string, msg types.AppNoticeMsg, topic string) bool {
	sourceDeviceAttr := strings.Split(sDA, "-")
	//规则数据源信息
	ruleTopic := sourceDeviceAttr[0]
	//规则设备SN信息
	ruleDevice := sourceDeviceAttr[1]
	//规则设备属性信息
	ruleAttr := sourceDeviceAttr[2]
	// 数据运行结果表达
	var flag = false
	if topic == ruleTopic {
		if msg.DevSN == ruleDevice {
			_, ok := msg.Param[ruleAttr]
			if ok {
				flag = getAllDevSNByRule(ruleDetail)
			}
		}
	}
	return flag
}
