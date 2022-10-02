package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var toolchain_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install or update a given toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_installCmd).Standalone()

	toolchain_installCmd.Flags().Bool("allow-downgrade", false, "Allow rustup to downgrade the toolchain to satisfy your component choice")
	toolchain_installCmd.Flags().StringP("component", "c", "", "Add specific components on installation")
	toolchain_installCmd.Flags().Bool("force", false, "Force an update, even if some components are missing")
	toolchain_installCmd.Flags().BoolP("help", "h", false, "Prints help information")
	toolchain_installCmd.Flags().Bool("no-self-update", false, "Don't perform self update when running the`rustup toolchain install` command")
	toolchain_installCmd.Flags().String("profile", "", "[possible values: minimal, default, complete]")
	toolchain_installCmd.Flags().StringP("target", "t", "", "Add specific targets on installation")
	toolchainCmd.AddCommand(toolchain_installCmd)

	carapace.Gen(toolchain_installCmd).FlagCompletion(carapace.ActionMap{
		"component": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAvailableComponents().Invoke(c).Filter(c.Parts).ToA()
		}),
		"profile": carapace.ActionValues("minimal", "default", "complete"),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionTargets(false).Invoke(c).ToMultiPartsA("-")
		}),
	})

	carapace.Gen(toolchain_installCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionDefaultChannels().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
