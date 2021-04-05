package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tass-io/cli/pkg/function"
)

var functionCmd = &cobra.Command{
	Use:     "function",
	Aliases: []string{"fn"},
	Short:   "Create, update and manage functions",
	Long:    "Create, update and manage functions",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	functionCmd.AddCommand(function.CreateCmd)
}
