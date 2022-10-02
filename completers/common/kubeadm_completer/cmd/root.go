package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubeadm",
	Short: "kubeadm: easily bootstrap a secure Kubernetes cluster",
	Long:  "https://kubernetes.io/docs/reference/setup-tools/kubeadm/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("add-dir-header", false, "If true, adds the file directory to the header of the log messages")
	rootCmd.PersistentFlags().Bool("alsologtostderr", false, "log to standard error as well as files")
	rootCmd.PersistentFlags().String("log-backtrace-at", "", "when logging hits line file:N, emit a stack trace")
	rootCmd.PersistentFlags().String("log-dir", "", "If non-empty, write log files in this directory")
	rootCmd.PersistentFlags().String("log-file", "", "If non-empty, use this log file")
	rootCmd.PersistentFlags().Uint64("log-file-max-size", 1800, "Defines the maximum size a log file can grow to. Unit is megabytes. If the value is 0, the maximum file size is unlimited.")
	rootCmd.PersistentFlags().Bool("logtostderr", true, "log to standard error instead of files")
	rootCmd.PersistentFlags().Bool("one-output", false, "If true, only write logs to their native severity level (vs also writing to each lower severity level)")
	rootCmd.PersistentFlags().String("rootfs", "", "[EXPERIMENTAL] The path to the 'real' host root filesystem.")
	rootCmd.PersistentFlags().Bool("skip-headers", false, "If true, avoid header prefixes in the log messages")
	rootCmd.PersistentFlags().Bool("skip-log-headers", false, "If true, avoid headers when opening log files")
	rootCmd.PersistentFlags().String("stderrthreshold", "", "logs at or above this threshold go to stderr")
	rootCmd.PersistentFlags().StringP("v", "v", "", "number for the log level verbosity")
	rootCmd.PersistentFlags().String("vmodule", "", "comma-separated list of pattern=N settings for file-filtered logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-dir":         carapace.ActionDirectories(),
		"log-file":        carapace.ActionFiles(),
		"rootfs":          carapace.ActionFiles(),
		"stderrthreshold": carapace.ActionValues("INFO", "WARNING", "ERROR", "FATAL"),
	})
}
