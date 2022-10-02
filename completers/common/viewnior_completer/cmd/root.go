package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "viewnior",
	Short: "simple, fast and elegant image viewer",
	Long:  "https://siyanpanayotov.com/project/viewnior",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().Bool("fullscreen", false, "")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().Bool("slideshow", false, "")
	rootCmd.Flags().Bool("version", false, "")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"display": os.ActionDisplays(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
