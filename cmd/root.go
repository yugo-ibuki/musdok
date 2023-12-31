package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "musdok",
		Short: "musdok is a CLI to add and search for the todo list",
		Long:  "musdok is a CLI to add and search for the todo list",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Help())
		},
	}
	cmd.AddCommand(newInitCmd())
	cmd.AddCommand(newAllCmd())
	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newUpdateCmd())
	cmd.AddCommand(newDeleteCmd())
	return cmd
}

func Execute() int {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
