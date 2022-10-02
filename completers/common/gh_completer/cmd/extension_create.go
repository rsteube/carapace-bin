package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var extension_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_createCmd).Standalone()
	extension_createCmd.Flags().String("precompiled", "", "Create a precompiled extension. Possible values: go, other")
	extensionCmd.AddCommand(extension_createCmd)

	carapace.Gen(extension_createCmd).FlagCompletion(carapace.ActionMap{
		"precompiled": carapace.ActionValues("go", "other"),
	})
}
