// Package vo 通用返回JSON结构体
// @author: kbj
// @date: 2023/2/3
package vo

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

type Page[T any] struct {
	Records  []T   `json:"records"`  // 分页数据
	Total    int64 `json:"total"`    // 总条目数
	PageSize int   `json:"pageSize"` // 每页数据量
	PageNum  int   `json:"pageNum"`  // 当前页数
}

type Ids struct {
	IDs []uint `params:"ids"`
}
