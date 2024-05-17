package dtos

type SocketDto struct {
	Code    string `json:"code"`
	Content any    `json:"content"`
}

type SocketErrorDto struct {
	Error string `json:"error"`
}
