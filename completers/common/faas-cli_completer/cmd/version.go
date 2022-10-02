package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the clients version information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	versionCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	versionCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	versionCmd.Flags().Bool("short-version", false, "Just print Git SHA")
	versionCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	versionCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	versionCmd.Flags().Bool("warn-update", true, "Check for new version and warn about updating")
	rootCmd.AddCommand(versionCmd)
}
