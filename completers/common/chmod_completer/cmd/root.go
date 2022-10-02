package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chmod",
	Short: "change file mode bits",
	Long:  "https://en.wikipedia.org/wiki/Chmod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("changes", "c", false, "like verbose but report only when a change is made")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("no-preserve-root", false, "do not treat '/' specially (the default)")
	rootCmd.Flags().Bool("preserve-root", false, "fail to operate recursively on '/'")
	rootCmd.Flags().Bool("quiet", false, "suppress most error messages")
	rootCmd.Flags().BoolP("recursive", "R", false, "change files and directories recursively")
	rootCmd.Flags().String("reference", "", "use RFILE's mode instead of MODE values")
	rootCmd.Flags().BoolP("silent", "f", false, "suppress most error messages")
	rootCmd.Flags().BoolP("verbose", "v", false, "output a diagnostic for every file processed")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("reference").Changed {
				return carapace.ActionFiles()
			} else {
				return fs.ActionFileModes()
			}
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
