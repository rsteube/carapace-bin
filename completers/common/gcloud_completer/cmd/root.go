package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "gcloud",
	Short:              "manage Google Cloud Platform resources and developer workflow",
	Long:               "https://cloud.google.com/sdk/gcloud/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO patch user@instance and --flag=optarg as in gcloud completion script

			if c.CallbackValue == "-" {
				return carapace.ActionValues("--").NoSpace() // seems shorthand flags aren't completed anyway so expand to longhand first
			}
			c.Setenv("CLOUDSDK_COMPONENT_MANAGER_DISABLE_UPDATE_CHECK", "1")
			return bridge.ActionArgcomplete("gcloud").Invoke(c).ToA()
		}),
	)
}
