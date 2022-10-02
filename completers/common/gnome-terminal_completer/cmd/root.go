package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gnome-terminal",
	Short: "A terminal emulator for GNOME",
	Long:  "https://help.gnome.org/users/gnome-terminal/stable/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().Bool("help-terminal", false, "Show terminal options")
	rootCmd.Flags().Bool("help-terminal-options", false, "Show per-terminal options")
	rootCmd.Flags().Bool("help-window-options", false, "Show per-window options")
	rootCmd.Flags().String("load-config", "", "Load a terminal configuration file")
	rootCmd.Flags().Bool("no-environment", false, "Do not pass the environment")
	rootCmd.Flags().Bool("preferences", false, "Show preferences window")
	rootCmd.Flags().BoolP("print-environment", "p", false, "Print environment variables to interact with the terminal")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress output")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase diagnostic verbosity")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"display":     os.ActionDisplays(),
		"load-config": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
