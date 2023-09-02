package cmd

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musdok/ent"
	"github.com/yugo-ibuki/musdok/ent/migrate"
	"log"
)

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init DB",
		Short: "Initialize the database",
		Long:  `Initialize the database`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("start initialize database")

			// DB initialization process
			if err := initDb(); err != nil {
				return err
			}

			fmt.Println("finish initialize database")
			return nil
		},
	}
}

func initDb() error {
	drv, err := sql.Open("sqlite3", "file:ent.db?_fk=1&cache=shared&_busy_timeout=5000")
	if err != nil {
		log.Fatalf("failed to open driver: %v", err)
		return err
	}
	defer drv.Close()

	client := ent.NewClient(ent.Driver(drv))

	// Run migration
	if err := client.Schema.Create(context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return err
	}
	return nil
}
