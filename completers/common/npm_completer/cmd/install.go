package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("audit", false, "Conduct security audit")
	installCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	installCmd.Flags().Bool("dry-run", false, "Only report changes")
	installCmd.Flags().Bool("fund", false, "Display funding message")
	installCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	installCmd.Flags().Bool("global-style", false, "Use global layout")
	installCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	installCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	installCmd.Flags().Bool("no-save", false, "Prevents saving to `dependencies`")
	installCmd.Flags().String("omit", "", "Exclude package")
	installCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	installCmd.Flags().BoolP("save", "S", false, "Package will appear in your `dependencies`")
	installCmd.Flags().Bool("save-dev", false, "Package will appear in your `devDependencies`")
	installCmd.Flags().BoolP("save-exact", "E", false, "Save exact package version")
	installCmd.Flags().Bool("save-optional", false, "Package will appear in your `optionalDependencies`")
	installCmd.Flags().Bool("save-peer", false, "Package will appear in your `peerDependencies`")
	installCmd.Flags().Bool("save-prod", false, "Package will appear in your `dependencies`.")
	installCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(installCmd)
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return action.ActionPackages(installCmd)
		}),
	)
}
