package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmus",
	Short: "Curses based music player",
	Long:  "https://cmus.github.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().String("listen", "", "listen on ADDR instead of $CMUS_SOCKET or $XDG_RUNTIME_DIR/cmus-socket")
	rootCmd.Flags().Bool("plugins", false, "list available plugins and exit")
	rootCmd.Flags().Bool("show-cursor", false, "always visible cursor")
	rootCmd.Flags().Bool("version", false, "show version")
}
