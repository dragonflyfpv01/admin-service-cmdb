package req

// ReqCreateInfraComponent request model for creating new infra component
type ReqCreateInfraComponent struct {
	Hostname        string `json:"hostname" validate:"required"`
	DNS             string `json:"dns" validate:"required"`
	Description     string `json:"description" validate:"required"`
	PublicInternet  string `json:"public_internet" validate:"required"`
	Class           string `json:"class" validate:"required"`
	IPAddress       string `json:"ipaddress" validate:"required"`
	Subnet          string `json:"subnet" validate:"required"`
	Site            string `json:"site" validate:"required"`
	ITComponentType string `json:"it_component_type" validate:"required"`
	RequestType     string `json:"request_type" validate:"required"`
	AppID           string `json:"appid" validate:"required"`
	VLAN            string `json:"vlan" validate:"required"`
	AppName         string `json:"app_name" validate:"required"`
	AppOwner        string `json:"app_owner" validate:"required"`
	Level           string `json:"level" validate:"required"`
	CIOwners        string `json:"ci_owners" validate:"required"`
	IMCM            string `json:"im_cm" validate:"required"`
}
