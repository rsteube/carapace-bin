package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "disconnect from given TCP/IP device or all",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disconnectCmd).Standalone()

	rootCmd.AddCommand(disconnectCmd)
}
