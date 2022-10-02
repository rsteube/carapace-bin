package cmd

import (
	"github.com/spf13/cobra"
)

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "List OpenFaaS namespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	namespacesCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	namespacesCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	namespacesCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	rootCmd.AddCommand(namespacesCmd)
}
