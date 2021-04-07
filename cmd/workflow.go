package cmd

import (
	"github.com/spf13/cobra"
)

var workflowCmd = &cobra.Command{
	Use:     "workflow",
	Aliases: []string{"wf"},
	Short:   "Create, update and manage workflows",
	Long:    "Create, update and manage workflows",
	Run:     func(cmd *cobra.Command, args []string) {},
}

// TODO: Add subcommand of workflows
func init() {
}
