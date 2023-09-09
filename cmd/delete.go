package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musdok/internal/repository"
	"os"
	"strconv"
)

func newDeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a todo",
		Long:  "Delete a todo",
		Run: func(cmd *cobra.Command, args []string) {
			idFlg := cmd.Flag("id").Value.String()

			id, err := strconv.Atoi(idFlg)
			if err != nil {
				fmt.Println("failed to convert id to int: %v", err)
				os.Exit(1)
			}
			delete(cmd.Context(), id)
			fmt.Printf("delete success! id: %d", id)
		},
	}
	addDeleteFlgs(deleteCmd)
	return deleteCmd
}

func addDeleteFlgs(c *cobra.Command) {
	c.Flags().String("id", "", "TODO id")
}

func delete(ctx context.Context, id int) {
	repo := repository.NewTodoRp()
	defer repo.Close()
	repo.Delete(ctx, id)
}
