package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/virtualbox"
	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "packages a running vagrant environment into a box",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packageCmd).Standalone()

	packageCmd.Flags().String("base", "", "Name of a VM in VirtualBox to package as a base box (VirtualBox Only)")
	packageCmd.Flags().String("include", "", "Comma separated additional files to package with the box")
	packageCmd.Flags().String("info", "", "Path to a custom info.json file containing additional box information")
	packageCmd.Flags().String("output", "", "Name of the file to output")
	packageCmd.Flags().String("vagrantfile", "", "Vagrantfile to package with the box")
	rootCmd.AddCommand(packageCmd)

	carapace.Gen(packageCmd).FlagCompletion(carapace.ActionMap{
		"base": virtualbox.ActionMachines(),
		"include": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"info":        carapace.ActionFiles(),
		"output":      carapace.ActionFiles(),
		"vagrantfile": carapace.ActionFiles(),
	})
}
