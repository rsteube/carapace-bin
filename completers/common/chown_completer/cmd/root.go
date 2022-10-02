package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chown",
	Short: "change file owner and group",
	Long:  "https://en.wikipedia.org/wiki/Chown",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "if a command line argument is a symbolic link to a directory, traverse it")
	rootCmd.Flags().BoolS("L", "L", false, "traverse every symbolic link to a directory encountered")
	rootCmd.Flags().BoolS("P", "P", false, "do not traverse any symbolic links (default)")
	rootCmd.Flags().BoolP("changes", "c", false, "like verbose but report only when a change is made")
	rootCmd.Flags().Bool("dereference", false, "affect the referent of each symbolic link")
	rootCmd.Flags().String("from", "", "change the owner and/or group of each file only if its current owner and/or group match those specified here.")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("no-dereference", "h", false, "affect symbolic links instead of any referenced file")
	rootCmd.Flags().Bool("no-preserve-root", false, "do not treat '/' specially (the default)")
	rootCmd.Flags().Bool("preserve-root", false, "fail to operate recursively on '/'")
	rootCmd.Flags().BoolP("recursive", "R", false, "operate on files and directories recursively")
	rootCmd.Flags().String("reference", "", "use RFILE's owner and group rather than specifying OWNER:GROUP values")
	rootCmd.Flags().StringP("silent", "f", "", "suppress most error messages")
	rootCmd.Flags().BoolP("verbose", "v", false, "output a diagnostic for every file processed")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"from":      os.ActionUserGroup(),
		"reference": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("reference").Changed {
				return carapace.ActionFiles()
			} else {
				return os.ActionUserGroup()
			}
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
