package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import keys or certificates from a PKCS file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
