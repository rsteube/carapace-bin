package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gnome-maps",
	Short: "A map application for GNOME",
	Long:  "https://apps.gnome.org/app/org.gnome.Maps/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().Bool("force-online", false, "Ignore network availability")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gapplication", false, "Show GApplication options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().Bool("local", false, "A path to a local tiles directory structure")
	rootCmd.Flags().Bool("local-tile-size", false, "Tile size for local tiles directory")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version of the program")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"display": os.ActionDisplays(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
