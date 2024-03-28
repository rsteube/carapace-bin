package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k3d",
	Short: "https://k3d.io/ -> Run k3s in Docker!",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("timestamps", false, "Enable Log timestamps")
	rootCmd.PersistentFlags().Bool("trace", false, "Enable super verbose output (trace logging)")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose output (debug logging)")
	rootCmd.Flags().Bool("version", false, "Show k3d and default k3s version")
}
