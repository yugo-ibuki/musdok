package cmd

import (
	"entgo.io/ent"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new todo",
		Long:  `Create a new todo`,
		Run: func(cmd *cobra.Command, args []string) {
			createTodo()
		},
	}
	return createCmd
}

func createTodo() {
	ent.Schema{}.Fields()
}
