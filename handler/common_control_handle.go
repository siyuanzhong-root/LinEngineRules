package handler

import (
	"LinEngineRules/constants"
	"LinEngineRules/initdata"
	"LinEngineRules/model"
	"LinEngineRules/types"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/20 17:12
CREATE_BY:GoLand.LinEngineRules
*/

// CommonControlHandle 通用控制结果方法
func CommonControlHandle(handleDetail string) error {
	controlInfo := strings.Split(handleDetail, ",")
	for _, cd := range controlInfo {
		controlDetail := strings.Split(cd, "-")
		controlLink := controlDetail[0]
		controlDevice := controlDetail[1]
		controlType := controlDetail[2]
		controlMsg := controlDetail[3]
		mqttPayload := fmt.Sprintf(constants.ControlDeviceSendJson, controlDevice, time.Now().UnixNano()/1e3, controlDevice, constants.ControlCommand, controlMsg)
		if controlType == constants.DirectControlType {
			err := commonControlExecute(controlLink, mqttPayload)
			if err != nil {
				return err
			}
		}
		if controlType == constants.DelayControlType {
			var err error
			delayMinute, _ := strconv.Atoi(controlDetail[4])
			dr := time.Duration(delayMinute) * time.Minute
			time.AfterFunc(dr, func() {
				err = commonControlExecute(controlLink, mqttPayload)
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func commonControlExecute(controlLink, mqttPayload string) error {
	if controlLink == "lora" {
		if err := initdata.Publish(constants.LoraDeviceControlTopic, 0, false, mqttPayload); err != nil {
			log.Println("下发Lora控制失败", err)
			return err
		}
	}
	if controlLink == "ble" {
		if err := initdata.Publish(constants.BleDeviceControlTopic, 0, false, mqttPayload); err != nil {
			log.Println("下发Ble控制失败", err)
			return err
		}
	}
	return nil
}

func updateRecordAndResult(flag bool, detail model.RuleDetail, msg types.AppNoticeMsg) {
	data, _ := json.Marshal(msg)
	var ruleRecord = model.ExprRuleRecord{
		Name:        detail.Name,
		Expr:        detail.SourceDeviceAttr,
		Record:      string(data),
		TriggerTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 规则条件判断结果处理未通过
	if !flag {
		ruleRecord.Result = constants.ExecuteFailedResult
		ruleRecord.ControlPublish = constants.ControlFailedResult
		err := ruleRecord.Insert()
		if err != nil {
			log.Println("插入规则记录表失败", err)
		}
		detail.FailedCount += 1
		err = detail.Update()
		if err != nil {
			log.Println("更新失败次数集失败", err)
			return
		}
	}

	// 规则条件判断结果处理通过
	if flag {
		ruleRecord.Result = constants.ExecuteSuccessResult
		err := CommonControlHandle(detail.HandleDetail)
		// 如果控制出错
		if err != nil {
			log.Println("规则：", detail.Name, "控制失败")
			ruleRecord.ControlPublish = constants.ControlFailedResult
			err := ruleRecord.Insert()
			if err != nil {
				log.Println("插入规则记录表失败", err)
			}
			detail.FailedCount = detail.FailedCount + 1
			err = detail.Update()
			if err != nil {
				log.Println("更新失败次数集失败", err)
				return
			}
			return
		}
		detail.SuccessCount = detail.SuccessCount + 1
		err = detail.Update()
		ruleRecord.ControlPublish = constants.ControlSuccessResult
		err = ruleRecord.Insert()
		if err != nil {
			log.Println("插入规则记录表失败", err)
		}
		if err != nil {
			log.Println("更新成功次数集失败", err)
			return
		}
	}

}
