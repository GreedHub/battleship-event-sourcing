package domain

type ApiResponse struct{
	Status uint32 `json:"status"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}

const (
	INTERNAL_SERVER_ERROR = "internal_server_error"
	OK = "ok"
)