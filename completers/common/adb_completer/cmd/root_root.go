package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var root_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "restart adbd with root permissions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(root_rootCmd).Standalone()

	rootCmd.AddCommand(root_rootCmd)
}
