package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "picard",
	Short: "Picard is a cross-platform music tagger written in Python",
	Long:  "https://picard.musicbrainz.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config-file", "c", "", "location of the configuration file")
	rootCmd.Flags().BoolP("debug", "d", false, "enable debug-level logging")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().BoolP("long-version", "V", false, "display long version information and exit")
	rootCmd.Flags().BoolP("no-player", "M", false, "disable built-in media player")
	rootCmd.Flags().BoolP("no-plugins", "P", false, "do not load any plugins")
	rootCmd.Flags().BoolP("no-restore", "N", false, "do not restore positions and/or sizes")
	rootCmd.Flags().BoolP("version", "v", false, "display version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
