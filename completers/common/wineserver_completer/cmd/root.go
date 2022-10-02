package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wineserver",
	Short: "the Wine server",
	Long:  "https://wiki.winehq.org/Wineserver",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("debug", "d", "", "set debug level to n or +1 if n not specified")
	rootCmd.Flags().BoolP("foreground", "f", false, "remain in the foreground for debugging")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message")
	rootCmd.Flags().StringP("kill", "k", "", "kill the current wineserver, optionally with signal n")
	rootCmd.Flags().StringP("persistent", "p", "", "make server persistent, optionally for n seconds")
	rootCmd.Flags().BoolP("version", "v", false, "display version information and exit")
	rootCmd.Flags().BoolP("wait", "w", false, "wait until the current wineserver terminates")

	rootCmd.Flag("debug").NoOptDefVal = " "
	rootCmd.Flag("kill").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"debug": carapace.ActionValuesDescribed(
			"0", "no debugging information",
			"1", "normal",
			"2", "extra verbose",
		),
		"kill": ps.ActionKillSignals(),
	})
}
