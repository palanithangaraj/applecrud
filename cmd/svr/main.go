package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"applecrud/pkg/api"
	"applecrud/pkg/config"
	"applecrud/pkg/db"
	"applecrud/pkg/util"
)

//PRELOG is the default prefix used for logging before the prefix from the config file is initialized.
const defaultPreLog = "TNG - applecrud - main.go: "

func main() {
	//Set appropriately
	runtime.GOMAXPROCS(2)

	//Load configs
	configPath := os.Getenv("CONFIG_FILE_PATH")
	if configPath == "" {
		configPath = "./config.json"
	}
	config, err := config.GetConfig(configPath)
	if err != nil {
		log.Tracef("%s Fatal config error: %v", defaultPreLog, err)
		log.Fatalf("%s Encountered fatal error initializing configuration, service being terminated", defaultPreLog)
	}

	util.SetLoggingFormat()
	log.Infof("%v main.go - BEGIN", defaultPreLog)
	(&config).Print() // write config printout if level = trace

	//Init datastore/Models
	db := db.Dao{}

	//create API routes
	r := mux.NewRouter().StrictSlash(true)

	//Handlers
	r.HandleFunc("/health", api.HealthCheck).Methods("GET")
	//Already handles the concurrency. 
	//
	//Ran the following in muti teminals
	//for ((i=1;i<=1000;i++)); do curl -v --header "Connection: keep-alive" "localhost:5000"; done
	//Tested with postman as well
	//TODO:
	//Could be able to customize accepting the connections and load balance with go-routines
	//for more concurrency.
	r.Handle("/", api.GetNamedJoke(&db, config)).Methods("GET")

	//start web server on given port with CORS enabled
	log.Debugf("Server listening on port %s\n", config.Port)
	log.Infof("%v main.go - END", defaultPreLog)
	log.Fatal(http.ListenAndServe(":"+config.Port, cors.Default().Handler(r)))
}
