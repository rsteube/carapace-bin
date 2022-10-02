package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var satisfyCmd = &cobra.Command{
	Use:   "satisfy",
	Short: "satisfy causes apt-get to satisfy the given dependency strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(satisfyCmd).Standalone()

	rootCmd.AddCommand(satisfyCmd)
}
