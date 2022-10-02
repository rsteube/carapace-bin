package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add dependency to a Cargo.toml manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().Bool("allow-prerelease", false, "Include prerelease versions when fetching from crates.io (e.g. '0.6.0-alpha')")
	addCmd.Flags().String("branch", "", "Specify a git branch to download the crate from")
	addCmd.Flags().BoolP("build", "B", false, "Add crate as build dependency")
	addCmd.Flags().BoolP("dev", "D", false, "Add crate as development dependency")
	addCmd.Flags().String("features", "", "Space-separated list of features to add. For an alternative approach to enabling features, consider installing the `cargo-feature` utility")
	addCmd.Flags().String("git", "", "Specify a git repository to download the crate from")
	addCmd.Flags().BoolP("help", "h", false, "Prints help information")
	addCmd.Flags().String("manifest-path", "", "Path to the manifest to add a dependency to")
	addCmd.Flags().Bool("no-default-features", false, "Set `default-features = false` for the added dependency")
	addCmd.Flags().Bool("offline", false, "Run without accessing the network")
	addCmd.Flags().Bool("optional", false, "Add as an optional dependency (for use in features)")
	addCmd.Flags().StringP("package", "p", "", "Package id of the crate to add this dependency to")
	addCmd.Flags().String("path", "", "Specify the path the crate should be loaded from")
	addCmd.Flags().BoolP("quiet", "q", false, "Do not print any output in case of success")
	addCmd.Flags().String("registry", "", "Registry to use")
	addCmd.Flags().StringP("rename", "r", "", "Rename a dependency in Cargo.toml")
	addCmd.Flags().BoolP("sort", "s", false, "Sort dependencies even if currently unsorted")
	addCmd.Flags().String("target", "", "Add as dependency to the given target platform")
	addCmd.Flags().String("upgrade", "", "Choose method of semantic version upgrade")
	addCmd.Flags().String("vers", "", "Specify the version to grab from the registry(crates.io). You can also specify version as part of name, e.g `cargo add bitflags@0.3.2`")
	addCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(addCmd)

	// TODO flag completion
	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := addCmd.Flag("git"); f.Changed && strings.HasPrefix(f.Value.String(), "https://github.com/") {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: f.Value.String(), Branches: true, Tags: true})
			}
			return carapace.ActionValues()
		}),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) == 1 && !util.HasPathPrefix(c.Args[0]) {
				// TODO limit to specific version
				return action.ActionGithubPackageFeatures(strings.Split(c.Args[0], "@")[0]).Invoke(c).Filter(c.Parts).ToA()
			}
			return carapace.ActionValues()
		}),
		"git": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRepositorySearch()
		}),
		"path":    carapace.ActionDirectories(),
		"upgrade": carapace.ActionValuesDescribed("none", "patch", "minor", "all", "default"),
		"vers": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 1 && !util.HasPathPrefix(c.Args[0]) {
				return action.ActionGithubPackageVersions(strings.Split(c.Args[0], "@")[0])
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionDirectories()
			}

			return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionGithubPackageSearch()
				case 1:
					return action.ActionGithubPackageVersions(c.Parts[0])
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
