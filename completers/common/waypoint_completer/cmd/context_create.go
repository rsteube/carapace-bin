package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var context_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_createCmd).Standalone()

	context_createCmd.Flags().String("server-addr", "", "Address for the server.")
	context_createCmd.Flags().String("server-auth-token", "", "Authentication token to use to connect to the server.")
	context_createCmd.Flags().Bool("server-require-auth", false, "If true, will send authentication details.")
	context_createCmd.Flags().Bool("server-tls", false, "If true, will connect to the server over TLS.")
	context_createCmd.Flags().Bool("server-tls-skip-verify", false, "If true, will not validate TLS cert presented by the server.")
	context_createCmd.Flags().Bool("set-default", false, "Set this context as the new default for the CLI. The default is false.")

	addGlobalOptions(context_createCmd)

	contextCmd.AddCommand(context_createCmd)
}
