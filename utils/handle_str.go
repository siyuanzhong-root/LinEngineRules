package utils

import (
	"log"
	"strconv"
	"strings"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/6 17:55
CREATE_BY:GoLand.LinEngineRules
*/

// DetermineResultByStr 根据str类型返回类型值
func DetermineResultByStr(ruleData, expr, reportData string) bool {
	ruData, err := strconv.ParseFloat(ruleData, 64)
	if err != nil {
		log.Println("规则预设参数值转换出错", err)
		return false
	}
	reData, err := strconv.ParseFloat(reportData, 64)
	if err != nil {
		log.Println("上报参数转换出错", err)
		return false
	}
	switch expr {
	case "=":
		if ruData == reData {
			return true
		}
	case ">":
		if reData > ruData {
			return true
		}
	case "<":
		if reData < ruData {
			return true
		}
	case "<=":
		if reData <= ruData {
			return true
		}
	case ">=":
		if reData >= ruData {
			return true
		}

	case "!=":
		if reData != ruData {
			return true
		}
	}

	return false
}

// DetermineLogicResult 根据逻辑操作符判断结果
func DetermineLogicResult(pre bool, logic string, now bool) bool {
	log.Println("传入数据是：", pre, logic, now)
	switch logic {
	case "&":
		return pre && now
	case "|":
		return pre || now
	}
	return false
}

// SpiltByDefineStr 多字符串分隔规则合集
func SpiltByDefineStr(s string) []string {
	var spiltItem = []rune{'&', '|'}
	split := func(r rune) bool {
		for _, v := range spiltItem {
			if v == r {
				return true
			}
		}
		return false
	}
	a := strings.FieldsFunc(s, split)
	return a
}

// HandleReturnFullResult 处理返回完整规则结构
func HandleReturnFullResult(expr string) []string {
	result := SpiltByDefineStr(expr)
	var tmp int
	var backData []string
	if len(result) == 1 {
		backData = append(backData, result[0])
	}
	if len(result) > 1 {
		for i, item := range result {
			var logic string
			tmp += len(item)
			if i == 0 {
				logic = string([]rune(expr)[len(item)])
			}
			if i < len(result)-1 {
				logic = string([]rune(expr)[tmp+i])
			}
			if logic != "" {
				backData = append(backData, item, logic)
			}
		}
		backData = append(backData, result[len(result)-1])
	}
	log.Println("完整解析结果是", backData)
	return backData
}
