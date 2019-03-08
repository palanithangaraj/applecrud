package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Sample
func TestUnMarshallJson(t *testing.T) {
	jokeIn := JokeType{Type: "success"}
	b, _ := json.Marshal(jokeIn)
	jokeOut := JokeType{}
	if err := json.Unmarshal(b, &jokeIn); err != nil {
		assert.True(t, jokeIn.Type == jokeOut.Type)
	}
}
