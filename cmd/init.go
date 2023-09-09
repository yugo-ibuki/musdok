package cmd

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musdok/ent/migrate"
	"github.com/yugo-ibuki/musdok/internal/db"
	"log"
)

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
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
	client, err := db.Client()
	if err != nil {
		log.Fatalf("failed to open driver: %v", err)
		return err
	}
	defer client.Close()

	// Run migration
	if err := client.Schema.Create(context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return err
	}
	return nil
}
