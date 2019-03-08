package apiclient

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const defaultLogMessage = "applecrud - apiclient.go:"

type ApiClient struct {
	UserAgent  string
	HttpClient *http.Client
}

type DefaultClient struct {
	Client ApiClient
}

type Request struct {
	TimeDelay int
}

func (c *ApiClient) Get(path string, headers http.Header) (Response, error) {
	response := Response{}
	request, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		log.Errorf("Failed to create GET : %v\n", err)
		return response, err
	}
	headers.Set("Accept", "application/json")
	headers.Set("User-Agent", c.UserAgent)
	request.Header = headers
	return c.Do(request)
}

func (c *ApiClient) Do(request *http.Request) (Response, error) {
	resp := Response{}
	request.Close = true
	response, err := c.HttpClient.Do(request)
	if err != nil {
		log.Errorf("Error sending HTTP request to %s: %v\n", request.URL, err)
		return resp, err
	}
	if response.Body != nil {
		resp.Body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Errorf("Do(): Failed to Read the Body of the Response")
			return resp, err
		}
	}
	response.Body.Close()
	resp.StatusCode = StatusCode(response.StatusCode)
	return resp, err
}

type Response struct {
	Body       []byte
	StatusCode StatusCode
}

type ResponseError struct {
	Error string `bson:"error" json:"error"`
}

// Error return error message
func (hr Response) Error() error {
	e := ResponseError{}
	if !hr.StatusCode.IsSuccessful() {
		json.Unmarshal(hr.Body, &e)
		return errors.New(e.Error)
	}
	return nil
}

type StatusCode int

func (s StatusCode) IsSuccessful() bool {
	return s/100 == 2
}
