package responses

type BaseResponse struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}
