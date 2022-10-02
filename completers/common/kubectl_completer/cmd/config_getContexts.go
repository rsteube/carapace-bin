package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_getContextsCmd = &cobra.Command{
	Use:   "get-contexts",
	Short: "Describe one or many contexts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getContextsCmd).Standalone()

	config_getContextsCmd.Flags().Bool("no-headers", false, "When using the default or custom-column output format, don't print headers (default print headers).")
	config_getContextsCmd.Flags().StringP("output", "o", "", "Output format. One of: name")
	configCmd.AddCommand(config_getContextsCmd)

	carapace.Gen(config_getContextsCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("name"),
	})
}
