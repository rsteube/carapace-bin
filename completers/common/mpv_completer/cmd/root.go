package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mpv",
	Short: "a media player",
	Long:  "https://mpv.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("fs", false, "fullscreen playback")
	rootCmd.Flags().Bool("list-options", false, "list all mpv options")
	rootCmd.Flags().Bool("no-audio", false, "do not play sound")
	rootCmd.Flags().Bool("no-video", false, "do not play video")
	rootCmd.Flags().String("playlist", "", "specify playlist file")
	rootCmd.Flags().String("start", "", "seek to given (percent, seconds, or hh:mm:ss) position")
	rootCmd.Flags().String("sub-file", "", "specify subtitle file to use")
	// TODO a lot more options listed by `mpv --list-options`

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"playlist": carapace.ActionFiles(),
		"sub-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
