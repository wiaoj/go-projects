package contracts

type ListDataResponse struct {
	//Message string      `json:"message" xml:"message"`
	Status Status      `json:"status" xml:"status"`
	Datas  interface{} `json:"data" xml:"data"`
}
