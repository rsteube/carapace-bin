package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Deprecated command. Use 'debug' instead.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	rootCmd.AddCommand(runCmd)
}
