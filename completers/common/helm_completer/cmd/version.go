package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the client version information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	versionCmd.Flags().BoolP("client", "c", true, "display client version information")
	versionCmd.Flags().Bool("short", false, "print the version number")
	versionCmd.Flags().String("template", "", "template for version string format")
	rootCmd.AddCommand(versionCmd)
}
