package util

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/2 13:54
CREATE_BY:GoLand.LinEngineRules
*/

// StrInArray 判断字符串是否在切片中
func StrInArray(sl []string, s string) bool {
	set := make(map[string]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	_, ok := set[s]
	return ok
}

// MapAppend Map追加数据
func MapAppend(rootMap map[string]interface{}, addedMap map[string]interface{}) map[string]interface{} {
	for key, value := range addedMap {
		rootMap[key] = value
	}
	return rootMap
}
