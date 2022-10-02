package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "deploys code in this environment to a configured destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		carapace.ActionValues("atlas", "ftp", "heroku", "sftp", "local-exec"),
	)
}
