package appmodel

// Response : Base response
type SuccessResponse struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

// Error : error detail
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors,omitempty"`
}

type PaginationResponsePagination struct {
	Page  *int `json:"page"`
	Total *int `json:"total"`
	Size  *int `json:"size"`
}

type PaginationResponseList struct {
	Pagination *PaginationResponsePagination `json:"pagination"`
	Content    interface{}                   `json:"content"`
}

type PaginationResponse struct {
	List *PaginationResponseList `json:"list"`
}
