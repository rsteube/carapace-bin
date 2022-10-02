package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove a Rust binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	uninstallCmd.Flags().String("bin", "", "Only uninstall the binary NAME")
	uninstallCmd.Flags().String("color", "", "Coloring: auto, always, never")
	uninstallCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	uninstallCmd.Flags().BoolP("help", "h", false, "Prints help information")
	uninstallCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	uninstallCmd.Flags().Bool("offline", false, "Run without accessing the network")
	uninstallCmd.Flags().StringP("package", "p", "", "Package to uninstall")
	uninstallCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	uninstallCmd.Flags().String("root", "", "Directory to uninstall packages from")
	uninstallCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"bin":     action.ActionTargets(uninstallCmd, action.TargetOpts{Bin: true}),
		"color":   action.ActionColorModes(),
		"package": action.ActionInstalledPackages(uninstallCmd.Flag("root").Value.String()),
		"root":    carapace.ActionDirectories(),
	})

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionInstalledPackages(uninstallCmd.Flag("root").Value.String()).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
