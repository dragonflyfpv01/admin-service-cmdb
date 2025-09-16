package model

type InfraComponent struct {
	ID              int    `json:"id" db:"id"`
	Hostname        string `json:"hostname" db:"hostname"`
	DNS             string `json:"dns" db:"dns"`
	Description     string `json:"description" db:"description"`
	PublicInternet  string `json:"public_internet" db:"public_internet"`
	Class           string `json:"class" db:"class"`
	IPAddress       string `json:"ipaddress" db:"ipaddress"`
	Subnet          string `json:"subnet" db:"subnet"`
	Site            string `json:"site" db:"site"`
	ITComponentType string `json:"it_component_type" db:"it_component_type"`
	RequestType     string `json:"request_type" db:"request_type"`
	AppID           string `json:"appid" db:"appid"`
	VLAN            string `json:"vlan" db:"vlan"`
	AppName         string `json:"app_name" db:"app_name"`
	AppOwner        string `json:"app_owner" db:"app_owner"`
	Level           string `json:"level" db:"level"`
	CIOwners        string `json:"ci_owners" db:"ci_owners"`
	IMCM            string `json:"im_cm" db:"im_cm"`
	Status          string `json:"status" db:"status"`
	CreatedAt       string `json:"created_at" db:"created_at"`
	CreateBy        string `json:"create_by" db:"create_by"`
}
