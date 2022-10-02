package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conky",
	Short: "A system  monitor for X originally based on the torsmo code",
	Long:  "https://github.com/brndnmtthws/conky",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("alignment", "a", "", "text alignment on screen")
	rootCmd.Flags().StringP("config", "c", "", "config file to load")
	rootCmd.Flags().BoolP("daemonize", "d", false, "daemonize, fork to background")
	rootCmd.Flags().BoolP("debug", "D", false, "increase debugging output, ie. -DD for more debugging")
	rootCmd.Flags().StringP("display", "X", "", "X11 display to use")
	rootCmd.Flags().BoolP("double-buffer", "b", false, "double buffer (prevents flickering)")
	rootCmd.Flags().StringP("font", "f", "", "font to use")
	rootCmd.Flags().BoolP("help", "h", false, "help")
	rootCmd.Flags().StringS("i", "i", "", "number of times to update conky (and quit)")
	rootCmd.Flags().StringP("interval", "u", "", "update interval")
	rootCmd.Flags().BoolP("own-window", "o", false, "create own window to draw")
	rootCmd.Flags().StringP("pause", "p", "", "pause for SECS seconds at startup before doing anything")
	rootCmd.Flags().BoolP("print-config", "C", false, "print the builtin default config to stdout")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.Flags().StringP("text", "t", "", "text to render, remember single quotes, like -t '$uptime'")
	rootCmd.Flags().BoolP("version", "v", false, "version")
	rootCmd.Flags().StringP("window-id", "w", "", "window id to draw")
	rootCmd.Flags().StringS("x", "x", "", "x position")
	rootCmd.Flags().StringP("xinerama-head", "m", "", "Xinerama monitor index (0=first)")
	rootCmd.Flags().StringS("y", "y", "", "y position")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"alignment": carapace.ActionMultiParts("_", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("top", "bottom", "middle").Invoke(c).Suffix("_").ToA()
			case 1:
				return carapace.ActionValues("left", "right", "middle")
			default:
				return carapace.ActionValues()
			}
		}),
		"config":  carapace.ActionFiles(),
		"display": os.ActionDisplays(),
	})
}
