package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "minikube",
	Short: "minikube quickly sets up a local Kubernetes cluster",
	Long:  "https://minikube.sigs.k8s.io/docs/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Bool("add_dir_header", false, "If true, adds the file directory to the header of the log messages")
	rootCmd.PersistentFlags().Bool("alsologtostderr", false, "log to standard error as well as files")
	rootCmd.PersistentFlags().StringP("bootstrapper", "b", "kubeadm", "The name of the cluster bootstrapper that will set up the Kubernetes cluster.")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "")
	rootCmd.PersistentFlags().String("log_backtrace_at", ":0", "when logging hits line file:N, emit a stack trace")
	rootCmd.PersistentFlags().String("log_dir", "", "If non-empty, write log files in this directory")
	rootCmd.PersistentFlags().String("log_file", "/tmp/minikube_d0bab2a63235e38c4dc1581a4373fd036e74a8bc_0.log", "If non-empty, use this log file")
	rootCmd.PersistentFlags().Uint64("log_file_max_size", 1800, "Defines the maximum size a log file can grow to. Unit is megabytes. If the value is 0, the maximum file size is unlimited.")
	rootCmd.PersistentFlags().Bool("logtostderr", false, "log to standard error instead of files")
	rootCmd.PersistentFlags().Bool("one_output", false, "If true, only write logs to their native severity level (vs also writing to each lower severity level)")
	rootCmd.PersistentFlags().StringP("profile", "p", "minikube", "The name of the minikube VM being used. This can be set to allow having multiple instances of minikube independently.")
	rootCmd.PersistentFlags().Bool("skip_headers", false, "If true, avoid header prefixes in the log messages")
	rootCmd.PersistentFlags().Bool("skip_log_headers", false, "If true, avoid headers when opening log files")
	rootCmd.PersistentFlags().Int("stderrthreshold", 2, "logs at or above this threshold go to stderr")
	rootCmd.PersistentFlags().String("user", "", "Specifies the user executing the operation. Useful for auditing operations executed by 3rd party tools. Defaults to the operating system username.")
	rootCmd.PersistentFlags().IntP("v", "v", 0, "number for the log level verbosity")
	rootCmd.PersistentFlags().StringArray("vmodule", []string{}, "comma-separated list of pattern=N settings for file-filtered logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log_dir":  carapace.ActionDirectories(),
		"log_file": carapace.ActionFiles(),
		"user":     os.ActionUsers(),
	})
}
