package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tass-io/cli/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of the Tass.",
	Long:  "Version information of the Tass.",
	Run:   version.GetVersionInfo,
}

func init() {}
