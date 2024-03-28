package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap toolkit components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrapCmd).Standalone()
	bootstrapCmd.PersistentFlags().String("arch", "", "cluster architecture, available options are: (amd64, arm, arm64)")
	bootstrapCmd.PersistentFlags().String("author-email", "", "author email for Git commits")
	bootstrapCmd.PersistentFlags().String("author-name", "Flux", "author name for Git commits")
	bootstrapCmd.PersistentFlags().String("branch", "main", "Git branch")
	bootstrapCmd.PersistentFlags().String("ca-file", "", "path to TLS CA file used for validating self-signed certificates")
	bootstrapCmd.PersistentFlags().String("cluster-domain", "cluster.local", "internal cluster domain")
	bootstrapCmd.PersistentFlags().String("commit-message-appendix", "", "string to add to the commit messages, e.g. '[ci skip]'")
	bootstrapCmd.PersistentFlags().StringSlice("components", []string{"source-controller", "kustomize-controller", "helm-controller", "notification-controller"}, "list of components, accepts comma-separated values")
	bootstrapCmd.PersistentFlags().StringSlice("components-extra", []string{}, "list of components in addition to those supplied or defaulted, accepts comma-separated values")
	bootstrapCmd.PersistentFlags().String("gpg-key-id", "", "key id for selecting a particular key")
	bootstrapCmd.PersistentFlags().String("gpg-key-ring", "", "path to GPG key ring for signing commits")
	bootstrapCmd.PersistentFlags().String("gpg-passphrase", "", "passphrase for decrypting GPG private key")
	bootstrapCmd.PersistentFlags().String("image-pull-secret", "", "Kubernetes secret name used for pulling the toolkit images from a private registry")
	bootstrapCmd.PersistentFlags().String("log-level", "info", "log level, available options are: (debug, info, error)")
	bootstrapCmd.PersistentFlags().String("manifests", "", "path to the manifest directory")
	bootstrapCmd.PersistentFlags().Bool("network-policy", true, "deny ingress access to the toolkit controllers from other namespaces using network policies")
	bootstrapCmd.PersistentFlags().String("private-key-file", "", "path to a private key file used for authenticating to the Git SSH server")
	bootstrapCmd.PersistentFlags().Bool("recurse-submodules", false, "when enabled, configures the GitRepository source to initialize and include Git submodules in the artifact it produces")
	bootstrapCmd.PersistentFlags().String("registry", "ghcr.io/fluxcd", "container registry where the toolkit images are published")
	bootstrapCmd.PersistentFlags().String("secret-name", "flux-system", "name of the secret the sync credentials can be found in or stored to")
	bootstrapCmd.PersistentFlags().String("ssh-ecdsa-curve", "", "SSH ECDSA public key curve (p256, p384, p521)")
	bootstrapCmd.PersistentFlags().String("ssh-hostname", "", "SSH hostname, to be used when the SSH host differs from the HTTPS one")
	bootstrapCmd.PersistentFlags().String("ssh-key-algorithm", "ecdsa", "SSH public key algorithm (rsa, ecdsa, ed25519)")
	bootstrapCmd.PersistentFlags().String("ssh-rsa-bits", "2048", "SSH RSA public key bit size (multiplies of 8)")
	bootstrapCmd.PersistentFlags().Bool("token-auth", false, "when enabled, the personal access token will be used instead of SSH deploy key")
	bootstrapCmd.PersistentFlags().StringSlice("toleration-keys", []string{}, "list of toleration keys used to schedule the components pods onto nodes with matching taints")
	bootstrapCmd.PersistentFlags().StringP("version", "v", "", "toolkit version, when specified the manifests are downloaded from https://github.com/fluxcd/flux2/releases")
	bootstrapCmd.PersistentFlags().Bool("watch-all-namespaces", true, "watch for custom resources in all namespaces, if set to false it will only watch the namespace where the toolkit is installed")
	rootCmd.AddCommand(bootstrapCmd)
}
