package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tass-io/cli/pkg/route"
)

var routeCmd = &cobra.Command{
	Use:     "route",
	Aliases: []string{"rt"},
	Short:   "Create, update and manage routes",
	Long:    "Create, update and manage routes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	routeCmd.AddCommand(route.CreateCmd)
	routeCmd.AddCommand(route.DeleteCmd)
	routeCmd.AddCommand(route.ListCmd)
	routeCmd.AddCommand(route.UpdateCmd)
}
