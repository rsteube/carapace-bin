package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/sdkmanager_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sdkmanager",
	Short: "Android SDK manager",
	Long:  "https://developer.android.com/studio/command-line/sdkmanager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("channel", "", "Include packages in channels up to <channelId>")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("include_obsolete", false, "show obsolete packages in the package listing")
	rootCmd.Flags().Bool("install", false, "installs or updates packages")
	rootCmd.Flags().Bool("licenses", false, "show and offer the option to accept licenses for all available packages")
	rootCmd.Flags().Bool("list", false, "all installed and available packages are printed out")
	rootCmd.Flags().Bool("list_installed", false, "all installed packages are printed out")
	rootCmd.Flags().Bool("no_https", false, "Force all connections to use http rather than https")
	rootCmd.Flags().String("package_file", "", "package file to use ")
	rootCmd.Flags().String("proxy", "", "Connect via a proxy of the given type")
	rootCmd.Flags().String("proxy_host", "", "IP or DNS address of the proxy to use")
	rootCmd.Flags().String("proxy_port", "", "Proxy port to connect to")
	rootCmd.Flags().String("sdk_root", "", "Use the specified SDK root instead of the SDK containing this tool")
	rootCmd.Flags().Bool("uninstall", false, "uninstall the listed packages")
	rootCmd.Flags().Bool("update", false, "all installed packages are updated to the latest version")
	rootCmd.Flags().Bool("verbose", false, "Enable verbose output.")
	rootCmd.Flags().Bool("version", false, "prints the current version of sdkmanager")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"channel": carapace.ActionValuesDescribed(
			"0", "Stable",
			"1", "Beta",
			"2", "Dev",
			"3", "Canary",
		),
		"package_file": carapace.ActionFiles(),
		"proxy":        carapace.ActionValues("http", "socks"),
		"sdk_root":     carapace.ActionDirectories(),
	})

	rootCmd.Flag("channel").NoOptDefVal = " "
	rootCmd.Flag("package_file").NoOptDefVal = " "
	rootCmd.Flag("sdk_root").NoOptDefVal = " "
	rootCmd.Flag("proxy").NoOptDefVal = " "
	rootCmd.Flag("proxy_host").NoOptDefVal = " "
	rootCmd.Flag("proxy_port").NoOptDefVal = " "

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("install").Changed {
				return action.ActionAvailablePackages(rootCmd).Invoke(c).Filter(c.Args).ToMultiPartsA(";")
			}
			if rootCmd.Flag("uninstall").Changed {
				return action.ActionInstalledPackages(rootCmd).Invoke(c).Filter(c.Args).ToA()
			}
			return carapace.ActionValues()
		}),
	)
}
