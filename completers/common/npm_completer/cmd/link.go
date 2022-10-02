package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Symlink a package folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().Bool("audit", false, "Conduct security audit")
	linkCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	linkCmd.Flags().Bool("dry-run", false, "Only report changes")
	linkCmd.Flags().Bool("fund", false, "Display funding message")
	linkCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	linkCmd.Flags().Bool("global-style", false, "Use global layout")
	linkCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	linkCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	linkCmd.Flags().Bool("no-save", false, "Prevents saving to `dependencies`")
	linkCmd.Flags().StringArray("omit", []string{""}, "Exclude package")
	linkCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	linkCmd.Flags().BoolP("save", "S", false, "Package will appear in your `dependencies`")
	linkCmd.Flags().Bool("save-dev", false, "Package will appear in your `devDependencies`")
	linkCmd.Flags().BoolP("save-exact", "E", false, "Save exact package version")
	linkCmd.Flags().Bool("save-optional", false, "Package will appear in your `optionalDependencies`")
	linkCmd.Flags().Bool("save-peer", false, "Package will appear in your `peerDependencies`")
	linkCmd.Flags().Bool("save-prod", false, "Package will appear in your `dependencies`.")
	linkCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(linkCmd)
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).FlagCompletion(carapace.ActionMap{
		"omit": carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(linkCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return action.ActionPackages(linkCmd)
		}),
	)
}
