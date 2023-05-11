package record

import (
	"LinEngineRules/constants"
	"LinEngineRules/model"
	rest "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/19 9:23
CREATE_BY:GoLand.LinEngineRules
*/

// SetupRouteAPI 数据源详情接口
func SetupRouteAPI(ws *restful.WebService) {
	ws.
		Route(
			ws.GET("rulerecord/page").
				To(ListAllExprRuleRecord).
				Doc("查询规则记录信息").
				Param(restful.QueryParameter("p", "页码")).
				Param(restful.QueryParameter("l", "页面容量")).
				Param(restful.QueryParameter("name", "规则名称")).
				Returns(http.StatusOK, constants.ResponseStatusOK, model.ExprRuleRecord{}).
				Metadata(rest.KeyOpenAPITags, []string{constants.ExprRuleRecordTag}))
}
