package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logcatCmd = &cobra.Command{
	Use:   "logcat",
	Short: "show device log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logcatCmd).Standalone()

	rootCmd.AddCommand(logcatCmd)
}
