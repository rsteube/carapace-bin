package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var jsonschemaCmd = &cobra.Command{
	Use:     "jsonschema",
	Short:   "Outputs nFPM's JSON schema",
	Aliases: []string{"schema"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(jsonschemaCmd).Standalone()
	jsonschemaCmd.Flags().StringP("output", "o", "-", "where to save the json schema")
	rootCmd.AddCommand(jsonschemaCmd)

	carapace.Gen(jsonschemaCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
