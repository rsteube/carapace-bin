package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var installTestCmd = &cobra.Command{
	Use:   "install-test",
	Short: "Install package(s) and run tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installTestCmd).Standalone()

	installTestCmd.Flags().Bool("audit", false, "Conduct security audit")
	installTestCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	installTestCmd.Flags().Bool("dry-run", false, "Only report changes")
	installTestCmd.Flags().Bool("fund", false, "Display funding message")
	installTestCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	installTestCmd.Flags().Bool("global-style", false, "Use global layout")
	installTestCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	installTestCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	installTestCmd.Flags().Bool("no-save", false, "Prevents saving to `dependencies`")
	installTestCmd.Flags().StringArray("omit", []string{""}, "Exclude package")
	installTestCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	installTestCmd.Flags().BoolP("save", "S", false, "Package will appear in your `dependencies`")
	installTestCmd.Flags().Bool("save-dev", false, "Package will appear in your `devDependencies`")
	installTestCmd.Flags().BoolP("save-exact", "E", false, "Save exact package version")
	installTestCmd.Flags().Bool("save-optional", false, "Package will appear in your `optionalDependencies`")
	installTestCmd.Flags().Bool("save-peer", false, "Package will appear in your `peerDependencies`")
	installTestCmd.Flags().Bool("save-prod", false, "Package will appear in your `dependencies`.")
	installTestCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(installTestCmd)
	rootCmd.AddCommand(installTestCmd)

	carapace.Gen(installTestCmd).FlagCompletion(carapace.ActionMap{
		"omit": carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(installTestCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return action.ActionPackages(installTestCmd)
		}),
	)
}
