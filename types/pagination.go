package types

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type SortEnum string

const (
	ASC  SortEnum = "ASC"
	DESC SortEnum = "DESC"
)

type Pagination struct {
	Take   int      `json:"take"`
	Skip   int      `json:"skip"`
	Sort   SortEnum `json:"sort"`
	Field  string   `json:"field"`
	Search string   `json:"search"`
}

func GetPaginationParams(r *http.Request) (Pagination, error) {
	query := r.URL.Query()

	// Default values
	pagination := Pagination{
		Take:   10,
		Skip:   0,
		Sort:   DESC,
		Field:  "created_at",
		Search: "",
	}

	// Extract and validate the pagination parameters
	if takeStr := query.Get("take"); takeStr != "" {
		take, err := strconv.Atoi(takeStr)
		if err != nil || take < 0 {
			return Pagination{}, fmt.Errorf("invalid take value")
		}
		pagination.Take = take
	}

	if skipStr := query.Get("skip"); skipStr != "" {
		skip, err := strconv.Atoi(skipStr)
		if err != nil || skip < 0 {
			return Pagination{}, fmt.Errorf("invalid skip value")
		}
		pagination.Skip = skip
	}

	if sortStr := query.Get("sort"); sortStr != "" {
		sortStr = strings.ToUpper(sortStr)
		if sortStr != string(ASC) && sortStr != string(DESC) {
			return Pagination{}, fmt.Errorf("invalid sort value")
		}
		pagination.Sort = SortEnum(sortStr)
	}

	if fieldStr := query.Get("field"); fieldStr != "" {
		pagination.Field = fieldStr
	}

	if searchStr := query.Get("search"); searchStr != "" {
		pagination.Search = searchStr
	}

	return pagination, nil
}
