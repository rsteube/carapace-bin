package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var root_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Display npm root",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(root_rootCmd).Standalone()
	root_rootCmd.Flags().BoolP("global", "g", false, "operate globally")

	rootCmd.AddCommand(root_rootCmd)
}
