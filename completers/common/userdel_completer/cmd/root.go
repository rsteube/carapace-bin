package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "userdel",
	Short: "delete a user account and related files",
	Long:  "https://linux.die.net/man/8/userdel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "force removal of files,")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("prefix", "P", "", "prefix directory where are located the /etc/* files")
	rootCmd.Flags().BoolP("remove", "r", false, "remove home directory and mail spool")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"prefix": carapace.ActionDirectories(),
		"root":   carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUsers(),
	)
}
