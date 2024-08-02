package response

import "net/http"

const (
	statusOK    = http.StatusOK
	statusError = "error"
)

type Response struct {
	Status int    `json:"statusCode"`
	Msg    string `json:"msg,omitempty"`
}

func OK() Response {
	return Response{
		Status: statusOK,
	}
}

func Error(StatusCode int, msg string) Response {
	return Response{
		Status: StatusCode,
		Msg:    msg,
	}
}
