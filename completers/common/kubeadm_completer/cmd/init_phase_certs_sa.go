package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_saCmd = &cobra.Command{
	Use:   "sa",
	Short: "Generate a private key for signing service account tokens along with its public key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_saCmd).Standalone()
	init_phase_certs_saCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_certsCmd.AddCommand(init_phase_certs_saCmd)

	carapace.Gen(init_phase_certs_saCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
	})
}
