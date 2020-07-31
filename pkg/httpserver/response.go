package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

type ErrorResponse struct {
	Error ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
