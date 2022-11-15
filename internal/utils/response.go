package utils

type ResponseBody struct {
	Success bool `json:"success"`
	Errors  any  `json:"errors,omitempty"`
	Data    any  `json:"data,omitempty"`
}
