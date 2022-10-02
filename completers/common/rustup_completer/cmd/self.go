package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var selfCmd = &cobra.Command{
	Use:   "self",
	Short: "Modify the rustup installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfCmd).Standalone()

	selfCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(selfCmd)
}
