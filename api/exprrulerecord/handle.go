package record

import (
	"LinEngineRules/constants"
	"LinEngineRules/model"
	"LinEngineRules/utils"
	"github.com/emicklei/go-restful/v3"
	"log"
	"time"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/18 17:06
CREATE_BY:GoLand.LinEngineRules
*/

// ListAllExprRuleRecord 查询所有规则记录
func ListAllExprRuleRecord(req *restful.Request, resp *restful.Response) {
	offset, limit, err := utils.GetOffsetAndLimit(req)
	if err != nil {
		log.Println("分页数据发生问题", err)
		return
	}
	name := req.QueryParameter("name")
	var exprRuleRecord model.ExprRuleRecord
	record, total := exprRuleRecord.PageRuleDetail(offset, limit, name)
	utils.RespWithDataAndCnt(resp, record, total)
}

func init() {
	var rd model.RuleDataSource
	constants.RuleDataSourceArray = rd.ListAllRuleDataSources()
	log.Println("获取到所有数据源是：", constants.RuleDataSourceArray)
	go deleteLazyData()
}

func deleteLazyData() {
	ticker := time.NewTicker(time.Hour * 24) //创建一个周期性定时器
	for {
		select {
		case <-ticker.C:
			var a model.AsynchronousHistoryData
			err := a.DeleteExpiredDataByTime()
			if err != nil {
				log.Println("清理数据出错", err)
				return
			}
		}

	}
}
