package model

// PaginationRequest chứa thông tin phân trang từ query parameters
type PaginationRequest struct {
	Page  int `query:"page"`  // Trang hiện tại (bắt đầu từ 1)
	Limit int `query:"limit"` // Số bản ghi trên mỗi trang
}

// PaginationResponse chứa thông tin phân trang trả về
type PaginationResponse struct {
	CurrentPage int   `json:"current_page"` // Trang hiện tại
	PerPage     int   `json:"per_page"`     // Số bản ghi trên mỗi trang
	TotalPages  int   `json:"total_pages"`  // Tổng số trang
	TotalItems  int64 `json:"total_items"`  // Tổng số bản ghi
	HasNext     bool  `json:"has_next"`     // Có trang kế tiếp hay không
	HasPrev     bool  `json:"has_prev"`     // Có trang trước hay không
}

// PaginatedResponse chứa dữ liệu và thông tin phân trang
type PaginatedResponse struct {
	Data       interface{}        `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

// GetOffset tính toán offset cho database query
func (p *PaginationRequest) GetOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return (p.Page - 1) * p.Limit
}

// Validate và set default values cho pagination request
func (p *PaginationRequest) Validate() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10 // Default limit
	}
	if p.Limit > 100 {
		p.Limit = 100 // Max limit
	}
}

// BuildPaginationResponse tạo response phân trang
func BuildPaginationResponse(req PaginationRequest, totalItems int64) PaginationResponse {
	totalPages := int((totalItems + int64(req.Limit) - 1) / int64(req.Limit))

	return PaginationResponse{
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
		HasNext:     req.Page < totalPages,
		HasPrev:     req.Page > 1,
	}
}
