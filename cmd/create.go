package cmd

import (
	"entgo.io/ent"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musdok/internal"
	"os"
)

func newCreateCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new todo",
		Long:  `Create a new todo`,
		Run: func(cmd *cobra.Command, args []string) {
			no := db.DatabaseExists()
			if no {
				fmt.Println("`musdok.db` already exists.")
				os.Exit(1)
			}

			createTodo()
		},
	}
	return createCmd
}

func createTodo() {
	ent.Schema{}.Fields()
}
