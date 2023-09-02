package db

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/yugo-ibuki/musdok/ent"
)

type Connector interface {
	Init() (*Sqlite, error)
}

type Sqlite struct {
	client *ent.Client
}

func NewSqlite() (*Sqlite, error) {
	return new(Sqlite).init()
}

func (s *Sqlite) init() (*Sqlite, error) {
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		fmt.Println("failed opening connection to sqlite:", err)
		return nil, err
	}
	defer drv.Close()
	client := ent.NewClient(ent.Driver(drv))
	return &Sqlite{client: client}, nil
}
