package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Flutter settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("analytics", false, "Enable reporting anonymously tool usage statistics and crash reports.")
	configCmd.Flags().String("android-sdk", "", "The Android SDK directory.")
	configCmd.Flags().String("android-studio-dir", "", "The Android Studio install directory.")
	configCmd.Flags().String("build-dir", "", "The relative path to override a projects build directory.")
	configCmd.Flags().Bool("clear-features", false, "Remove all configured features and restore them to the default values.")
	configCmd.Flags().Bool("clear-ios-signing-cert", false, "Clear the saved development certificate choice used to sign apps for iOS device deployment.")
	configCmd.Flags().Bool("enable-android", false, "Enable Flutter for Android.")
	configCmd.Flags().Bool("enable-custom-devices", false, "Enable Early support for custom device types.")
	configCmd.Flags().Bool("enable-fuchsia", false, "Enable Flutter for Fuchsia.")
	configCmd.Flags().Bool("enable-ios", false, "Enable Flutter for iOS.")
	configCmd.Flags().Bool("enable-linux-desktop", false, "Enable beta-quality support for desktop on Linux.")
	configCmd.Flags().Bool("enable-macos-desktop", false, "Enable beta-quality support for desktop on macOS.")
	configCmd.Flags().Bool("enable-web", false, "Enable Flutter for web.")
	configCmd.Flags().Bool("enable-windows-desktop", false, "Enable beta-quality support for desktop on Windows.")
	configCmd.Flags().Bool("enable-windows-uwp-desktop", false, "Enable  Flutter for Windows UWP.")
	configCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	configCmd.Flags().Bool("no-analytics", false, "Disable reporting anonymously tool usage statistics and crash reports.")
	configCmd.Flags().Bool("no-enable-android", false, "Disable Flutter for Android.")
	configCmd.Flags().Bool("no-enable-custom-devices", false, "Disable Early support for custom device types.")
	configCmd.Flags().Bool("no-enable-fuchsia", false, "Disable Flutter for Fuchsia.")
	configCmd.Flags().Bool("no-enable-ios", false, "Disable Flutter for iOS.")
	configCmd.Flags().Bool("no-enable-linux-desktop", false, "Disable beta-quality support for desktop on Linux.")
	configCmd.Flags().Bool("no-enable-macos-desktop", false, "Disable beta-quality support for desktop on macOS.")
	configCmd.Flags().Bool("no-enable-web", false, "Disable Flutter for web.")
	configCmd.Flags().Bool("no-enable-windows-desktop", false, "Disable beta-quality support for desktop on Windows.")
	configCmd.Flags().Bool("no-enable-windows-uwp-desktop", false, "Disable Flutter for Windows UWP.")
	configCmd.Flags().Bool("no-single-widget-reload-optimization", false, "Disable Hot reload optimization for changes to class body of a single widget.")
	configCmd.Flags().Bool("single-widget-reload-optimization", false, "Enable Hot reload optimization for changes to class body of a single widget.")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"android-sdk":        carapace.ActionDirectories(),
		"android-studio-dir": carapace.ActionDirectories(),
		"build-dir":          carapace.ActionDirectories(),
	})
}
