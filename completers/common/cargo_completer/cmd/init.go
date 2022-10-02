package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new cargo package in an existing directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	initCmd.Flags().Bool("bin", false, "Use a binary (application) template [default]")
	initCmd.Flags().String("color", "", "Coloring: auto, always, never")
	initCmd.Flags().String("edition", "", "Edition to set for the crate generated [possible values: 2015, 2018]")
	initCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	initCmd.Flags().BoolP("help", "h", false, "Prints help information")
	initCmd.Flags().Bool("lib", false, "Use a library template")
	initCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	initCmd.Flags().String("name", "", "Set the resulting package name, defaults to the directory name")
	initCmd.Flags().Bool("offline", false, "Run without accessing the network")
	initCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	initCmd.Flags().String("registry", "", "Registry to use")
	initCmd.Flags().String("vcs", "", "Initialize a new repository for the given version control system")
	initCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"color":    action.ActionColorModes(),
		"edition":  carapace.ActionValues("2015", "2018"),
		"registry": action.ActionRegistries(),
		"vcs":      carapace.ActionValues("git", "hg", "pijul", "vcs", "none"),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
