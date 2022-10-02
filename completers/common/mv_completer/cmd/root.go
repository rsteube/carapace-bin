package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mv",
	Short: "move (rename) files",
	Long:  "https://linux.die.net/man/1/mv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("b", "b", false, "like --backup but does not accept an argument")
	rootCmd.Flags().String("backup", "", "make a backup of each existing destination file")
	rootCmd.Flags().BoolP("context", "Z", false, "set SELinux security context of destination file to default type")
	rootCmd.Flags().BoolP("force", "f", false, "do not prompt before overwriting")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("interactive", "i", false, "prompt before overwrite")
	rootCmd.Flags().BoolP("no-clobber", "n", false, "do not overwrite an existing file")
	rootCmd.Flags().BoolP("no-target-directory", "T", false, "treat DEST as a normal file")
	rootCmd.Flags().Bool("strip-trailing-slashes", false, "remove any trailing slashes from each SOURCE argument")
	rootCmd.Flags().StringP("suffix", "S", "", "override the usual backup suffix")
	rootCmd.Flags().StringP("target-directory", "t", "", "move all SOURCE arguments into DIRECTORY")
	rootCmd.Flags().BoolP("update", "u", false, "move only when the SOURCE file is newer than the destination file or when the destination file is missing")
	rootCmd.Flags().BoolP("verbose", "v", false, "explain what is being done")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backup": carapace.ActionValuesDescribed(
			"none", "never make backups (even if --backup is given)",
			"off", "never make backups (even if --backup is given)",
			"numbered", "make numbered backups",
			"t", "make numbered backups",
			"existing", "numbered if numbered backups exist, simple otherwise",
			"nil", "numbered if numbered backups exist, simple otherwise",
			"simple", "always make simple backups",
			"never", "always make simple backups",
		),
		"target-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

}
