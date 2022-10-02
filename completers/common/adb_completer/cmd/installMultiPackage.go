package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var installMultiPackageCmd = &cobra.Command{
	Use:   "install-multi-package",
	Short: "push one or more packages to the device and install them atomically",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installMultiPackageCmd).Standalone()

	installMultiPackageCmd.Flags().String("abi", "", "override platform's default ABI")
	installMultiPackageCmd.Flags().BoolS("d", "d", false, "allow version code downgrade (debuggable packages only)")
	installMultiPackageCmd.Flags().Bool("date-check-agent", false, "update deployment agent when local version is newer and using fast deploy")
	installMultiPackageCmd.Flags().Bool("fastdeploy", false, "use fast deploy")
	installMultiPackageCmd.Flags().Bool("force-agent", false, "force update of deployment agent when using fast deploy")
	installMultiPackageCmd.Flags().BoolS("g", "g", false, "grant all runtime permissions")
	installMultiPackageCmd.Flags().Bool("instant", false, "cause the app to be installed as an ephemeral install app")
	installMultiPackageCmd.Flags().Bool("local-agent", false, "locate agent files from local source build (instead of SDK location)")
	installMultiPackageCmd.Flags().Bool("no-fastdeploy", false, "prevent use of fast deploy")
	installMultiPackageCmd.Flags().Bool("no-streaming", false, "always push APK to device and invoke Package Manager as separate steps")
	installMultiPackageCmd.Flags().BoolS("p", "p", false, "partial application install (install-multiple only)")
	installMultiPackageCmd.Flags().BoolS("r", "r", false, "replace existing application")
	installMultiPackageCmd.Flags().Bool("streaming", false, "force streaming APK directly into Package Manager")
	installMultiPackageCmd.Flags().BoolS("t", "t", false, "allow test packages")
	installMultiPackageCmd.Flags().Bool("version-check-agent", false, "update deployment agent when local version has different version code and using fast deploy")
	rootCmd.AddCommand(installMultiPackageCmd)

	carapace.Gen(installMultiPackageCmd).FlagCompletion(carapace.ActionMap{
		"abi": carapace.ActionValues("armeabi-v7a", "arm64-v8a", "x86", "x86_64"),
	})

	carapace.Gen(installMultiPackageCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".apk"),
	)
}
