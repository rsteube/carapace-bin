package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "Print runtime TypeScript declarations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(typesCmd).Standalone()

	rootCmd.AddCommand(typesCmd)
}
