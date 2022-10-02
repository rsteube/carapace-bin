package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var json2cCmd = &cobra.Command{
	Use:   "json2c",
	Short: "Creates a keymap.c from a QMK Configurator export",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(json2cCmd).Standalone()

	json2cCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	json2cCmd.Flags().StringP("output", "o", "", "File to write to")
	json2cCmd.Flags().BoolP("quiet", "q", false, "Quiet mode, only output error messages")
	rootCmd.AddCommand(json2cCmd)

	carapace.Gen(json2cCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(json2cCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
