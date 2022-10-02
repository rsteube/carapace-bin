package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_portsCmd = &cobra.Command{
	Use:   "ports",
	Short: "List ports in a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_portsCmd).Standalone()
	codespace_portsCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_portsCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	codespace_portsCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	codespace_portsCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	codespaceCmd.AddCommand(codespace_portsCmd)

	carapace.Gen(codespace_portsCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionCodespacePortFields().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
