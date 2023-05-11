package datasource

import (
	"LinEngineRules/model"
	"LinEngineRules/utils"
	"github.com/emicklei/go-restful/v3"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/18 17:06
CREATE_BY:GoLand.LinEngineRules
*/

// ListAllDataSource 查询所有数据源信息
func ListAllDataSource(req *restful.Request, resp *restful.Response) {
	var rd model.RuleDataSource
	details := rd.ListAllRuleDataSources()
	utils.RespOKWithData(resp, details)
}
