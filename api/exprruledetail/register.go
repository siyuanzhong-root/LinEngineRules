package detail

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

// SetupRouteAPI 规则详情结果集
func SetupRouteAPI(ws *restful.WebService) {
	ws.
		Route(
			ws.GET("details/page").
				To(PageAllExprRule).
				Doc("分页查询所有规则").
				Returns(http.StatusOK, constants.ResponseStatusOK, model.RuleDetail{}).
				Param(restful.QueryParameter("p", "页码")).
				Param(restful.QueryParameter("l", "页面容量")).
				Param(restful.QueryParameter("name", "规则名称")).
				Param(restful.QueryParameter("status", "规则状态")).
				Param(restful.QueryParameter("handleName", "处理方法名")).
				Metadata(rest.KeyOpenAPITags, []string{constants.DetailTag})).
		Route(
			ws.POST("details/update").
				To(UpdateRuleDetail).
				Doc("更新规则详情").
				Returns(http.StatusOK, constants.ResponseStatusOK, "result").
				Param(restful.BodyParameter("ruleDetail", "规则详情")).
				Metadata(rest.KeyOpenAPITags, []string{constants.DetailTag})).
		Route(
			ws.POST("details/add").
				To(AddRuleDetail).
				Doc("新增规则详情").
				Returns(http.StatusOK, constants.ResponseStatusOK, "result").
				Param(restful.BodyParameter("ruleDetail", "规则详情")).
				Metadata(rest.KeyOpenAPITags, []string{constants.DetailTag})).
		Route(
			ws.POST("details/delete").
				To(DeleteRuleDetail).
				Doc("删除规则详情").
				Returns(http.StatusOK, constants.ResponseStatusOK, "result").
				Param(restful.BodyParameter("ruleDetail", "规则详情")).
				Metadata(rest.KeyOpenAPITags, []string{constants.DetailTag})).
		Route(
			ws.GET("details/list").
				To(ListAllRuleDetail).
				Doc("获取到所有规则列表").
				Returns(http.StatusOK, constants.ResponseStatusOK, model.RuleDetail{}).
				Metadata(rest.KeyOpenAPITags, []string{constants.DetailTag}))
}
