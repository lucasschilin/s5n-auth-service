package dto

type DefaultError struct {
	Code   int
	Detail string
}

type DefaultDetailResponse struct {
	Detail string `json:"detail"`
}

type DefaultMessageResponse struct {
	Message string `json:"message"`
}
