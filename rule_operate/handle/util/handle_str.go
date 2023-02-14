package util

import (
	"strconv"
	"strings"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/2/6 17:55
CREATE_BY:GoLand.LinEngineRules
*/

func HandleStrToArrayBy0(str string) []string {
	var res []string
	res = strings.Split(str, "-")
	return res
}
func DetermineStrType(str string) string {
	boolFlag := strings.ContainsAny(str, "true") || strings.ContainsAny(str, "false")
	_, err := strconv.ParseFloat(str, 64)
	intFlag := true
	for i := range str {
		if (str[i] < 48 || str[i] > 59) && str[i] != '.' {
			intFlag = false
		}
	}
	floatFlag := strings.Contains(str, ".") && err == nil && intFlag
	if boolFlag {
		return "bool"
	}
	if floatFlag {
		return "float"
	}
	if intFlag {
		return "int"
	}
	return "string"
}
