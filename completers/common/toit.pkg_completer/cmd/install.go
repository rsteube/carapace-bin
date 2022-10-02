package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/toit.pkg_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Installs a package in the current project, or downloads all dependencies",
	Aliases: []string{"download", "fetch"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().Bool("local", false, "Treat package argument as local path")
	installCmd.Flags().String("name", "", "The name used for the 'import' clause")
	installCmd.Flags().Bool("recompute", false, "Recompute dependencies")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if installCmd.Flag("local").Changed {
				return carapace.ActionDirectories()
			}

			return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.Batch(
						action.ActionPackages(),
						git.ActionRepositorySearch(), // TODO verify if https prefix is ok
					).ToA()
				case 1:
					return action.ActionPackageVersions(c.Parts[0])
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
