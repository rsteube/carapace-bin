package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "readlink",
	Short: "print resolved symbolic links or canonical file names",
	Long:  "https://linux.die.net/man/1/readlink",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("canonicalize", "f", false, "canonicalize by following every symlink in")
	rootCmd.Flags().BoolP("canonicalize-existing", "e", false, "canonicalize by following every symlink in")
	rootCmd.Flags().BoolP("canonicalize-missing", "m", false, "canonicalize by following every symlink in")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("no-newline", "n", false, "do not output the trailing delimiter")
	rootCmd.Flags().BoolP("quiet", "q", false, "")
	rootCmd.Flags().BoolP("silent", "s", false, "suppress most error messages (on by default)")
	rootCmd.Flags().BoolP("verbose", "v", false, "report error messages")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero", "z", false, "end each output line with NUL, not newline")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
