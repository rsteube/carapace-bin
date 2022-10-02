package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chroot",
	Short: "run command or interactive shell with special root directory",
	Long:  "https://en.wikipedia.org/wiki/Chroot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("groups", "", "specify supplementary groups as g1,g2,..,gN")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("skip-chdir", false, "do not change working directory to '/'")
	rootCmd.Flags().String("userspec", "", "specify user and group (ID or name) to use")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"groups": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionGroups().Invoke(c).Filter(c.Parts).ToA()
		}),
		"userspec": os.ActionUserGroup(),
	})
}
