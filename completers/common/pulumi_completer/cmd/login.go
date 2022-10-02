package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the Pulumi service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()
	loginCmd.PersistentFlags().StringP("cloud-url", "c", "", "A cloud URL to log in to")
	loginCmd.PersistentFlags().String("default-org", "", "A default org to associate with the login. Please note, currently, only the managed and self-hosted backends support organizations")
	loginCmd.PersistentFlags().BoolP("local", "l", false, "Use Pulumi in local-only mode")
	rootCmd.AddCommand(loginCmd)
}
