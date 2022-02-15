package controller

type HTTPError struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"process successfully"`
}
