package resource

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, p *Pagination) (*ListResponse, error) {
	pg := ListParams{
		Limit:  p.Limit,
		Offset: (p.Page - 1) * p.Limit,
	}
	res, err := s.repo.ListAvailableResources(ctx, &pg)
	if err != nil {
		return nil, err
	}
	response := ListResponse{
		Pagination: Pagination{
			Page:  p.Page,
			Limit: p.Limit,
		},
		Data: res,
	}
	return &response, nil
}


func (s *Service) Get(ctx context.Context, id int64) (*ListResource, error) {
	res, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	list := &ListResource{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Location:    res.Location,
		IsActive:    res.IsActive,
		CreatedAt:   res.CreatedAt,
	}
	return list, nil


}
