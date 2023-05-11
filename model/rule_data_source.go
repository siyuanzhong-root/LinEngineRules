package model

import (
	"LinEngineRules/initdata"
	"log"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/19 14:43
CREATE_BY:GoLand.LinEngineRules
*/

// RuleDataSource 规则数据源
type RuleDataSource struct {
	ID        uint   `gorm:"column:id" json:"id" description:"数据源ID"`
	Name      string `gorm:"column:name" json:"name" description:"数据源名称"`
	Detail    string `gorm:"column:detail" json:"detail"`
	Timestamp string `gorm:"timestamp" json:"timestamp"`
}

// TableName 返回规则表表名
func (r *RuleDataSource) TableName() string {
	return "rule_data_source"
}

// ListAllRuleDataSources 获取所有规则名
func (r *RuleDataSource) ListAllRuleDataSources() (result []RuleDataSource) {
	var rd []RuleDataSource
	err := initdata.EngineDB.Table(r.TableName()).Find(&rd).Error
	if err != nil {
		log.Println("获取数据流失败", err)
		return []RuleDataSource{}
	}
	return rd
}
