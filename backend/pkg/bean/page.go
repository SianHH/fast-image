package bean

import (
	"reflect"
)

var cfg = Cfg{
	MinPage:     1,
	MaxSize:     100,
	DefaultSize: 10,
}

type Cfg struct {
	MinPage     int // 最小分页
	MaxSize     int // 最大条目
	DefaultSize int // 默认条目
}

// SetCfg 设置分页参数
func SetCfg(minPage, maxSize, defaultSize int) {
	cfg = Cfg{
		MinPage:     minPage,
		MaxSize:     maxSize,
		DefaultSize: defaultSize,
	}
}

type PageParam struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

func (p *PageParam) GetOffset() int {
	if p.Page < cfg.MinPage {
		p.Page = cfg.MinPage
	}
	return (p.Page - 1) * p.Size
}

func (p *PageParam) GetLimit() int {
	if p.Size > cfg.MaxSize {
		p.Size = cfg.MaxSize
	}
	if p.Size == 0 {
		p.Size = cfg.DefaultSize
	}
	return p.Size
}

func (p *PageParam) HasPage(total int64) bool {
	return total > int64(p.GetOffset())
}

type resultPage struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

func NewPage(list any, total int64) resultPage {
	// 将nil处理为[]
	of := reflect.ValueOf(list)
	switch of.Kind() {
	case reflect.Slice:
		if of.Len() == 0 {
			list = make([]int, 0)
		}
	default:

	}
	return resultPage{
		List:  list,
		Total: total,
	}
}
