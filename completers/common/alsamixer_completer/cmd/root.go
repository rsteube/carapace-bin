package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "alsamixer",
	Short: "soundcard mixer for ALSA soundcard driver, with ncurses interface",
	Long:  "https://en.wikipedia.org/wiki/Alsamixer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("abstraction", "a", "", "mixer abstraction level: none/basic")
	rootCmd.Flags().StringP("card", "c", "", "sound card number or id")
	rootCmd.Flags().StringP("config", "f", "", "configuration file")
	rootCmd.Flags().StringP("device", "D", "", "mixer device name")
	rootCmd.Flags().BoolP("help", "h", false, "this help")
	rootCmd.Flags().BoolP("mouse", "m", false, "enable mouse")
	rootCmd.Flags().BoolP("no-color", "g", false, "toggle using of colors")
	rootCmd.Flags().BoolP("no-config", "F", false, "do not load configuration file")
	rootCmd.Flags().BoolP("no-mouse", "M", false, "disable mouse")
	rootCmd.Flags().StringP("view", "V", "", "starting view mode: playback/capture/all")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"abstraction": carapace.ActionValues("none", "basic"),
		"card":        os.ActionSoundCards(),
		"config":      carapace.ActionFiles(),
		"view":        carapace.ActionValues("playback", "capture", "all"),
	})
}
