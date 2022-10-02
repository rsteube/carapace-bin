package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mknod",
	Short: "make block or character special files",
	Long:  "https://linux.die.net/man/3/mknod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("Z", "Z", false, "set the SELinux security context to default type")
	rootCmd.Flags().String("context", "", "like -Z, or if CTX is specified then set the SELinux")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("mode", "m", "", "set file permission bits to MODE, not a=rw - umask")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"mode": fs.ActionFileModes(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValuesDescribed("b", "block (buffered) special file", "c", "FIFO", "u", "character (unbuffered) special file", "p", "character (unbuffered) special file"),
	)
}
