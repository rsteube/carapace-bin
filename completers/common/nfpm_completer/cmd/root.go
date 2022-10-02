package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nfpm",
	Short: "Packages apps on RPM, Deb and APK formats based on a YAML configuration file",
	Long:  "https://github.com/goreleaser/nfpm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "help for nfpm")
	rootCmd.Flags().BoolP("version", "v", false, "version for nfpm")
}
