package cmd

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_createCmd).Standalone()
	codespace_createCmd.Flags().StringP("branch", "b", "", "repository branch")
	codespace_createCmd.Flags().Bool("default-permissions", false, "do not prompt to accept additional permissions requested by the codespace")
	codespace_createCmd.Flags().String("devcontainer-path", "", "path to the devcontainer.json file to use when creating codespace")
	codespace_createCmd.Flags().Duration("idle-timeout", 0, "allowed inactivity before codespace is stopped, e.g. \"10m\", \"1h\"")
	codespace_createCmd.Flags().StringP("location", "l", "", "location: {EastUs|SouthEastAsia|WestEurope|WestUs2} (determined automatically if not provided)")
	codespace_createCmd.Flags().StringP("machine", "m", "", "hardware specifications for the VM")
	codespace_createCmd.Flags().StringP("repo", "r", "", "repository name with owner: user/repo")
	codespace_createCmd.Flags().Duration("retention-period", 0, "allowed time after shutting down before the codespace is automatically deleted (maximum 30 days), e.g. \"1h\", \"72h\"")
	codespace_createCmd.Flags().BoolP("status", "s", false, "show status of post-create command and dotfiles")
	codespaceCmd.AddCommand(codespace_createCmd)

	// TODO devcontainer-path
	carapace.Gen(codespace_createCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(codespace_createCmd),
		"devcontainer-path": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			path := filepath.Dir(c.CallbackValue)
			path = strings.TrimPrefix(path, "/")
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return action.ActionContents(codespace_createCmd, path, codespace_createCmd.Flag("branch").Value.String())
			})
		}),
		"location": carapace.ActionValues("EastUs", "SouthEastAsia", "WestEurope", "WestUs2"),
		"machine":  action.ActionCodespaceMachines(),
		"repo":     action.ActionOwnerRepositories(codespace_createCmd),
	})
}
