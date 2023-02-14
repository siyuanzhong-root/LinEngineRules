package execute

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/siyuanzhong-root/LinEngineRules/data_driver"
	"github.com/siyuanzhong-root/LinEngineRules/rule_expr"
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/constants"
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/handle/types"
	"github.com/siyuanzhong-root/LinEngineRules/rule_operate/handle/util"
	"log"
	"strconv"
	"strings"
	"sync"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/2 10:28
CREATE_BY:GoLand.LinEngineRules
*/

// DriverEngine 引擎驱动
func DriverEngine() {
	if err := mqtt.Subscribe(constants.MQTTListenALLTopic, constants.MQTTQos, HandleMqttMessage); err != nil {
		log.Println("接受消息出错")
		return
	}
}

// HandleMqttMessage 处理消息
func HandleMqttMessage(msg []byte, topic string) {
	//接受MQTT上报消息
	var appMsg types.AppNoticeMsg
	json.Unmarshal(msg, &appMsg)
	var wg sync.WaitGroup
	wg.Add(len(constants.RuleExchangeArray))
	for i, ruleExchange := range constants.RuleExchangeArray {
		go func(exchange map[string]interface{}, i int) {
			for key := range exchange {
				if strings.Contains(key, "-") {
					source := util.HandleStrToArrayBy0(key)[0]
					deviceSN := util.HandleStrToArrayBy0(key)[1]
					attribute := util.HandleStrToArrayBy0(key)[2]
					log.Println("****线程", i, "取到规则名", source, deviceSN, attribute, appMsg.Param[attribute])
					data, f := appMsg.Param[attribute]
					d := fmt.Sprint(data)
					if source == topic && deviceSN == appMsg.DevSN && f != false {
						switch util.DetermineStrType(d) {
						case "string":
							constants.CacheDataOperate[i][attribute] = d
						case "int":
							log.Println("获取到类型值是INT")
							dataInt, _ := strconv.Atoi(d)
							constants.CacheDataOperate[i][attribute] = dataInt
						case "bool":
							dataBool, _ := strconv.ParseBool(d)
							constants.CacheDataOperate[i][attribute] = dataBool
						case "float":
							dataFloat, _ := strconv.ParseFloat(d, 64)
							constants.CacheDataOperate[i][attribute] = dataFloat
						}
						go DataCache(i)
					}
				}
			}
			defer wg.Done()
		}(ruleExchange, i)
	}
	wg.Wait()
	log.Println("消息体结构是", constants.CacheDataOperate)
	var flag = true
	for in, va := range constants.CacheDataOperate {
		for _, value := range va {
			if value == "" || len(va) < len(constants.RuleExchangeArray[in]) {
				flag = false
			}
		}
		if flag == true {
			log.Println("****线程", in, "进入判断触发操作******")
			//进行判断触发
			HandleExprData(constants.CacheDataOperate[in], constants.ExprStrArray[in])
			constants.CacheDataOperate[in] = map[string]interface{}{}
		}
	}
}

func DataCache(i int) {
	defer constants.TimeSchedule.Stop()
	for {
		select {
		case <-constants.TimeSchedule.C:
			constants.CacheDataOperate[i] = map[string]interface{}{}
			log.Println("****线程", i, "进入数据清理了一次*********")
		}
	}
}
func HandleExprData(controlMap map[string]interface{}, expr string) {
	log.Println("执行了一次,规则数据是：", controlMap, "表达式是：", expr)
	engine, err := rule_expr.NewEngine(expr)
	if err != nil {
		fmt.Println("rules has error", err.Error())
	}
	result, err := engine.RunRule(controlMap)
	if err != nil {
		fmt.Println("error has occurred", err.Error())
	}
	log.Println("规则执行结果", result)
}
