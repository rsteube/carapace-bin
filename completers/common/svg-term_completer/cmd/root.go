package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "svg-term",
	Short: "Share terminal sessions as razor-sharp animated SVG everywhere",
	Long:  "https://github.com/marionebl/svg-term-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("at", "", "timestamp of frame to render in ms [number]")
	rootCmd.Flags().String("cast", "", "asciinema cast id to download [string], required if no stdin provided [string]")
	rootCmd.Flags().String("command", "", "command to record [string]")
	rootCmd.Flags().String("from", "", "lower range of timeline to render in ms [number]")
	rootCmd.Flags().String("height", "", "height in lines [number]")
	rootCmd.Flags().Bool("help", false, "print this help [boolean]")
	rootCmd.Flags().String("in", "", "json file to use as input [string]")
	rootCmd.Flags().Bool("no-cursor", false, "disable cursor rendering [boolean]")
	rootCmd.Flags().Bool("no-optimize", false, "disable svgo optimization [boolean]")
	rootCmd.Flags().String("out", "", "output file, emits to stdout if omitted, [string]")
	rootCmd.Flags().String("padding", "", "distance between text and image bounds, [number]")
	rootCmd.Flags().String("padding-x", "", "distance between text and image bounds on x axis [number]")
	rootCmd.Flags().String("padding-y", "", "distance between text and image bounds on y axis [number]")
	rootCmd.Flags().String("profile", "", "terminal profile file to use, requires --term [string]")
	rootCmd.Flags().String("term", "", "terminal profile format, requires --profile [string]")
	rootCmd.Flags().String("to", "", "upper range of timeline to render in ms [number]")
	rootCmd.Flags().String("width", "", "width in columns [number]")
	rootCmd.Flags().Bool("window", false, "render with window decorations [boolean]")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(".json"),
		"out":     carapace.ActionFiles(),
		"profile": carapace.ActionFiles(),
		"term":    carapace.ActionValues("iterm2", "xrdb", "xresources", "terminator", "konsole", "terminal", "remmina", "termite", "tilda", "xcfe"),
	})
}
