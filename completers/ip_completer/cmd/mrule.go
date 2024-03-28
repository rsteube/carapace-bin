package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mruleCmd = &cobra.Command{
	Use:   "mrule",
	Short: "rule in multicast routing policy database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mruleCmd).Standalone()

	rootCmd.AddCommand(mruleCmd)
}
