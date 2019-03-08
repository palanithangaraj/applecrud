package db

import (
	"errors"

	"applecrud/pkg/model"
)

type MockDao struct {
}

//InitializeMock creates a Mock Dao instance that can be used in unit tests
//where the interface is used for DI
func InitializeMock() *MockDao {
	return &MockDao{}
}

func (dao *MockDao) GetNamedJoke(joke model.JokeType) error {
	var err error
	if joke.Type != "success" {
		err = errors.New("not found")
	}
	return err
}
