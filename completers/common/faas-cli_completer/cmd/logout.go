package cmd

import (
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from OpenFaaS gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	logoutCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	rootCmd.AddCommand(logoutCmd)
}
