package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helm",
	Short: "The Helm package manager for Kubernetes.",
	Long:  "https://helm.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Bool("add-dir-header", false, "If true, adds the file directory to the header of the log messages")
	rootCmd.PersistentFlags().Bool("alsologtostderr", false, "log to standard error as well as files")
	rootCmd.PersistentFlags().Bool("debug", false, "enable verbose output")
	rootCmd.Flags().BoolP("help", "h", false, "help for helm")
	rootCmd.PersistentFlags().String("kube-apiserver", "", "the address and the port for the Kubernetes API server")
	rootCmd.PersistentFlags().StringArray("kube-as-group", []string{}, "group to impersonate for the operation, this flag can be repeated to specify multiple groups.")
	rootCmd.PersistentFlags().String("kube-as-user", "", "username to impersonate for the operation")
	rootCmd.PersistentFlags().String("kube-ca-file", "", "the certificate authority file for the Kubernetes API server connection")
	rootCmd.PersistentFlags().String("kube-context", "", "name of the kubeconfig context to use")
	rootCmd.PersistentFlags().String("kube-token", "", "bearer token used for authentication")
	rootCmd.PersistentFlags().String("kubeconfig", "", "path to the kubeconfig file")
	rootCmd.PersistentFlags().String("log-backtrace-at", ":0", "when logging hits line file:N, emit a stack trace")
	rootCmd.PersistentFlags().String("log-dir", "", "If non-empty, write log files in this directory")
	rootCmd.PersistentFlags().String("log-file", "", "If non-empty, use this log file")
	rootCmd.PersistentFlags().Uint64("log-file-max-size", 1800, "Defines the maximum size a log file can grow to. Unit is megabytes. If the value is 0, the maximum file size is unlimited.")
	rootCmd.PersistentFlags().Bool("logtostderr", true, "log to standard error instead of files")
	rootCmd.PersistentFlags().StringP("namespace", "n", "", "namespace scope for this request")
	rootCmd.PersistentFlags().Bool("one-output", false, "If true, only write logs to their native severity level (vs also writing to each lower severity level)")
	rootCmd.PersistentFlags().String("registry-config", "/home/user/.config/helm/registry.json", "path to the registry config file")
	rootCmd.PersistentFlags().String("repository-cache", "/home/user/.cache/helm/repository", "path to the file containing cached repository indexes")
	rootCmd.PersistentFlags().String("repository-config", "/home/user/.config/helm/repositories.yaml", "path to the file containing repository names and URLs")
	rootCmd.PersistentFlags().Bool("skip-headers", false, "If true, avoid header prefixes in the log messages")
	rootCmd.PersistentFlags().Bool("skip-log-headers", false, "If true, avoid headers when opening log files")
	rootCmd.PersistentFlags().Int("stderrthreshold", 2, "logs at or above this threshold go to stderr")
	rootCmd.PersistentFlags().IntP("v", "v", 0, "number for the log level verbosity")
	rootCmd.PersistentFlags().StringSlice("vmodule", []string{}, "comma-separated list of pattern=N settings for file-filtered logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"kube-as-group": os.ActionGroups(),
		"kube-as-user":  os.ActionUsers(),
		"kube-ca-file":  carapace.ActionFiles(),
		// TODO "kube-context"
		"kubeconfig": carapace.ActionFiles(),
		"log-dir":    carapace.ActionDirectories(),
		"log-file":   carapace.ActionFiles(),
		// TODO namespace
		"registry-config":  carapace.ActionFiles(),
		"repository-cache": carapace.ActionFiles(),
	})
}
