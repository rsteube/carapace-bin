package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/mousepad_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mousepad",
	Short: "Mousepad is a simple text editor for the Xfce desktop environment",
	Long:  "http://users.xfce.org/~benny/xfce/apps.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("column", "c", "", "Column number the cursor to position to (COLUMN >= 0 from start, COLUMN < 0 from end)")
	rootCmd.Flags().Bool("disable-server", false, "Do not register with the D-BUS session message bus")
	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().StringP("encoding", "e", "", "Encoding to be used to open files (leave empty to open files in the encoding dialog)")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gapplication", false, "Show GApplication options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().StringP("line", "l", "", "Line number the cursor to position to (LINE > 0 from top, LINE < 0 from bottom)")
	rootCmd.Flags().Bool("list-encodings", false, "Display a list of possible encodings to open files")
	rootCmd.Flags().StringP("opening-mode", "o", "", "File opening mode: \"tab\", \"window\" or \"mixed\" (open tabs in a new window)")
	rootCmd.Flags().Bool("preferences", false, "Open the preferences dialog")
	rootCmd.Flags().BoolP("quit", "q", false, "Quit a running Mousepad primary instance")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"display":      os.ActionDisplays(),
		"encoding":     action.ActionEncoding(),
		"opening-mode": carapace.ActionValues("tab", "window", "mixed"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
