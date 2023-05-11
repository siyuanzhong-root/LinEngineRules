package model

import (
	"LinEngineRules/initdata"
	"log"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/18 16:56
CREATE_BY:GoLand.LinEngineRules
*/

// RuleDetail 规则详情
type RuleDetail struct {
	ID               uint   `gorm:"column:id" json:"id" description:"规则ID"`
	Name             string `gorm:"column:name" json:"name" description:"规则名"`
	Status           string `gorm:"column:status" json:"status"`
	IsAsync          int    `gorm:"column:is_async" json:"is_async"`
	SourceDeviceAttr string `gorm:"column:source_device_attr" json:"source_device_attr"`
	HandleName       string `gorm:"column:handle_name" json:"handle_name"`
	Timestamp        string `gorm:"timestamp" json:"timestamp"`
	SuccessCount     int    `gorm:"success_count" json:"success_count"`
	FailedCount      int    `gorm:"failed_count" json:"failed_count"`
	HandleDetail     string `gorm:"column:handle_detail" json:"handle_detail"`
}

// TableName 返回规则表表名
func (r *RuleDetail) TableName() string {
	return "expr_rule_detail"
}

// PageRuleDetail 获取所有规则名
func (r *RuleDetail) PageRuleDetail(offset, limit int, status, handleName, name string) (result []RuleDetail, total int64) {
	var number int64
	var rd []RuleDetail
	tx := initdata.EngineDB.Table(r.TableName())
	if status != "" {
		tx = tx.Where("status =?", status)
	}
	if handleName != "" {
		tx = tx.Where("handle_name =?", handleName)
	}
	if name != "" {
		tx = tx.Where("name= ?", name)
	}
	err := tx.Count(&number).Offset(offset).Limit(limit).Order("timestamp DESC").Find(&rd).Error
	if err != nil {
		log.Println("获取到规则名失败", err)
		return []RuleDetail{}, 0
	}
	return rd, number
}

// ListAllRuleDetails 获取所有规则名
func (r *RuleDetail) ListAllRuleDetails() (result []RuleDetail) {
	var rd []RuleDetail
	err := initdata.EngineDB.Table(r.TableName()).Find(&rd).Error
	if err != nil {
		log.Println("获取到规则名失败", err)
		return []RuleDetail{}
	}
	return rd
}

// Update 更新规则详情
func (r *RuleDetail) Update() error {
	return initdata.EngineDB.Table(r.TableName()).Updates(r).Error
}

// Insert 新增规则名
func (r *RuleDetail) Insert() error {
	return initdata.EngineDB.Table(r.TableName()).Save(r).Error
}

// Delete 删除规则
func (r *RuleDetail) Delete() error {
	return initdata.EngineDB.Table(r.TableName()).Delete(r).Error
}
