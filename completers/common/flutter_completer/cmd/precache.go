package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var precacheCmd = &cobra.Command{
	Use:   "precache",
	Short: "Populate the Flutter tool's cache of binary artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(precacheCmd).Standalone()

	precacheCmd.Flags().BoolP("all-platforms", "a", false, "Precache artifacts for all host platforms.")
	precacheCmd.Flags().BoolP("force", "f", false, "Force re-downloading of artifacts.")
	precacheCmd.Flags().Bool("fuchsia", false, "Precache artifacts for Fuchsia development.")
	precacheCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	precacheCmd.Flags().Bool("ios", false, "Precache artifacts for iOS development.")
	precacheCmd.Flags().Bool("linux", false, "Precache artifacts for Linux desktop development.")
	precacheCmd.Flags().Bool("macos", false, "Precache artifacts for macOS desktop development.")
	precacheCmd.Flags().Bool("no-fuchsia", false, "Do not precache artifacts for Fuchsia development.")
	precacheCmd.Flags().Bool("no-ios", false, "Do not precache artifacts for iOS development.")
	precacheCmd.Flags().Bool("no-linux", false, "Do not precache artifacts for Linux desktop development.")
	precacheCmd.Flags().Bool("no-macos", false, "Do not precache artifacts for macOS desktop development.")
	precacheCmd.Flags().Bool("no-universal", false, "Do not precache artifacts required for any development platform.")
	precacheCmd.Flags().Bool("no-web", false, "Do not precache artifacts for web development.")
	precacheCmd.Flags().Bool("no-windows", false, "Do not precache artifacts for Windows desktop development.")
	precacheCmd.Flags().Bool("no-winuwp", false, "Do not precache artifacts for Windows UWP desktop development.")
	precacheCmd.Flags().Bool("universal", false, "Precache artifacts required for any development platform.")
	precacheCmd.Flags().Bool("web", false, "Precache artifacts for web development.")
	precacheCmd.Flags().Bool("windows", false, "Precache artifacts for Windows desktop development.")
	precacheCmd.Flags().Bool("winuwp", false, "Precache artifacts for Windows UWP desktop development.")
	rootCmd.AddCommand(precacheCmd)
}
