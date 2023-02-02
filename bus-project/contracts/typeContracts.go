package contracts

type CreateTypeRequest struct {
	Value string `json:"value" binding:"required"`
}

type CreatedTypeResponse struct {
	Id    uint   `json:"id"`
	Value string `json:"value"`
}

type DeletedTypeResponse struct {
	Id uint `json:"id"`
}

type TypeResponse struct {
	Id    uint   `json:"id"`
	Value string `json:"value"`
}
