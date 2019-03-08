package db

import (
	"applecrud/pkg/model"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type DataStore interface {
	ReadJoke(name *model.Name, jokeType *model.JokeType) string
}

type Dao struct {
}

//Just mimics the DB layer
func (dao *Dao) ReadJoke(name *model.Name, jokeType *model.JokeType) string {
	if name == nil || jokeType == nil {
		log.Fatalf("Error - reading date")
	}

	return fmt.Sprintf("%s %s's %s", name.Name, name.Surname, jokeType.Value.Joke)
}
