package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var analyticsCmd = &cobra.Command{
	Use:   "analytics",
	Short: "Configures the gathering of Angular CLI usage metrics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyticsCmd).Standalone()

	rootCmd.AddCommand(analyticsCmd)

	carapace.Gen(analyticsCmd).PositionalCompletion(
		carapace.ActionValues(
			"cli.analyticsSharing",
			"cli.analyticsSharing.tracking",
			"cli.analyticsSharing.uuid",
		),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[0] {
			case "cli.analyticsSharing":
				return carapace.ActionValuesDescribed("undefined", "turn off analytics")
			case "cli.analyticsSharing.uuid":
				return carapace.ActionValuesDescribed("", "generate random user ID")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
