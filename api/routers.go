package api

import (
	"LinEngineRules/api/exprruledetail"
	record "LinEngineRules/api/exprrulerecord"
	"LinEngineRules/api/ruledatasource"
	"LinEngineRules/handler"
	"github.com/emicklei/go-restful/v3"
)

func init() {
	ws := &restful.WebService{}
	ws.Path("/api/rule").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	//规则详情接口集
	detail.SetupRouteAPI(ws)
	//数据源详情接口集
	datasource.SetupRouteAPI(ws)
	//记录源接口集
	record.SetupRouteAPI(ws)
	restful.Add(ws)
	handler.ConsumeMqDataCenter()
}
