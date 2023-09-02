package cmd

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init DB",
		Short: "Initialize the database",
		Long:  `Initialize the database`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("start initialize database")

			// DB initialization process
			//if err := initDb(); err != nil {
			//	return err
			//}

			fmt.Println("finish initialize database")
			return nil
		},
	}
}

//func initDb() error {

//createTableQuery := `
//CREATE TABLE IF NOT EXISTS users (
//	id INTEGER PRIMARY KEY AUTOINCREMENT,
//	name TEXT NOT NULL,
//	age INTEGER
//)
//`
//_, err = db.Exec(createTableQuery)
//if err != nil {
//	log.Fatalf("テーブル作成に失敗: %v", err)
//}

//fmt.Println("start initialize database")
//return nil
//}
