package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musdok/ent"
	"github.com/yugo-ibuki/musdok/internal/db"
	"github.com/yugo-ibuki/musdok/internal/repository"
	"os"
)

func newCreateCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new todo",
		Long:  "Create a new todo",
		Run: func(cmd *cobra.Command, args []string) {
			ttl := cmd.Flag("title").Value.String()
			dsp := cmd.Flag("desc").Value.String()

			//check if the database exists
			yes := db.DatabaseExists()
			if !yes {
				fmt.Println("Do `init` command first to create db")
				os.Exit(1)
			}

			id := create(cmd.Context(), ent.Todo{
				Title:       ttl,
				Description: dsp,
			})
			fmt.Printf("create success! id: %d", id)
		},
	}
	addCreateFlgs(createCmd)
	return createCmd
}

func addCreateFlgs(c *cobra.Command) {
	c.Flags().String("title", "", "TODO title")
	c.Flags().String("desc", "", "TODO description")
}

func create(ctx context.Context, todo ent.Todo) int {
	repo := repository.NewTodoRp()
	defer repo.Close()
	id := repo.Create(ctx, todo)
	return id
}
