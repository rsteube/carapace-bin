package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var refsCmd = &cobra.Command{
	Use:   "refs",
	Short: "Start up in refs view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refsCmd).Standalone()

	rootCmd.AddCommand(refsCmd)
}
