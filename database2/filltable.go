package database2

type ReadModel struct {
	IdExternalApp uint64 `json:"ExternalAppID"`
	IdProduct     int    `json:"CustomerMid"`
	IdSuperUser   string `json:"CustomerEmail"`
	Description   uint64 `json:"ExternalAppID"`
	CustomerMid   uint64 `json:"SuperuserID"`
	CustomerEmail bool   `json:"Activated"`
	Activated     bool   `json:"Activated"`
}
