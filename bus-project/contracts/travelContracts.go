package contracts

type GetTravelBusPropertyDTO struct {
	Id    uint   `json:"id"`
	Value string `json:"value"`
}

type GetTravelSeatDTO struct {
	Id         uint                       `json:"id"`
	Count      int                        `json:"count"`
	Properties []GetTravelSeatPropertyDTO `json:"seat_properties"`
}

type GetTravelSeatPropertyDTO struct {
	Id     uint `json:"id"`
	No     int  `json:"no"`
	Gender bool `json:"gender"`
}

type GetTravelBusDTO struct {
	Id          uint                      `json:"id"`
	PlateNumber string                    `json:"plate_number"`
	BusBrand    string                    `json:"brand"`
	BusModel    string                    `json:"model"`
	Type        string                    `json:"type"`
	Properties  []GetTravelBusPropertyDTO `json:"properties"`
	Seats       []GetTravelSeatDTO        `json:"seats"`
}

type GetTravelResponse struct {
	Id    uint              `json:"id"`
	Fee   float32           `json:"fee"`
	From  string            `json:"from"`
	To    string            `json:"to"`
	Day   string            `json:"day"`
	Time  string            `json:"time"`
	Buses []GetTravelBusDTO `json:"buses"`
}

type BuyTicketRequest struct {
	TravelId uint `json:"travel_id"`
	BusId    uint `json:"bus_id"`
	Gender   bool `json:"gender"`
	No       int  `json:"no"`
}

type CreateTravelRequest struct {
	Fee   float32 `json:"fee"`
	From  string  `json:"from"`
	To    string  `json:"to"`
	Date  string  `json:"date"`
	BusId uint    `json:"bus_id"`
}
