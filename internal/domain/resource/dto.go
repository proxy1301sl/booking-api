package resource

import "time"

type Resource struct {
	ID          int64
	Name        string
	Description string
	Location    string
	IsActive    bool
	CreatedAt   time.Time
}

type ListResponse struct {
	Pagination Pagination
	Data       []Resource
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListParams struct {
	Limit  int
	Offset int
}
