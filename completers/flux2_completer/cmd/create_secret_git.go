package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_secret_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Create or update a Kubernetes secret for Git authentication",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secret_gitCmd).Standalone()
	create_secret_gitCmd.Flags().String("ca-file", "", "path to TLS CA file used for validating self-signed certificates")
	create_secret_gitCmd.Flags().StringP("password", "p", "", "basic authentication password")
	create_secret_gitCmd.Flags().String("private-key-file", "", "path to a passwordless private key file used for authenticating to the Git SSH server")
	create_secret_gitCmd.Flags().EcdsaCurve("ssh-ecdsa-curve", p384, "SSH ECDSA public key curve (p256, p384, p521)")
	create_secret_gitCmd.Flags().PublicKeyAlgorithm("ssh-key-algorithm", ecdsa, "SSH public key algorithm (rsa, ecdsa, ed25519)")
	create_secret_gitCmd.Flags().RsaKeyBits("ssh-rsa-bits", 2048, "SSH RSA public key bit size (multiplies of 8)")
	create_secret_gitCmd.Flags().String("url", "", "git address, e.g. ssh://git@host/org/repository")
	create_secret_gitCmd.Flags().StringP("username", "u", "", "basic authentication username")
	create_secretCmd.AddCommand(create_secret_gitCmd)
}
