package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var schema_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check a Pulumi package schema for errors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schema_checkCmd).Standalone()
	schemaCmd.AddCommand(schema_checkCmd)
}
