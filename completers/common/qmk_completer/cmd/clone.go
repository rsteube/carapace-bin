package cmd

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a qmk_firmware fork",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().String("baseurl", "https://github.com", "The URL all git operations start from")
	cloneCmd.Flags().StringP("branch", "b", "", "The branch to clone")
	cloneCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			url := cloneCmd.Flag("baseurl").Value.String()
			fork := "qmk/qmk_firmware"
			if len(c.Args) > 0 {
				fork = c.Args[0]
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: fmt.Sprintf("%v/%v", url, fork), Branches: true, Tags: true})
		}),
	})

	carapace.Gen(cloneCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !cloneCmd.Flag("baseurl").Changed {
				return gh.ActionOwnerRepositories(gh.HostOpts{}) // TODO support different basurls
			}
			return carapace.ActionValues()
		}),
		carapace.ActionDirectories(),
	)
}
