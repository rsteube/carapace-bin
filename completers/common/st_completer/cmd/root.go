package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "st",
	Short: "simple terminal",
	Long:  "https://st.suckless.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("T", "T", "", "defines the window title (default 'st').")
	rootCmd.Flags().BoolS("a", "a", false, "disable alternate screens in terminal")
	rootCmd.Flags().StringS("c", "c", "", "defines the window class (default $TERM).")
	rootCmd.Flags().StringS("e", "e", "", "st executes command instead of the shell.")
	rootCmd.Flags().StringS("f", "f", "", "defines the font to use when st is run.")
	rootCmd.Flags().StringS("g", "g", "", "defines the X11 geometry string.")
	rootCmd.Flags().BoolS("i", "i", false, "will fixate the position given with the -g option.")
	rootCmd.Flags().StringS("l", "l", "", "use a tty line instead of a pseudo terminal.")
	rootCmd.Flags().StringS("n", "n", "", "defines the window instance name (default $TERM).")
	rootCmd.Flags().StringS("o", "o", "", "writes all the I/O to iofile.")
	rootCmd.Flags().StringS("t", "t", "", "defines the window title (default 'st').")
	rootCmd.Flags().BoolS("v", "v", false, "prints version information to stderr, then exits.")
	rootCmd.Flags().StringS("w", "w", "", "embeds st within the window identified by windowid")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"e": carapace.ActionFiles(),
	})
}
