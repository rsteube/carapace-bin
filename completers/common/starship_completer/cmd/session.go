package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Generate random session key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessionCmd).Standalone()

	sessionCmd.Flags().BoolP("help", "h", false, "Prints help information")
	sessionCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(sessionCmd)
}
