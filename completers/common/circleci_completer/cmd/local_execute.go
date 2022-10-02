package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var local_executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Run a job in a container on the local machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(local_executeCmd).Standalone()
	local_executeCmd.Flags().String("branch", "", "Git branch")
	local_executeCmd.Flags().String("checkout-key", "~/.ssh/id_rsa", "Git Checkout key")
	local_executeCmd.Flags().StringP("config", "c", ".circleci/config.yml", "config file")
	local_executeCmd.Flags().StringArrayP("env", "e", []string{}, "Set environment variables, e.g. `-e VAR=VAL`")
	local_executeCmd.Flags().Int("index", 0, "node index of parallelism")
	local_executeCmd.Flags().String("job", "build", "job to be executed")
	local_executeCmd.Flags().Int("node-total", 1, "total number of parallel nodes")
	local_executeCmd.Flags().StringP("org-slug", "o", "", "organization slug (for example: github/example-org), used when a config depends on private orbs belonging to that org")
	local_executeCmd.Flags().String("repo-url", "", "Git Url")
	local_executeCmd.Flags().String("revision", "", "Git Revision")
	local_executeCmd.Flags().Bool("skip-checkout", true, "use local path as-is")
	local_executeCmd.Flags().StringArrayP("volume", "v", []string{}, "Volume bind-mounting")
	localCmd.AddCommand(local_executeCmd)

	// TODO flag completion
	carapace.Gen(local_executeCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := local_executeCmd.Flag("repo-url"); flag.Changed {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: flag.Value.String(), Branches: true, Tags: true})
			}
			return carapace.ActionValues()
		}),
		"checkout-key": carapace.ActionFiles(),
		"config":       carapace.ActionFiles(),
	})
}
