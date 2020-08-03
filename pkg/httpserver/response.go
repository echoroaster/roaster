package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

func DataResponse(value interface{}) map[string]interface{} {
	return map[string]interface{}{
		"data": value,
	}
}

type ErrorResponse struct {
	Error ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
