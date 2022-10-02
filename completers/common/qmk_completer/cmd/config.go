package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read and write configuration settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().BoolP("all", "a", false, "Show all configuration options.")
	configCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	configCmd.Flags().Bool("read-only", false, "Operate in read-only mode.") // skip -ro
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionConfigs().Invoke(c).Filter(c.Args).ToMultiPartsA(".")
			case 1:
				return action.ActionConfigValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
