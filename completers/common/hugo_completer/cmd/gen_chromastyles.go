package cmd

import (
	"github.com/spf13/cobra"
)

var gen_chromastylesCmd = &cobra.Command{
	Use:   "chromastyles",
	Short: "Generate CSS stylesheet for the Chroma code highlighter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	gen_chromastylesCmd.PersistentFlags().String("highlightStyle", "bg:#ffffcc", "style used for highlighting lines (see https://github.com/alecthomas/chroma)")
	gen_chromastylesCmd.PersistentFlags().String("linesStyle", "", "style used for line numbers (see https://github.com/alecthomas/chroma)")
	gen_chromastylesCmd.PersistentFlags().String("style", "friendly", "highlighter style (see https://help.farbox.com/pygments.html)")
	genCmd.AddCommand(gen_chromastylesCmd)

	// TODO styles
}
