package config

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strings"
)

//PRELOG is the default prefix the Config struct uses for logging before it initializes the prefix from the config file.
const defaultPreLog = "TNG - applecrud - config.go: "

//Config models the configuration data for the api
type Config struct {
	Port     string    `json:"port"`
	PreLog   string    `json:"preLog"`
	LogLevel log.Level `json:"log-level"`
}

//GetConfig gets the configuration values for the api using the file in the supplied configPath. If the file is not found,
//or if the file is missing certain required config value, the function initializes default config values.
func GetConfig(configPath string) (Config, error) {
	c := Config{}
	var err error
	if _, err = os.Stat(configPath); os.IsNotExist(err) { // search localhost path for config
		configPath = "./config.json"

		if _, err = os.Stat(configPath); os.IsNotExist(err) { // search ENV variables for config
			c, err = loadConfigFromENV(c)
			if err != nil {
				return c, err
			}
			configPath = ""
		}
	}

	if configPath != "" { // configurations path is set to "" if no configuration is found
		c, err = loadConfigFromFile(c, configPath)
	}

	log.SetLevel(c.LogLevel)
	return c, err
}

func loadConfigFromENV(c Config) (Config, error) {
	var err error

	// set log level first before any errors are returned below
	if level, err := fetchEnvVariable("LOG_LEVEL"); err != nil {
		c.LogLevel, err = log.ParseLevel(level)
		if err != nil {
			c.LogLevel = log.ErrorLevel
			err = nil
		}
	}

	c.Port = os.Getenv("PORT")
	c.PreLog = os.Getenv("PRELOG")
	if c.Port == "" || c.PreLog == "" {
		err = errors.New("1 or more ENV variables were empty")
	}
	return c, err
}

//if the config loaded from the file errors, no defaults will be loaded and the app will exit.
func loadConfigFromFile(c Config, configPath string) (Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Errorf(defaultPreLog+"Error opening config file: %v", err)
	} else {
		defer file.Close()
		var conf Config
		err = json.NewDecoder(file).Decode(&conf)
		if err != nil {
			log.Errorf(defaultPreLog+"Error decoding config file: %v", err)
		} else {
			c.Port = conf.Port
			c.PreLog = conf.PreLog
			c.LogLevel = conf.LogLevel
		}
	}

	return c, err
}

func decodeURL(encodedURL string) (string, error) {
	url, err := base64.StdEncoding.DecodeString(encodedURL)
	if err != nil {
		log.Errorf(defaultPreLog+"Error decoding db url: %v", err)
	}

	return string(url), err
}

func fetchEnvVariable(key string) (val string, err error) {
	val, found := os.LookupEnv(key)
	if !found {
		err = fmt.Errorf("environment variable %s not set", key)
	}
	return val, err
}
func (c *Config) Print() {
	result := make(map[string]string)
	n := reflect.ValueOf(c).Elem()
	for i := 0; i < n.NumField(); i++ {
		tag := n.Type().Field(i).Tag.Get("json")
		value := n.Field(i).Interface()
		if strings.ToLower(string(tag)) == "databaseurl" {
			value = "*************"
		}
		result[tag] = " " + fmt.Sprintln(value)
	}
	log.Tracef("CURRENT CONFIG: \n %v", result)
}
