package dtos

type PaginationRequest struct {
	Limit int `form:"limit" json:"limit"`
	Page  int `form:"page" json:"page"`
}

type PaginationResponse[T any] struct {
	Page            int   `json:"page"`
	Limit           int   `json:"limit"`
	TotalItems      int64 `json:"totalItems"`
	TotalPages      int64 `json:"totalPages"`
	HasPreviousPage bool  `json:"hasPreviousPage"`
	HasNextPage     bool  `json:"hasNextPage"`
	Items           *[]*T `json:"items"`
}

func (p *PaginationRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationRequest) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 15
	}
	return p.Limit
}

func (p *PaginationRequest) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
