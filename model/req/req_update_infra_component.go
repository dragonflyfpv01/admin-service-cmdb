package req

// ReqUpdateInfraComponent request để cập nhật thông tin infra component
// Không bao gồm status và create_by (không được phép sửa)
// created_at sẽ được cập nhật tự động trong handler
type ReqUpdateInfraComponent struct {
	ID              int    `json:"id" validate:"required"`
	Hostname        string `json:"hostname" validate:"required"`
	DNS             string `json:"dns"`
	Description     string `json:"description"`
	PublicInternet  string `json:"public_internet"`
	Class           string `json:"class"`
	IPAddress       string `json:"ipaddress"`
	Subnet          string `json:"subnet"`
	Site            string `json:"site"`
	ITComponentType string `json:"it_component_type"`
	RequestType     string `json:"request_type"`
	AppID           string `json:"appid"`
	VLAN            string `json:"vlan"`
	AppName         string `json:"app_name"`
	AppOwner        string `json:"app_owner"`
	Level           string `json:"level"`
	CIOwners        string `json:"ci_owners"`
	IMCM            string `json:"im_cm"`
}
