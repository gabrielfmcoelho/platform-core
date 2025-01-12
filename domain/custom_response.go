package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Omits Data field if nil
}
