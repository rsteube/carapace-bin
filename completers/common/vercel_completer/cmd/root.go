package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vercel_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vercel",
	Short: "Develop. Preview. Ship.",
	Long:  "https://vercel.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSliceP("build-env", "b", []string{}, "Similar to `--env` but for build time only.")
	rootCmd.Flags().BoolP("confirm", "c", false, "Confirm default options and skip questions")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode")
	rootCmd.Flags().StringSliceP("env", "e", []string{}, "Include an env var during run time (e.g.: `-e KEY=value`). Can appear many times.")
	rootCmd.Flags().BoolP("force", "f", false, "Force a new deployment even if nothing has changed")
	rootCmd.Flags().StringP("global-config", "Q", "", "Path to the global `.vercel` directory")
	rootCmd.Flags().StringP("local-config", "A", "", "Path to the local `vercel.json` file")
	rootCmd.Flags().StringSliceP("meta", "m", []string{}, "Add metadata for the deployment (e.g.: `-m KEY=value`). Can appear many times.")
	rootCmd.Flags().BoolP("no-clipboard", "C", false, "Do not attempt to copy URL to clipboard")
	rootCmd.Flags().BoolP("platform-version", "V", false, "Set the platform version to deploy to")
	rootCmd.Flags().Bool("prod", false, "Create a production deployment")
	rootCmd.Flags().BoolP("public", "p", false, "Deployment is public (`/_src` is exposed)")
	rootCmd.Flags().String("regions", "", "Set default regions to enable the deployment on")
	rootCmd.Flags().StringP("scope", "S", "", "Set a custom scope")
	rootCmd.Flags().StringP("token", "t", "", "Login token")
	rootCmd.Flags().BoolP("version", "v", false, "Output the version number")
	rootCmd.Flags().Bool("with-cache", false, "Retain build cache when using \"--force\"")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Output usage information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"global-config": carapace.ActionDirectories(),
		"local-config":  carapace.ActionFiles(),
		"regions": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionRegions().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionDirectories()
			}
			return carapace.ActionValues()
		}),
	)
}
