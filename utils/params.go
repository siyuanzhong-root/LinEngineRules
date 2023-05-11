package utils

import (
	"github.com/emicklei/go-restful/v3"
	"strconv"
)

func GetID(req *restful.Request) (uint, error) {
	idStr := req.PathParameter("id")
	id, err := strconv.Atoi(idStr)
	return uint(id), err
}

func GetOffsetAndLimit(req *restful.Request) (int, int, error) {
	pageStr := req.QueryParameter("p")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return -1, -1, err
	}

	limitStr := req.QueryParameter("l")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return -1, -1, err
	}

	offset := (page - 1) * limit
	return offset, limit, nil
}
