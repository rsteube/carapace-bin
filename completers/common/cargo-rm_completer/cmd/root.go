package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cargo-rm",
	Short: "Remove a dependency from a Cargo.toml manifest file",
	Long:  "https://github.com/killercup/cargo-edit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("build", "B", false, "Remove crate as build dependency")
	rootCmd.Flags().BoolP("dev", "D", false, "Remove crate as development dependency")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().String("manifest-path", "", "Path to the manifest to remove a dependency from")
	rootCmd.Flags().StringP("package", "p", "", "Package id of the crate to remove this dependency from")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print any output in case of success")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")

	// TODO package?
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
