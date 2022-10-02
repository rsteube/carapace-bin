package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pip"
	"github.com/spf13/cobra"
)

var freezeCmd = &cobra.Command{
	Use:   "freeze",
	Short: "Output installed packages in requirements format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freezeCmd).Standalone()

	freezeCmd.Flags().Bool("all", false, "Do not skip pip, distribute, setuptools, wheel")
	freezeCmd.Flags().String("exclude", "", "Exclude specified package from the output")
	freezeCmd.Flags().Bool("exclude-editable", false, "Exclude editable package from output.")
	freezeCmd.Flags().StringP("find-links", "f", "", "URL for finding packages, which will be added to the output.")
	freezeCmd.Flags().BoolP("local", "l", false, "If in a virtualenv that has global access, do not output globally-installed packages.")
	freezeCmd.Flags().StringArray("path", []string{}, "Restrict to the specified installation path for listing packages")
	freezeCmd.Flags().StringArrayP("requirement", "r", []string{}, "Use the order in the given requirements file and its comments when generating output.")
	freezeCmd.Flags().Bool("user", false, "Only output packages installed in user-site.")
	rootCmd.AddCommand(freezeCmd)

	carapace.Gen(freezeCmd).FlagCompletion(carapace.ActionMap{
		"exclude":     pip.ActionInstalledPackages(),
		"path":        carapace.ActionDirectories(),
		"requirement": carapace.ActionFiles(),
	})
}
