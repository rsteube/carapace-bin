package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pub"
	"github.com/spf13/cobra"
)

var pub_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a dependency to pubspec.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_addCmd).Standalone()

	pub_addCmd.Flags().BoolP("dev", "d", false, "Adds package to the development dependencies instead.")
	pub_addCmd.Flags().StringP("directory", "C", "", "Run this in the directory <dir>.")
	pub_addCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change")
	pub_addCmd.Flags().String("git-path", "", "Path of git package in repository")
	pub_addCmd.Flags().String("git-ref", "", "Git branch or commit to be retrieved")
	pub_addCmd.Flags().String("git-url", "", "Git URL of the package")
	pub_addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_addCmd.Flags().String("hosted-url", "", "URL of package host server")
	pub_addCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	pub_addCmd.Flags().Bool("no-precompile", false, "Do not build executables in immediate dependencies.")
	pub_addCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	pub_addCmd.Flags().String("path", "", "Local path")
	pub_addCmd.Flags().Bool("precompile", false, "Build executables in immediate dependencies.")
	pub_addCmd.Flags().String("sdk", "", "SDK source for package")
	pubCmd.AddCommand(pub_addCmd)

	carapace.Gen(pub_addCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
		// TODO "git-path":
		"git-ref": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := pub_addCmd.Flag("git-url"); flag.Changed {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: flag.Value.String(), Branches: true, Tags: true})
			}
			return carapace.ActionMessage("git-url not set")
		}),
		"path": carapace.ActionDirectories(),
		// TODO "sdk":
	})

	carapace.Gen(pub_addCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return pub.ActionPackageSearch().Invoke(c).Suffix(":").ToA()
			case 1:
				return pub.ActionPackageVersions(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
