package contracts

type CreateBrandRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreatedBrandResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type DeletedBrandResponse struct {
	Id uint `json:"id"`
}

type BrandResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
