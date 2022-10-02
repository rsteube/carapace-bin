package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/dfc_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dfc",
	Short: "report file system space usage information with style",
	Long:  "https://github.com/rolinh/dfc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("M", "M", false, "do not print \"mounted on\"")
	rootCmd.Flags().BoolS("T", "T", false, "show filesystem type")
	rootCmd.Flags().BoolS("W", "W", false, "wide filename (un truncate)")
	rootCmd.Flags().BoolS("a", "a", false, "print all mounted filesystem")
	rootCmd.Flags().BoolS("b", "b", false, "do not show the graph bar")
	rootCmd.Flags().StringS("c", "c", "", "choose color mode.")
	rootCmd.Flags().BoolS("d", "d", false, "show used size")
	rootCmd.Flags().StringS("e", "e", "", "export to specified format.")
	rootCmd.Flags().BoolS("f", "f", false, "disable auto-adjust mode (force display)")
	rootCmd.Flags().BoolS("h", "h", false, "print this message")
	rootCmd.Flags().BoolS("i", "i", false, "info about inodes")
	rootCmd.Flags().BoolS("l", "l", false, "only show information about locally mounted file systems")
	rootCmd.Flags().BoolS("m", "m", false, "use metric (SI unit)")
	rootCmd.Flags().BoolS("n", "n", false, "do not print header")
	rootCmd.Flags().BoolS("o", "o", false, "show mount flags")
	rootCmd.Flags().StringS("p", "p", "", "filter by file system name.")
	rootCmd.Flags().StringS("q", "q", "", "sort the output.")
	rootCmd.Flags().BoolS("s", "s", false, "sum the total usage")
	rootCmd.Flags().StringS("t", "t", "", "filter by file system type.")
	rootCmd.Flags().StringS("u", "u", "", "choose the unit in which to show the values.")
	rootCmd.Flags().BoolS("v", "v", false, "print program version")
	rootCmd.Flags().BoolS("w", "w", false, "use a wider bar")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionValuesDescribed(
			"always", "Color will always be used, no matter what \"stdout\" is.",
			"auto", "Color is used only if \"stdout\" is a terminal.",
			"never", "Color will never be used.",
		),
		"e": carapace.ActionValuesDescribed(
			"csv", "Output as \"comma separated value\" file type.",
			"html", "Output is HTML formated.",
			"json", "Output is JSON formated.",
			"tex", "Output is TeX formated.",
			"text", "Text output (default).",
		),
		"p": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionFSNames().Invoke(c).Filter(c.Parts).ToA()
		}),
		"q": carapace.ActionValues("name", "type", "mount"),
		"t": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return fs.ActionFilesystemTypes().Invoke(c).Filter(c.Parts).ToA()
		}),
		"u": carapace.ActionValuesDescribed(
			"h", "Human readable",
			"b", "Show bytes.",
			"k", "Show size using Kio.",
			"m", "Show size using Mio.",
			"g", "Show size using Gio.",
			"t", "Show size using Tio.",
			"p", "Show size using Pio.",
			"e", "Show size using Eio.",
			"z", "Show size using Zio.",
			"y", "Show size using Yio.",
		),
	})

}
