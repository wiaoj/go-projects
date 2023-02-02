package contracts

type SingleResponse struct {
	Message string `json:"message" xml:"message"` //form:"message"
	Status  Status `json:"status" xml:"status"`
}

type SingleDataResponse struct {
	Message string      `json:"message" xml:"message"`
	Status  Status      `json:"status" xml:"status"`
	Data    interface{} `json:"data" xml:"data"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Status string

const (
	SuccessStatus Status = "success"
	ErrorStatus   Status = "failed"
)
