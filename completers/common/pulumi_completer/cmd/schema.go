package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Analyze package schemas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemaCmd).Standalone()
	rootCmd.AddCommand(schemaCmd)
}
