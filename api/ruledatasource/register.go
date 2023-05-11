package datasource

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
			ws.GET("datasource/listAll").
				To(ListAllDataSource).
				Doc("查询数据源列表").
				Returns(http.StatusOK, constants.ResponseStatusOK, model.RuleDataSource{}).
				Metadata(rest.KeyOpenAPITags, []string{constants.DataSourceTag}))
}
