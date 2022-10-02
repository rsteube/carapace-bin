package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show log output from an application's current deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	addGlobalOptions(logsCmd)

	rootCmd.AddCommand(logsCmd)
}
