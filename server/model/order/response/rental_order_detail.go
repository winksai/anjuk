package response

type RentalOrderDetail struct {
	HouseTitle   string        `json:"houseTitle"`
	HouseImages  []string      `json:"houseImages"`
	TenantName   string        `json:"tenantName"`
	LandlordName string        `json:"landlordName"`
	RentAmount   float64       `json:"rentAmount"`
	Status       string        `json:"status"`
	SignedAt     string        `json:"signedAt"`
	Contract     *ContractInfo `json:"contract,omitempty"`
}

type ContractInfo struct {
	ContractURL string `json:"contractUrl"`
	Status      string `json:"status"`
	SignTime    string `json:"signTime"`
}
