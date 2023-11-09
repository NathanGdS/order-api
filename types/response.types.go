package types

type CustomResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	Messages []string `json:"messages"`
}
