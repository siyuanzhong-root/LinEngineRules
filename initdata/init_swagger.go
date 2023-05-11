package initdata

import "github.com/go-openapi/spec"

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/19 10:06
CREATE_BY:GoLand.LinEngineRules
*/

// EnrichSwaggerObject swagger配置信息
func EnrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:   "规则可视化后台API",
			Version: "0.0.1",
		},
	}
	swo.Tags = []spec.Tag{
		{TagProps: spec.TagProps{Name: "规则详情信息"}},
		{TagProps: spec.TagProps{Name: "系统概览"}},
	}
	swo.Schemes = []string{"http"}
}
