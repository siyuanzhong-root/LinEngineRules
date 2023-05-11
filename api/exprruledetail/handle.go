package detail

import (
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

// PageAllExprRule 查询所有规则信息
func PageAllExprRule(req *restful.Request, resp *restful.Response) {
	offset, limit, err := utils.GetOffsetAndLimit(req)
	if err != nil {
		log.Println("分页数据发生问题", err)
		return
	}
	status := req.QueryParameter("status")
	handleName := req.QueryParameter("handleName")
	name := req.QueryParameter("name")
	var ruleDetail model.RuleDetail
	details, total := ruleDetail.PageRuleDetail(offset, limit, status, handleName, name)
	utils.RespWithDataAndCnt(resp, details, total)
}

// AddRuleDetail 新增规则信息
func AddRuleDetail(req *restful.Request, resp *restful.Response) {
	var ruleDetail model.RuleDetail
	err := req.ReadEntity(&ruleDetail)
	if err != nil {
		log.Println("新增获取规则详情入参失败", err)
		utils.RespErrWithData(resp, "新增获取规则详情入参失败"+err.Error())
		return
	}
	ruleDetail.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	err = ruleDetail.Insert()
	if err != nil {
		log.Println("插入规则数据出错", err)
		utils.RespErrWithData(resp, "插入规则数据出错"+err.Error())
		return
	}
	utils.RespOKWithData(resp, "新增成功")
}

// UpdateRuleDetail 更新规则信息
func UpdateRuleDetail(req *restful.Request, resp *restful.Response) {
	var ruleDetail model.RuleDetail
	err := req.ReadEntity(&ruleDetail)
	if err != nil {
		log.Println("更新获取规则详情入参失败", err)
		utils.RespErrWithData(resp, "更新获取规则详情入参失败"+err.Error())
		return
	}
	err = ruleDetail.Update()
	if err != nil {
		log.Println("更新规则数据出错", err)
		utils.RespErrWithData(resp, "更新规则数据出错"+err.Error())
		return
	}
	utils.RespOKWithData(resp, "更新成功")
}

// DeleteRuleDetail 删除规则信息
func DeleteRuleDetail(req *restful.Request, resp *restful.Response) {
	var ruleDetail model.RuleDetail
	err := req.ReadEntity(&ruleDetail)
	if err != nil {
		log.Println("删除获取规则详情入参失败", err)
		utils.RespErrWithData(resp, "删除获取规则详情入参失败"+err.Error())
		return
	}
	err = ruleDetail.Delete()
	if err != nil {
		log.Println("删除规则数据出错", err)
		utils.RespErrWithData(resp, "删除规则数据出错"+err.Error())
		return
	}
	utils.RespOKWithData(resp, "删除成功")
}

// ListAllRuleDetail 获取所有的规则列表信息
func ListAllRuleDetail(req *restful.Request, resp *restful.Response) {
	var ruleDetail model.RuleDetail
	utils.RespOKWithData(resp, ruleDetail.ListAllRuleDetails())
}
