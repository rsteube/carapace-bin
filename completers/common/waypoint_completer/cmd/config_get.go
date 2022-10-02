package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get config variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()

	config_getCmd.Flags().Bool("json", false, "Output in JSON The default is false.")
	config_getCmd.Flags().String("label", "", "Labels to set for this operation.")
	config_getCmd.Flags().Bool("raw", false, "Output in key=val The default is false.")
	config_getCmd.Flags().Bool("runner", false, "Show configuration that is set on runners.")

	addGlobalOptions(config_getCmd)

	configCmd.AddCommand(config_getCmd)
}
