package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gpasswd",
	Short: "administer /etc/group and /etc/gshadow",
	Long:  "https://linux.die.net/man/1/gpasswd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("add", "a", "", "add USER to GROUP")
	rootCmd.Flags().StringP("administrators", "A", "", "set the list of administrators for GROUP")
	rootCmd.Flags().StringP("delete", "d", "", "remove USER from GROUP")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("members", "M", "", "set the list of members of GROUP")
	rootCmd.Flags().BoolP("remove-password", "r", false, "remove the GROUP's password")
	rootCmd.Flags().BoolP("restrict", "R", false, "restrict access to GROUP to its members")
	rootCmd.Flags().StringP("root", "Q", "", "directory to chroot into")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add": os.ActionUsers(),
		"administrators": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionUsers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"delete": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return os.ActionGroupMembers(c.Args[0])
			}
			return os.ActionUsers()
		}),
		"members": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionUsers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"root": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionGroups(),
	)
}
