package spec

type APIError struct {
	Code        string `json:"code"`
	HTTPStatus  int    `json:"httpStatus"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}
