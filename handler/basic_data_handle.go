package handler

import (
	"LinEngineRules/constants"
	"LinEngineRules/initdata"
	"LinEngineRules/model"
	"LinEngineRules/types"
	"encoding/json"
	"log"
	"time"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/19 14:35
CREATE_BY:GoLand.LinEngineRules
*/

// ConsumeMqDataCenter 消费MQ数据中心方法
func ConsumeMqDataCenter() {
	for _, source := range constants.RuleDataSourceArray {
		err := initdata.Subscribe(source.Detail, constants.MQTTQosCount, dataHandler)
		if err != nil {
			log.Println("监听主题", source.Detail, "出错", err)
			return
		}
	}
}

func dataHandler(msg []byte, topic string) {
	//数据收集流程
	log.Println("接收到Topic为:", topic, ";Data是:\n", string(msg))
	var appMsg types.AppNoticeMsg
	err := json.Unmarshal(msg, &appMsg)
	if err != nil {
		log.Println("上报数据流出现问题", err)
		return
	}
	var ac model.AsynchronousHistoryData
	ac.DevSN = appMsg.DevSN
	param, err := json.Marshal(appMsg.Param)
	if err != nil {
		log.Println("解析param参数出问题", err)
		return
	}
	ac.ReceiveData = string(param)
	ac.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	err = ac.Insert()
	if err != nil {
		log.Println("数据插入失败", err)
		return
	}
	var rd model.RuleDetail

	for _, detail := range rd.ListAllRuleDetails() {
		if detail.Status == constants.StatusEnable {
			if detail.IsAsync == constants.NotAsyncRuleSymbol {
				//执行同步任务
				InTimeHandle(detail, appMsg, topic)
			}
			if detail.IsAsync == constants.IsAsyncRuleSymbol {
				//执行异步任务
				AsyncRuleHandle(detail, appMsg, topic)
			}
		}
	}
}
