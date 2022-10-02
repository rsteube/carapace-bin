package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nano",
	Short: "Nano's ANOther editor, inspired by Pico",
	Long:  "https://www.nano-editor.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("afterends", "y", false, "Make Ctrl+Right stop at word ends")
	rootCmd.Flags().BoolP("atblanks", "a", false, "When soft-wrapping, do it at whitespace")
	rootCmd.Flags().BoolP("autoindent", "i", false, "Automatically indent new lines")
	rootCmd.Flags().BoolP("backup", "B", false, "Save backups of existing files")
	rootCmd.Flags().StringP("backupdir", "C", "", "Directory for saving unique backup files")
	rootCmd.Flags().BoolP("boldtext", "D", false, "Use bold instead of reverse video text")
	rootCmd.Flags().BoolP("bookstyle", "O", false, "Leading whitespace means new paragraph")
	rootCmd.Flags().BoolP("breaklonglines", "b", false, "Automatically hard-wrap overlong lines")
	rootCmd.Flags().BoolP("constantshow", "c", false, "Constantly show cursor position")
	rootCmd.Flags().BoolP("cutfromcursor", "k", false, "Cut from cursor to end of line")
	rootCmd.Flags().BoolP("emptyline", "e", false, "Keep the line below the title bar empty")
	rootCmd.Flags().StringP("fill", "r", "", "Set width for hard-wrap and justify")
	rootCmd.Flags().StringP("guidestripe", "J", "", "Show a guiding bar at this column")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help text and exit")
	rootCmd.Flags().BoolP("historylog", "H", false, "Log & read search/replace string history")
	rootCmd.Flags().BoolP("ignorercfiles", "I", false, "Don't look at nanorc files")
	rootCmd.Flags().BoolP("indicator", "q", false, "Show a position+portion indicator")
	rootCmd.Flags().BoolP("jumpyscrolling", "j", false, "Scroll per half-screen, not per line")
	rootCmd.Flags().BoolP("linenumbers", "l", false, "Show line numbers in front of the text")
	rootCmd.Flags().BoolP("locking", "G", false, "Use (vim-style) lock files")
	rootCmd.Flags().BoolP("magic", "!", false, "Also try magic to determine syntax")
	rootCmd.Flags().BoolP("minibar", "_", false, "Show a feedback bar at the bottom")
	rootCmd.Flags().BoolP("mouse", "m", false, "Enable the use of the mouse")
	rootCmd.Flags().BoolP("multibuffer", "F", false, "Read a file into a new buffer by default")
	rootCmd.Flags().BoolP("noconvert", "N", false, "Don't convert files from DOS/Mac format")
	rootCmd.Flags().BoolP("nohelp", "x", false, "Don't show the two help lines")
	rootCmd.Flags().BoolP("nonewlines", "L", false, "Don't add an automatic newline")
	rootCmd.Flags().BoolP("noread", "n", false, "Do not read the file (only write it)")
	rootCmd.Flags().BoolP("nowrap", "w", false, "Don't hard-wrap long lines [default]")
	rootCmd.Flags().StringP("operatingdir", "o", "", "Set operating directory")
	rootCmd.Flags().BoolP("positionlog", "P", false, "Log & read location of cursor position")
	rootCmd.Flags().BoolP("preserve", "p", false, "Preserve XON (^Q) and XOFF (^S) keys")
	rootCmd.Flags().BoolP("quickblank", "U", false, "Wipe status bar upon next keystroke")
	rootCmd.Flags().StringP("quotestr", "Q", "", "Regular expression to match quoting")
	rootCmd.Flags().BoolP("rawsequences", "K", false, "Fix numeric keypad key confusion problem")
	rootCmd.Flags().StringP("rcfile", "f", "", "Use only this file for configuring nano")
	rootCmd.Flags().BoolP("rebinddelete", "d", false, "Fix Backspace/Delete confusion problem")
	rootCmd.Flags().BoolP("restricted", "R", false, "Restrict access to the filesystem")
	rootCmd.Flags().BoolP("saveonexit", "t", false, "Save changes on exit, don't prompt")
	rootCmd.Flags().BoolP("showcursor", "g", false, "Show cursor in file browser & help text")
	rootCmd.Flags().BoolP("smarthome", "A", false, "Enable smart home key")
	rootCmd.Flags().BoolP("softwrap", "S", false, "Display overlong lines on multiple rows")
	rootCmd.Flags().StringP("speller", "s", "", "Use this alternative spell checker")
	rootCmd.Flags().BoolP("stateflags", "%", false, "Show some states on the title bar")
	rootCmd.Flags().BoolP("suspendable", "z", false, "Enable suspension")
	rootCmd.Flags().StringP("syntax", "Y", "", "Syntax definition to use for coloring")
	rootCmd.Flags().StringP("tabsize", "T", "", "Make a tab this number of columns wide")
	rootCmd.Flags().BoolP("tabstospaces", "E", false, "Convert typed tabs to spaces")
	rootCmd.Flags().BoolP("trimblanks", "M", false, "Trim tail spaces when hard-wrapping")
	rootCmd.Flags().BoolP("unix", "u", false, "Save a file by default in Unix format")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information and exit")
	rootCmd.Flags().BoolP("view", "v", false, "View mode (read-only)")
	rootCmd.Flags().BoolP("wordbounds", "W", false, "Detect word boundaries more accurately")
	rootCmd.Flags().StringP("wordchars", "X", "", "Which other characters are word parts")
	rootCmd.Flags().BoolP("zap", "Z", false, "Let Bsp and Del erase a marked region")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backupdir":    carapace.ActionDirectories(),
		"operatingdir": carapace.ActionDirectories(),
		"rcfile":       carapace.ActionFiles(),
		"speller":      carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
