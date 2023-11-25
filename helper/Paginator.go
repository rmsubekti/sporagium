package helper

import (
	"math"

	"gorm.io/gorm"
)

type Paginator struct {
	Limit      int    `query:"limit" form:"limit, omitempty" json:"limit"`
	Page       int    `query:"page" form:"page, omitempty" json:"page"`
	Sort       string `query:"sort" form:"sort, omitempty" json:"sort"`
	Order      string `query:"order" form:"order" json:"order"`
	Query      string `query:"query" form:"query, omitempty" json:"query"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Rows       any    `json:"rows"`
	Offset     int    `json:"-"`
}

func (p *Paginator) Scopes() func(db *gorm.DB) *gorm.DB {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset).Limit(p.Limit).Order(p.Order + " " + p.Sort)
	}
}

func (p *Paginator) SetCount(countCondition *gorm.DB) (err error) {
	var total int64
	if err = countCondition.Count(&total).Error; err != nil {
		return
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	p.Offset = (p.Page - 1) * p.Limit

	if len(p.Sort) < 1 {
		p.Sort = "asc"
	}

	if len(p.Order) < 1 {
		p.Order = "id"
	}

	p.TotalRows = total
	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.Limit)))

	return
}
func (p *Paginator) SetNCount(total int64) (err error) {

	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	p.Offset = (p.Page - 1) * p.Limit

	if len(p.Sort) < 1 {
		p.Sort = "asc"
	}

	if len(p.Order) < 1 {
		p.Order = "id"
	}

	p.TotalRows = total
	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.Limit)))

	return
}

func (p *Paginator) Paginate(rows any) {
	p.Rows = rows
}
