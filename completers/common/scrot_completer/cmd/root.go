package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scrot",
	Short: "command line screen capture utility",
	Long:  "https://en.wikipedia.org/wiki/Scrot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("autoselect", "a", false, "non-interactively choose a rectangle of x,y,w,h")
	rootCmd.Flags().BoolP("border", "b", false, "When selecting a window, grab wm border too")
	rootCmd.Flags().StringP("class", "C", "", "Window class name. Associative with options: -k")
	rootCmd.Flags().BoolP("count", "c", false, "show a countdown before taking the shot")
	rootCmd.Flags().StringP("delay", "d", "", "wait NUM seconds before taking a shot")
	rootCmd.Flags().BoolP("display", "D", false, "Set DISPLAY target other than current")
	rootCmd.Flags().StringP("exec", "e", "", "run APP on the resulting screenshot")
	rootCmd.Flags().BoolP("focused", "u", false, "use the currently focused window")
	rootCmd.Flags().BoolP("freeze", "f", false, "Freeze the screen when the selection is used: --select")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("line", "l", false, "Indicates the style of the line when the selection is used: --select")
	rootCmd.Flags().BoolP("multidisp", "m", false, "For multiple heads, grab shot from each and join them together.")
	rootCmd.Flags().BoolP("note", "n", false, "Draw a text note.")
	rootCmd.Flags().BoolP("overwrite", "o", false, "By default scrot does not overwrite the files, use this option to allow it.")
	rootCmd.Flags().BoolP("pointer", "p", false, "Capture the mouse pointer.")
	rootCmd.Flags().StringP("quality", "q", "", "Image quality (1-100) high value means high size, low compression")
	rootCmd.Flags().StringP("script", "S", "", "Imlib2 script commands")
	rootCmd.Flags().BoolP("select", "s", false, "interactively choose a window or rectangle with the mouse")
	rootCmd.Flags().BoolP("silent", "z", false, "Prevent beeping")
	rootCmd.Flags().BoolP("stack", "k", false, "Capture stack/overlapped windows and join them together.")
	rootCmd.Flags().StringP("thumb", "t", "", "generate thumbnail too")
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exec": carapace.ActionFiles(),
	})
}
