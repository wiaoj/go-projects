package contracts

type CreateBusRequest struct {
	PlateNumber string                `json:"plate_number"`
	SeatsCount  int                   `json:"seats"`
	BusModelId  uint                  `json:"bus_model_id"`
	TypeId      uint                  `json:"type_id"`
	Properties  []CreateBusProperties `json:"properties"`
}

type CreateBusProperties struct {
	Id uint `json:"id"`
}

type CreatedBusResponse struct {
	Id uint `json:"id"`
	// PlateNumber string `json:"plate_number"`
	// Seats       int    `json:"seats"`
	// BusBrand    string `json:"brand"`
	// BusModel    string `json:"model"`
	// Type        string `json:"type"`
	// Properties  any    `json:"properties"`
}
type ListBusResponse struct {
	Id          uint   `json:"id"`
	PlateNumber string `json:"plate_number"`
	Seats       int    `json:"seats"`
	BusBrand    string `json:"brand"`
	BusModel    string `json:"model"`
	Type        string `json:"type"`
	Properties  any    `json:"properties"`
}
type GetByIdBusResponse struct {
	Id          uint   `json:"id"`
	PlateNumber string `json:"plate_number"`
	Seats       int    `json:"seats"`
	BusBrand    string `json:"brand"`
	BusModel    string `json:"model"`
	Type        string `json:"type"`
	Properties  any    `json:"properties"`
}
