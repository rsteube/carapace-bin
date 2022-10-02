package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pub_global_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an executable from a globally activated package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_global_runCmd).Standalone()

	pub_global_runCmd.Flags().Bool("enable-asserts", false, "Enable assert statements.")
	pub_global_runCmd.Flags().String("enable-experiment", "", "Runs the executable in a VM with the given experiments enabled.")
	pub_global_runCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_global_runCmd.Flags().Bool("no-enable-asserts", false, "Disable assert statements.")
	pub_global_runCmd.Flags().Bool("no-sound-null-safety", false, "Do not override the default null safety execution mode.")
	pub_global_runCmd.Flags().Bool("sound-null-safety", false, "Override the default null safety execution mode.")
	pub_globalCmd.AddCommand(pub_global_runCmd)

	carapace.Gen(pub_global_runCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionActivePackages().Invoke(c).Suffix(":").ToA()
			case 1:
				return action.ActionActivePackageExecutables(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
