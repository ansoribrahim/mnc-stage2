package data

type Response struct {
	Status  string       `json:"status"`
	Result  *interface{} `json:"result,omitempty"`
	Message *string      `json:"message,omitempty"`
}
