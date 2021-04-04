package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tass-io/cli/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(version.VersionCmd)
}
