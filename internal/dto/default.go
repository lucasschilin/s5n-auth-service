package dto

type ControllerError struct {
	Code   int
	Detail string
}

type DefaultDetailResponse struct {
	Detail string `json:"detail"`
}

type DefaultMessageResponse struct {
	Message string `json:"message"`
}
