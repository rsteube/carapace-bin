package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mktemp",
	Short: "create a temporary file or directory",
	Long:  "https://linux.die.net/man/1/mktemp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("directory", "d", false, "create a directory, not a file")
	rootCmd.Flags().BoolP("dry-run", "u", false, "do not create anything; merely print a name (unsafe)")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress diagnostics about file/dir-creation failure")
	rootCmd.Flags().String("suffix", "", "append SUFF to TEMPLATE")
	rootCmd.Flags().BoolS("t", "t", false, "interpret TEMPLATE as a single file name component")
	rootCmd.Flags().StringP("tmpdir", "p", "", "interpret TEMPLATE relative to DIR")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	rootCmd.Flag("tmpdir").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"tmpdir": carapace.ActionDirectories(),
	})
}
