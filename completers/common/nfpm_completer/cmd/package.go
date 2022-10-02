package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:     "package",
	Short:   "Creates a package based on the given the given config file and flags",
	Aliases: []string{"pkg", "p"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packageCmd).Standalone()
	packageCmd.Flags().StringP("config", "f", "nfpm.yaml", "config file to be used")
	packageCmd.Flags().StringP("packager", "p", "", "which packager implementation to use [apk|deb|rpm]")
	packageCmd.Flags().StringP("target", "t", "", "where to save the generated package (filename, folder or empty for current folder)")
	rootCmd.AddCommand(packageCmd)

	carapace.Gen(packageCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionFiles(),
		"packager": carapace.ActionValues("apk", "deb", "rpm"),
		"target":   carapace.ActionFiles(),
	})
}
