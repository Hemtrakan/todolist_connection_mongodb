package httpClients

import (
	"bytes"
	"io"
	"net/http"
)

type IHttpClient interface {
	Get(url string) (httpResponse HttpResponseClients, Error error)
	Post(url string, headers map[string]string, payload []byte) (httpResponse HttpResponseClients, Error error)
	Put(url string, headers map[string]string, payload []byte) (httpResponse HttpResponseClients, Error error)
	Delete(url string) (httpResponse HttpResponseClients, Error error)
	DeleteWhitPayload(url string, headers map[string]string, payload []byte) (httpResponse HttpResponseClients, Error error)
}

type httpClient struct {
}

func NewHttpClient() IHttpClient {
	return httpClient{}
}

func (h httpClient) Get(url string) (httpResponse HttpResponseClients, Error error) {
	method := "GET"
	httpResponse, Error = send(method, url, nil, nil)
	return
}

func (h httpClient) Post(url string, headers map[string]string, payload []byte) (httpResponse HttpResponseClients, Error error) {
	method := "POST"
	//if _, ok := headers["Content-Type"]; !ok {
	//	// If not, initialize it with an empty slice
	//	headers["Content-Type"] = []string{}
	//}
	//headers["Content-Type"] = append(headers["Content-Type"], "application/json")

	httpResponse, Error = send(method, url, headers, payload)
	return
}

func (h httpClient) Put(url string, headers map[string]string, payload []byte) (httpResponse HttpResponseClients, Error error) {
	method := "PUT"
	//if _, ok := headers["Content-Type"]; !ok {
	//	// If not, initialize it with an empty slice
	//	headers["Content-Type"] = []string{}
	//}
	//headers["Content-Type"] = append(headers["Content-Type"], "application/json")
	httpResponse, Error = send(method, url, headers, payload)
	return
}

func (h httpClient) Delete(url string) (httpResponse HttpResponseClients, Error error) {
	method := "DELETE"
	httpResponse, Error = send(method, url, nil, nil)
	return
}

func (h httpClient) DeleteWhitPayload(url string, headers map[string]string, payload []byte) (httpResponse HttpResponseClients, Error error) {
	method := "DELETE"
	//if _, ok := headers["Content-Type"]; !ok {
	//	// If not, initialize it with an empty slice
	//	headers["Content-Type"] = []string{}
	//}
	//headers["Content-Type"] = append(headers["Content-Type"], "application/json")
	httpResponse, Error = send(method, url, headers, payload)
	return
}

func send(method, url string, headers map[string]string, payload []byte) (httpResponse HttpResponseClients, Error error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		Error = err
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		Error = err
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		Error = err
		return
	}

	httpResponse.Status = resp.Status
	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Body = string(responseBody)
	return
}
