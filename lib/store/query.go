package store

import (
	"fmt"
	"strings"
)

const (
	defaultLimit   = 10
	defaultSort    = "created_at"
	defaultSortDir = "DESC"
)

type QueryParams struct {
	Limit    int      `json:"limit" query:"limit"`
	Offset   int      `json:"offset" query:"offset"`
	Keywords string   `json:"keywords" query:"keywords"`
	Sort     []string `json:"sort" query:"sort"`
	Filter   []string `json:"filter" query:"filter"`
}

type QueryResponse struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	TotalData int64  `json:"totalRows"`
	Keywords  string `json:"keywords"`
}

func (q *QueryParams) BuildPagination(f []string) (limit string, offset string, sort string, filter string, keywords string) {
	limit = fmt.Sprintf("LIMIT %d", defaultLimit)
	offset = "OFFSET 0"
	sort = fmt.Sprintf("ORDER BY %s %s", defaultSort, defaultSortDir)
	filter = "1"
	keywords = "1"

	if q.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", q.Limit)
	}

	if q.Offset > 1 {
		offset = fmt.Sprintf("OFFSET %d", q.Offset)
	}

	if len(q.Sort) > 0 {
		sortDir := "DESC"
		if q.Sort[1] == "asc" {
			sortDir = "ASC"
		}
		sort = fmt.Sprintf("ORDER BY %s %s", q.Sort[0], sortDir)
	}

	if len(q.Filter) >= 3 {
		filter = fmt.Sprintf("%s %s '%v'", q.Filter[0], q.Filter[1], q.Filter[2])
	}

	if q.Keywords != "" {
		if len(f) > 0 {
			key := make([]string, 0)
			for _, v := range f {
				val := fmt.Sprintf("%s LIKE '%s'", v, "%"+q.Keywords+"%")
				key = append(key, val)
			}
			keywords = strings.Join(key, " OR ")
		}
	}

	return
}

func (q *QueryParams) BuildQueryResponse(totalData int64) (res QueryResponse) {
	res.Keywords = q.Keywords
	res.Limit = q.Limit
	res.TotalData = totalData
	res.Offset = q.Offset

	return res
}
