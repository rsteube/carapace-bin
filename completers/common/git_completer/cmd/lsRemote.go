package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var lsRemoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "List references in a remote repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsRemoteCmd).Standalone()

	lsRemoteCmd.Flags().Bool("exit-code", false, "exit with exit code 2 if no matching refs are found")
	lsRemoteCmd.Flags().Bool("get-url", false, "take url.<base>.insteadOf into account")
	lsRemoteCmd.Flags().BoolP("heads", "h", false, "limit to heads")
	lsRemoteCmd.Flags().BoolP("quiet", "q", false, "do not print remote URL")
	lsRemoteCmd.Flags().Bool("refs", false, "do not show peeled tags")
	lsRemoteCmd.Flags().StringP("server-option", "o", "", "option to transmit")
	lsRemoteCmd.Flags().String("sort", "", "field name to sort on")
	lsRemoteCmd.Flags().Bool("symref", false, "show underlying ref in addition to the object pointed by it")
	lsRemoteCmd.Flags().BoolP("tags", "t", false, "limit to tags")
	lsRemoteCmd.Flags().String("upload-pack", "", "path of git-upload-pack on the remote host")
	rootCmd.AddCommand(lsRemoteCmd)

	carapace.Gen(lsRemoteCmd).FlagCompletion(carapace.ActionMap{
		"sort": git.ActionFieldNames(),
	})

	carapace.Gen(lsRemoteCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(lsRemoteCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[0], Branches: true, Tags: true})
		}),
	)
}
