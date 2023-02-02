package contracts

type CreateModelRequest struct {
	BrandId uint   `json:"brand_id" bindig:"required"`
	Value   string `json:"value" binding:"required"`
}

type CreatedModelResponse struct {
	Id      uint   `json:"id"`
	BrandId string `json:"brand_id"`
	Value   string `json:"value"`
}

type DeletedModelResponse struct {
	Id uint `json:"id"`
}

type ModelResponse struct {
	Id      uint   `json:"id"`
	BrandId uint   `json:"brand_id"`
	Value   string `json:"value"`
}
