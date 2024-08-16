package httpClients

import (
	"net/http"
)

type Response struct {
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	ErrorData ErrorData `json:"error,omitempty"`
}

type ErrorData struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type HttpResponseClients struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       string
}

type CallAPIProducer struct {
	Data       string `json:"data"`
	TopicName  string `json:"topicName"`
	MessageKey string `json:"messageKey"`
}
