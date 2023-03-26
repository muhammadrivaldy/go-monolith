package models

import (
	"math"
)

type (
	// Pagination is a object for pagination query
	Pagination struct {
		Page  int `validate:"required,min=1"`
		Limit int `validate:"required,min=1"`
	}

	// PaginationResponse ..
	PaginationResponse struct {
		Page      int `json:"page"`
		Limit     int `json:"limit"`
		TotalPage int `json:"total_page"`
		TotalData int `json:"total_data"`
	}
)

func (p Pagination) GetOffset() int {

	if p.Page <= 1 {
		return 0
	}

	page := p.Page - 1

	return page * p.Limit

}

func (p *PaginationResponse) SetTotalPage() {

	if p.TotalData <= p.Limit {
		p.TotalPage = 1
		return
	}

	totalPage := float64(p.TotalData) / float64(p.Limit)

	p.TotalPage = int(math.Ceil(totalPage))

}
