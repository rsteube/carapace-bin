package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate config files/directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	validateCmd.Flags().String("config-format", "", "Config files are in this format irrespective of their extension.")
	validateCmd.Flags().Bool("quiet", false, "When given, a successful run will produce no output.")
	rootCmd.AddCommand(validateCmd)

	carapace.Gen(validateCmd).FlagCompletion(carapace.ActionMap{
		"config-format": carapace.ActionValues("hcl", "json"),
	})

	carapace.Gen(validateCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
