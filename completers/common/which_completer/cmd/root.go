package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "which",
	Short: "Write the full path of COMMAND(s) to standard output",
	Long:  "https://linux.die.net/man/1/which",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all,", "a", false, "Print all matches in PATH, not just the first")
	rootCmd.Flags().Bool("help,", false, "Print this help and exit successfully.")
	rootCmd.Flags().BoolP("read-alias,", "i", false, "Read list of aliases from stdin.")
	rootCmd.Flags().Bool("read-functions", false, "Read shell functions from stdin.")
	rootCmd.Flags().Bool("show-dot", false, "Don't expand a dot to current directory in output.")
	rootCmd.Flags().Bool("show-tilde", false, "Output a tilde for HOME directory for non-root.")
	rootCmd.Flags().Bool("skip-alias", false, "Ignore option --read-alias; don't read stdin.")
	rootCmd.Flags().Bool("skip-dot", false, "Skip directories in PATH that start with a dot.")
	rootCmd.Flags().Bool("skip-functions", false, "Ignore option --read-functions; don't read stdin.")
	rootCmd.Flags().Bool("skip-tilde", false, "Skip directories in PATH that start with a tilde.")
	rootCmd.Flags().Bool("tty-only", false, "Stop processing options on the right if not on tty.")
	rootCmd.Flags().BoolP("version,", "v", false, "Print version and exit successfully.")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		os.ActionPathExecutables(),
	)
}
