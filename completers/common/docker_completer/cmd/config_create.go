package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a config from a file or STDIN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_createCmd).Standalone()

	config_createCmd.Flags().StringP("label", "l", "", "Config labels")
	config_createCmd.Flags().String("template-driver", "", "Template driver")
	configCmd.AddCommand(config_createCmd)

	carapace.Gen(config_createCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionFiles(),
	)
}
