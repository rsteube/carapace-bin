package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flux",
	Short: "Command line utility for assembling Kubernetes CD pipelines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().String("as", "", "Username to impersonate for the operation. User could be a regular user or a service account in a namespace.")
	rootCmd.PersistentFlags().StringArray("as-group", []string{}, "Group to impersonate for the operation, this flag can be repeated to specify multiple groups.")
	rootCmd.PersistentFlags().String("as-uid", "", "UID to impersonate for the operation.")
	rootCmd.PersistentFlags().String("cache-dir", "/home/rsteube/.kube/cache", "Default cache directory")
	rootCmd.PersistentFlags().String("certificate-authority", "", "Path to a cert file for the certificate authority")
	rootCmd.PersistentFlags().String("client-certificate", "", "Path to a client certificate file for TLS")
	rootCmd.PersistentFlags().String("client-key", "", "Path to a client key file for TLS")
	rootCmd.PersistentFlags().String("cluster", "", "The name of the kubeconfig cluster to use")
	rootCmd.PersistentFlags().String("context", "", "The name of the kubeconfig context to use")
	rootCmd.Flags().BoolP("help", "h", false, "help for flux")
	rootCmd.PersistentFlags().Bool("insecure-skip-tls-verify", false, "If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure")
	rootCmd.PersistentFlags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	rootCmd.PersistentFlags().StringP("namespace", "n", "flux-system", "If present, the namespace scope for this CLI request")
	rootCmd.PersistentFlags().String("server", "", "The address and port of the Kubernetes API server")
	rootCmd.PersistentFlags().String("timeout", "", "timeout for this operation")
	rootCmd.PersistentFlags().String("tls-server-name", "", "Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used")
	rootCmd.PersistentFlags().String("token", "", "Bearer token for authentication to the API server")
	rootCmd.PersistentFlags().String("user", "", "The name of the kubeconfig user to use")
	rootCmd.PersistentFlags().Bool("verbose", false, "print generated objects")
	rootCmd.Flags().BoolP("version", "v", false, "version for flux")
}
