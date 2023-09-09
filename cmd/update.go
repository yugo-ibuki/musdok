package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musdok/ent"
	"github.com/yugo-ibuki/musdok/internal/db"
	"github.com/yugo-ibuki/musdok/internal/repository"
	"os"
	"strconv"
)

func newUpdateCmd() *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update a new todo",
		Long:  "Update a new todo",
		Run: func(cmd *cobra.Command, args []string) {
			idFlg := cmd.Flag("id").Value.String()
			ttl := cmd.Flag("title").Value.String()
			dsp := cmd.Flag("desc").Value.String()

			id, err := strconv.Atoi(idFlg)
			if err != nil {
				fmt.Println("failed to convert id to int: %v", err)
				os.Exit(1)
			}

			fmt.Println(ttl, dsp)

			//check if the database exists
			yes := db.DatabaseExists()
			if !yes {
				fmt.Println("Do `init` command first to create db")
				os.Exit(1)
			}

			update(cmd.Context(), id, ent.Todo{
				Title:       ttl,
				Description: dsp,
			})
		},
	}
	addUpdateFlgs(updateCmd)
	return updateCmd
}

func addUpdateFlgs(c *cobra.Command) {
	c.Flags().String("id", "", "todo id")
	c.Flags().String("title", "", "TODO title")
	c.Flags().String("desc", "", "TODO description")
}

func update(ctx context.Context, id int, todo ent.Todo) {
	repo := repository.NewTodoRp()
	defer repo.Close()
	repo.Update(ctx, id, todo)
}
