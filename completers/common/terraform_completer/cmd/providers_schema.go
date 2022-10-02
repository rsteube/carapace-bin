package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var providers_schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Show schemas for the providers used in the configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providers_schemaCmd).Standalone()
	providers_schemaCmd.Flags().BoolS("json", "json", false, "Prints out a json representation")

	providersCmd.AddCommand(providers_schemaCmd)
}
