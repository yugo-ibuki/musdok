package db

import (
	"os"
)

const (
	path = "./musdok.db"
)

func DatabaseExists() bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
