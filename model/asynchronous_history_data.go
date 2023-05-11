package model

import (
	"LinEngineRules/initdata"
	"log"
	"time"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/19 17:41
CREATE_BY:GoLand.LinEngineRules
*/

// AsynchronousHistoryData 数据历史记录表
type AsynchronousHistoryData struct {
	ID          uint   `gorm:"column:id" json:"id" description:"数据ID"`
	DevSN       string `gorm:"column:dev_sn" json:"dev_sn" description:"设备SN"`
	ReceiveData string `gorm:"column:receive_data" json:"receive_data"`
	Timestamp   string `gorm:"column:timestamp" json:"timestamp"`
}

// TableName 返回规则表表名
func (a *AsynchronousHistoryData) TableName() string {
	return "asynchronous_history_data"
}

// Insert 插入数据
func (a *AsynchronousHistoryData) Insert() error {
	return initdata.EngineDB.Table(a.TableName()).Save(a).Error
}

// QueryNewerDataByDevSN 根据SN查询最新的数据
func (a *AsynchronousHistoryData) QueryNewerDataByDevSN() AsynchronousHistoryData {
	var async AsynchronousHistoryData
	err := initdata.EngineDB.Table(a.TableName()).Where("dev_sn=?", a.DevSN).
		Order("timestamp DESC").Limit(1).Find(&async).Error
	if err != nil {
		log.Println("查询最新数据出错", err)
		return AsynchronousHistoryData{}
	}
	return async
}

// DeleteExpiredDataByTime 查询过期时间
func (a *AsynchronousHistoryData) DeleteExpiredDataByTime() error {
	expiredTime := time.Now().AddDate(0, 0, -30)
	err := initdata.EngineDB.Table(a.TableName()).Where("timestamp < ?", expiredTime).Delete(AsynchronousHistoryData{}).Error
	if err != nil {
		log.Println("删除数据出错", err)
		return err
	}
	return nil
}
