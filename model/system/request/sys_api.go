package request

type SearchApiParams struct {
	Id       int64  `json:"id"`
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
