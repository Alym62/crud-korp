package pkg

type PageResponse[T any] struct {
	List       []T `json:"list"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}
