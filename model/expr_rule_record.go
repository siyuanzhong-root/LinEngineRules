package model

import (
	"LinEngineRules/initdata"
	"log"
	"time"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/20 17:55
CREATE_BY:GoLand.LinEngineRules
*/

// ExprRuleRecord 规则执行记录
type ExprRuleRecord struct {
	ID             uint   `gorm:"column:id" json:"id" description:"规则ID"`
	Name           string `gorm:"column:name" json:"name" description:"规则名"`
	Expr           string `gorm:"column:expr" json:"expr"`
	Record         string `gorm:"column:record" json:"record"`
	Result         int    `gorm:"column:result" json:"result"`
	TriggerTime    string `gorm:"column:trigger_time" json:"trigger_time"`
	ControlPublish int    `gorm:"column:control_publish" json:"control_publish"`
}

// TableName 返回规则表表名
func (e *ExprRuleRecord) TableName() string {
	return "expr_rule_record"
}

// Insert 添加规则执行记录
func (e *ExprRuleRecord) Insert() error {
	return initdata.EngineDB.Table(e.TableName()).Save(e).Error
}

// PageRuleDetail 获取规则记录分页
func (e *ExprRuleRecord) PageRuleDetail(offset, limit int, name string) (result []ExprRuleRecord, total int64) {
	var number int64
	var exprRuleRecord []ExprRuleRecord
	tx := initdata.EngineDB.Table(e.TableName())
	if name != "" {
		tx = tx.Where("name= ?", name)
	}
	err := tx.Count(&number).Offset(offset).Limit(limit).Order("trigger_time DESC").Find(&exprRuleRecord).Error
	if err != nil {
		log.Println("获取到规则名失败", err)
		return []ExprRuleRecord{}, 0
	}
	return exprRuleRecord, number
}

// DeleteByName 删除规则记录
func (e *ExprRuleRecord) DeleteByName() error {
	return initdata.EngineDB.Table(e.TableName()).Where("name = ?", e.Name).Delete(ExprRuleRecord{}).Error
}

// DeleteExpiredByName 删除规则记录
func (e *ExprRuleRecord) DeleteExpiredByName() error {
	expiredTime := time.Now().AddDate(0, 0, -30)
	return initdata.EngineDB.Table(e.TableName()).Where("trigger_time < ?", expiredTime).Delete(ExprRuleRecord{}).Error
}
