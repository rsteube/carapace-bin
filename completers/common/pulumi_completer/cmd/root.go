package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pulumi",
	Short: "Pulumi command line",
	Long:  "https://www.pulumi.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().String("color", "auto", "Colorize output. Choices are: always, never, raw, auto")
	rootCmd.PersistentFlags().StringP("cwd", "C", "", "Run pulumi as if it had been started in another directory")
	rootCmd.PersistentFlags().Bool("disable-integrity-checking", false, "Disable integrity checking of checkpoint files")
	rootCmd.PersistentFlags().BoolP("emoji", "e", false, "Enable emojis in the output")
	rootCmd.Flags().BoolP("help", "h", false, "help for pulumi")
	rootCmd.PersistentFlags().Bool("logflow", false, "Flow log settings to child processes (like plugins)")
	rootCmd.PersistentFlags().Bool("logtostderr", false, "Log to stderr instead of to files")
	rootCmd.PersistentFlags().Bool("non-interactive", false, "Disable interactive mode for all commands")
	rootCmd.PersistentFlags().String("profiling", "", "Emit CPU and memory profiles and an execution trace to '[filename].[pid].{cpu,mem,trace}', respectively")
	rootCmd.PersistentFlags().String("tracing", "", "Emit tracing to the specified endpoint. Use the `file:` scheme to write tracing data to a local file")
	rootCmd.PersistentFlags().String("tracing-header", "", "Include the tracing header with the given contents.")
	rootCmd.PersistentFlags().IntP("verbose", "v", 0, "Enable verbose logging (e.g., v=3); anything >3 is very verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("always", "never", "raw", "auto").StyleF(style.ForKeyword),
		"cwd":   carapace.ActionDirectories(),
		"tracing": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("file:")
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
