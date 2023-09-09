package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musdok/ent"
	"github.com/yugo-ibuki/musdok/internal/repository"
	"os"
)

func newAllCmd() *cobra.Command {
	allCmd := &cobra.Command{
		Use:   "all",
		Short: "shows all todos",
		Long:  "shows all todos",
		Run: func(cmd *cobra.Command, args []string) {
			allTodos := all(cmd.Context())
			// shows all todos one by one
			for _, todo := range allTodos {
				fmt.Printf("ID: %d / Title: %s / Desc: %s / Created: %s / Updated: %s\n",
					todo.ID,
					todo.Title,
					todo.Description,
					todo.CreatedAt,
					todo.UpdatedAt,
				)
			}
		},
	}
	return allCmd
}

func all(ctx context.Context) []*ent.Todo {
	repo := repository.NewTodoRp()
	todos, err := repo.Client.Todo.Query().All(ctx)
	if err != nil {
		fmt.Println("failed to query todos: %v", err)
		os.Exit(1)
	}

	return todos
}
