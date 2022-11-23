package database

import (
	"errors"
	"workshop/src/types"
)

type Database struct {
	url string
}

var DB Database

func Init(url string) {
	DB = Database{url}
}

func (self *Database) Save(key, value string) (int, error) {
	if value == "err" {
		return -1, errors.New("This value is not allowed")
	} else {
		return 1, nil
	}
}

func (self *Database) delete(key, value string) (int, error) {
	if value == "err" {
		return -1, errors.New("This value is not allowed")
	} else {
		return 1, nil
	}
}

func (self *Database) InsertUser(user types.User) (int, error) {
	// Save user

	return 1, nil
}

func (self *Database) InsertBenefit(user types.Benefit) (int, error) {
	// Save user

	return 1, nil
}
