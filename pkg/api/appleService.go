package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"applecrud/pkg/apiclient"
	"applecrud/pkg/config"
	"applecrud/pkg/db"
	"applecrud/pkg/model"
)

const defaultPreLog = "TNG - applecrud - appleService.go "

//TODO: Worker pool

//HealthCheck defines a health check route
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetNamedJoke(dao db.DataStore, config config.Config) http.Handler {
	//TODO: Additional validations
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		pipe := make(chan []byte)
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		client := apiclient.ApiClient{UserAgent: fmt.Sprintf("apple Service"), HttpClient: http.DefaultClient}
		if client.HttpClient == nil {
			log.Printf("no http client found")
		}

		//Call Async
		//TODO: Could be a for range loop with a map of resultype vs. url
		//Could be cached or, generated on the fly
		//Did not want to use the long spinning "select" for this requirement
		url := "http://uinames.com/api/"
		go get(pipe, client, url)
		//We could make this on the fly if we want - may not be out of the scope
		url = "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=\\[nerdy\\]"
		go get(pipe, client, url)

		//Get the values in order
		var name model.Name
		json.Unmarshal(<-pipe, &name)
		var jokeType model.JokeType
		json.Unmarshal(<-pipe, &jokeType)

		result := ""
		if result = dao.ReadJoke(&name, &jokeType); result == "" {
			log.Errorf(defaultPreLog+": GetNamedItem: ERROR : Name: %s : JokeType: %s\n", name.Name, jokeType.Type)
			respondWithError(w, 401, "No Result")
			return
		}
		data, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
}

func get(pipe chan []byte, client apiclient.ApiClient, url string) {
	log.Printf("Sending API - %s\n", url)
	response, err := client.Get(url, http.Header{})
	log.Printf("Sending API - Response: %v, err:%v  url: %s\n", string(response.Body), err, url)
	if err == nil {
		pipe <- response.Body
	}
}

//Utility function to convert an error message into a JSON response.
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

//Utility function to convert the payload into a JSON response.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
