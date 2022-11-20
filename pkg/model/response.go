package model

type MessageResponse struct {
	ErrorCode int64  `json:"errorCode"`
	Message   string `json:"message"`
}
