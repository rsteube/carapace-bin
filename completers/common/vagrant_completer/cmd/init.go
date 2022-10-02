package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes a new Vagrant environment by creating a Vagrantfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("box-version", "", "Version of the box to add")
	initCmd.Flags().BoolP("force", "f", false, "Overwrite existing Vagrantfile")
	initCmd.Flags().BoolP("minimal", "m", false, "Use minimal Vagrantfile template (no help comments). Ignored with --template")
	initCmd.Flags().String("output", "", "Output path for the box. '-' for stdout")
	initCmd.Flags().String("template", "", "Path to custom Vagrantfile template")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"box-version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionCloudBoxVersions(c.Args[0], "")
			}
			return carapace.ActionValues("")
		}),
		"output":   carapace.ActionFiles(),
		"template": carapace.ActionFiles(),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
	)
}
