package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display system-wide information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_infoCmd).Standalone()

	system_infoCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	systemCmd.AddCommand(system_infoCmd)

	rootAlias(system_infoCmd, func(cmd *cobra.Command, isAlias bool) {})
}
