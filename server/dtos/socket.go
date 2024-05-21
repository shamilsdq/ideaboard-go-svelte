package dtos

type SocketDto struct {
	Code    string `json:"code" validate:"required,min=0"`
	Content any    `json:"content" validate:"required"`
}

type SocketErrorDto struct {
	Errors []string `json:"errors"`
}
