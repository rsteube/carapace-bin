package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swarm_caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Display and rotate the root CA",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_caCmd).Standalone()

	swarm_caCmd.Flags().String("ca-cert", "", "Path to the PEM-formatted root CA certificate to use for the")
	swarm_caCmd.Flags().String("ca-key", "", "Path to the PEM-formatted root CA key to use for the new cluster")
	swarm_caCmd.Flags().String("cert-expiry", "", "Validity period for node certificates (ns|us|ms|s|m|h)")
	swarm_caCmd.Flags().BoolP("detach", "d", false, "Exit immediately instead of waiting for the root rotation to")
	swarm_caCmd.Flags().String("external-ca", "", "Specifications of one or more certificate signing endpoints")
	swarm_caCmd.Flags().BoolP("quiet", "q", false, "Suppress progress output")
	swarm_caCmd.Flags().Bool("rotate", false, "Rotate the swarm CA - if no certificate or key are provided,")
	swarmCmd.AddCommand(swarm_caCmd)

	carapace.Gen(swarm_caCmd).FlagCompletion(carapace.ActionMap{
		"ca-cert": carapace.ActionFiles(".crt"),
		"ca-key":  carapace.ActionFiles(".key"),
	})
}
