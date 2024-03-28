package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_source_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Create or update a GitRepository source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_source_gitCmd).Standalone()
	create_source_gitCmd.Flags().String("branch", "", "git branch")
	create_source_gitCmd.Flags().String("ca-file", "", "path to TLS CA file used for validating self-signed certificates")
	create_source_gitCmd.Flags().String("git-implementation", "", "the Git implementation to use, available options are: (go-git, libgit2)")
	create_source_gitCmd.Flags().StringP("password", "p", "", "basic authentication password")
	create_source_gitCmd.Flags().String("private-key-file", "", "path to a passwordless private key file used for authenticating to the Git SSH server")
	create_source_gitCmd.Flags().Bool("recurse-submodules", false, "when enabled, configures the GitRepository source to initialize and include Git submodules in the artifact it produces")
	create_source_gitCmd.Flags().String("secret-ref", "", "the name of an existing secret containing SSH or basic credentials")
	create_source_gitCmd.Flags().BoolP("silent", "s", false, "assumes the deploy key is already setup, skips confirmation")
	create_source_gitCmd.Flags().EcdsaCurve("ssh-ecdsa-curve", p384, "SSH ECDSA public key curve (p256, p384, p521)")
	create_source_gitCmd.Flags().PublicKeyAlgorithm("ssh-key-algorithm", ecdsa, "SSH public key algorithm (rsa, ecdsa, ed25519)")
	create_source_gitCmd.Flags().RsaKeyBits("ssh-rsa-bits", 2048, "SSH RSA public key bit size (multiplies of 8)")
	create_source_gitCmd.Flags().String("tag", "", "git tag")
	create_source_gitCmd.Flags().String("tag-semver", "", "git tag semver range")
	create_source_gitCmd.Flags().String("url", "", "git address, e.g. ssh://git@host/org/repository")
	create_source_gitCmd.Flags().StringP("username", "u", "", "basic authentication username")
	create_sourceCmd.AddCommand(create_source_gitCmd)
}
