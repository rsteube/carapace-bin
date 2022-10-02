package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var jdwpCmd = &cobra.Command{
	Use:   "jdwp",
	Short: "list pids of processes hosting a JDWP transport",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(jdwpCmd).Standalone()

	rootCmd.AddCommand(jdwpCmd)
}
