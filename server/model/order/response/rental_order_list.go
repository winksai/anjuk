package response

type RentalOrderListItem struct {
	ID           int     `json:"id"`
	HouseTitle   string  `json:"houseTitle"`
	TenantName   string  `json:"tenantName"`
	LandlordName string  `json:"landlordName"`
	RentStart    string  `json:"rentStart"`
	RentEnd      string  `json:"rentEnd"`
	RentAmount   float64 `json:"rentAmount"`
	Deposit      float64 `json:"deposit"`
	Status       string  `json:"status"`
	SignedAt     string  `json:"signedAt"`
	CreatedAt    string  `json:"createdAt"`
}
