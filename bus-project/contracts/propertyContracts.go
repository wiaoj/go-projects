package contracts

type CreatePropertyRequest struct {
	Value string `json:"value" binding:"required"`
}

type CreatedPropertyResponse struct {
	Id    uint   `json:"id"`
	Value string `json:"value"`
}

type DeletedPropertyResponse struct {
	Id uint `json:"id"`
}

type BusPropertyResponse struct {
	Id    uint   `json:"id"`
	Value string `json:"value"`
}

type PropertyResponse struct {
	Id    uint   `json:"id"`
	Value string `json:"value"`
}
