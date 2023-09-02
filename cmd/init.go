package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init DB",
		Short: "Initialize the database",
		Long:  `Initialize the database`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("init database")
			return nil
		},
	}
}
