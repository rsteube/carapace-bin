package cmd

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your computer for qmk_firmware",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().String("baseurl", "https://github.com", "The URL all git operations start from")
	setupCmd.Flags().StringP("branch", "b", "", "The branch to clone")
	setupCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	setupCmd.Flags().StringP("home", "H", "", "The location for QMK Firmware")
	setupCmd.Flags().BoolP("no", "n", false, "Answer no to all questions")
	setupCmd.Flags().BoolP("yes", "y", false, "Answer yes to all questions")
	rootCmd.AddCommand(setupCmd)

	carapace.Gen(setupCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			url := setupCmd.Flag("baseurl").Value.String()
			fork := "qmk/qmk_firmware"
			if len(c.Args) > 0 {
				fork = c.Args[0]
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: fmt.Sprintf("%v/%v", url, fork), Branches: true, Tags: true})
		}),
		"home": carapace.ActionDirectories(),
	})

	carapace.Gen(setupCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !setupCmd.Flag("baseurl").Changed {
				return gh.ActionOwnerRepositories(gh.HostOpts{}) // TODO support different baseurl
			}
			return carapace.ActionValues()
		}),
	)
}
