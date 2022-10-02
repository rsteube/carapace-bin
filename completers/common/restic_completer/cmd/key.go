package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Manage keys (passwords)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyCmd).Standalone()
	keyCmd.Flags().String("host", "", "the hostname for new keys")
	keyCmd.Flags().String("new-password-file", "", "`file` from which to read the new password")
	keyCmd.Flags().String("user", "", "the username for new keys")
	rootCmd.AddCommand(keyCmd)

	carapace.Gen(keyCmd).FlagCompletion(carapace.ActionMap{
		"host":              net.ActionHosts(),
		"new-password-file": carapace.ActionFiles(),
		"user":              os.ActionUsers(),
	})
}
