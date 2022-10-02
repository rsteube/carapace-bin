package cmd

import (
	"github.com/spf13/cobra"
)

var tunnelCmd = &cobra.Command{
	Use:   "tunnel",
	Short: "Connect to LoadBalancer services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	tunnelCmd.Flags().BoolP("cleanup", "c", true, "call with cleanup=true to remove old tunnels")
	rootCmd.AddCommand(tunnelCmd)
}
