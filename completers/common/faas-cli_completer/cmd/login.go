package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to OpenFaaS gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	loginCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	loginCmd.Flags().StringP("password", "p", "", "Gateway password")
	loginCmd.Flags().BoolP("password-stdin", "s", false, "Reads the gateway password from stdin")
	loginCmd.Flags().Duration("timeout", 5*time.Second, "Override the timeout for this API call")
	loginCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	loginCmd.Flags().StringP("username", "u", "admin", "Gateway username")
	rootCmd.AddCommand(loginCmd)
}
