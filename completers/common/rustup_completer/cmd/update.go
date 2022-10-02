package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Rust toolchains and rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("force", false, "Force an update, even if some components are missing")
	updateCmd.Flags().BoolP("help", "h", false, "Prints help information")
	updateCmd.Flags().Bool("no-self-update", false, "Don't perform self update when running the `rustup update` command")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionToolchains().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
