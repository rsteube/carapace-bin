package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cargo-upgrade",
	Short: "Update dependencies as recorded in the local lock file",
	Long:  "https://github.com/killercup/cargo-edit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "[deprecated in favor of `--workspace`]")
	rootCmd.Flags().Bool("allow-prerelease", false, "Include prerelease versions when fetching from crates.io (e.g. 0.6.0-alpha')")
	rootCmd.Flags().Bool("dry-run", false, "Print changes to be made without making them")
	rootCmd.Flags().String("exclude", "", "Crates to exclude and not upgrade")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().String("manifest-path", "", "Path to the manifest to upgrade")
	rootCmd.Flags().Bool("offline", false, "Run without accessing the network")
	rootCmd.Flags().StringP("package", "p", "", "Package id of the crate to add this dependency to")
	rootCmd.Flags().Bool("skip-compatible", false, "Only update a dependency if the new version is semver incompatible")
	rootCmd.Flags().Bool("to-lockfile", false, "Upgrade all packages to the version in the lockfile")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.Flags().Bool("workspace", false, "Upgrade all packages in the workspace")

	// TODO exclude / package
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(".toml"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return cargo.ActionDependencies(cargo.DependencyOpts{
				Path:           rootCmd.Flag("manifest-path").Value.String(),
				IncludeVersion: false,
			}).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
