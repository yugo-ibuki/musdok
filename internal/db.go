package db

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/yugo-ibuki/musdok/ent"
	"os"
)

const (
	db     = "sqlite3"
	dbName = "musdok.db"
	source = "file:" + dbName + "?_fk=1&cache=shared&_busy_timeout=5000"
)

func Client() (*ent.Client, error) {
	drv, err := sql.Open(db, source)
	if err != nil {
		fmt.Println("failed to open driver: %v", err)
		return nil, err
	}

	return ent.NewClient(ent.Driver(drv)), nil
}

func DatabaseExists() bool {
	path := "./" + dbName
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
