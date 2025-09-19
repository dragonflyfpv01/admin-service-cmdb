package req

// ReqUpdateStatus request để cập nhật status với cả ID và hostname
type ReqUpdateStatus struct {
	ID        int    `json:"id" validate:"required"`
	Hostname  string `json:"hostname" validate:"required"`
	NewStatus string `json:"new_status" validate:"required"`
}
