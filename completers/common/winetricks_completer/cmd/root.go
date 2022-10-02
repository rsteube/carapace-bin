package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/winetricks_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "winetricks",
	Short: "manage virtual Windows environments using Wine",
	Long:  "https://wiki.winehq.org/Winetricks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("country", "", "Set country code to CC and don't detect your IP address")
	rootCmd.Flags().BoolP("ddrescue", "r", false, "Retry hard when caching scratched discs")
	rootCmd.Flags().BoolP("force", "f", false, "Don't check whether packages were already installed")
	rootCmd.Flags().Bool("gui", false, "Show gui diagnostics even when driven by commandline")
	rootCmd.Flags().BoolP("help", "h", false, "Display this message and exit")
	rootCmd.Flags().Bool("isolate", false, "Install each app or game in its own bottle (WINEPREFIX)")
	rootCmd.Flags().BoolP("keep_isos", "k", false, "Cache isos (allows later installation without disc)")
	rootCmd.Flags().Bool("no-clean", false, "Don't delete temp directories (useful during debugging)")
	rootCmd.Flags().Bool("self-update", false, "Update this application to the last version")
	rootCmd.Flags().BoolP("torify", "t", false, "Run downloads under torify, if available")
	rootCmd.Flags().BoolP("unattended", "q", false, "Don't ask any questions, just install automatically")
	rootCmd.Flags().BoolP("verbose", "v", false, "Echo all commands as they are executed")
	rootCmd.Flags().Bool("verify", false, "Run (automated) GUI tests for verbs, if available")
	rootCmd.Flags().BoolP("version", "V", false, "Display version and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionVerbs(),
				carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValuesDescribed(
							"arch", "create wineprefix with 32 or 64 bit",
							"wineprefix", "select WINEPREFIX",
						).Invoke(c).Suffix("=").ToA()
					case 1:
						switch c.Parts[0] {
						case "arch":
							return carapace.ActionValues("32", "64")
						case "wineprefix":
							return action.ActionPrefixes()

						}
					}
					return carapace.ActionValues()
				}),
			).Invoke(c).Merge().ToA()
		}),
	)
}
