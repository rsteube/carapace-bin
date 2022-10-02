package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new template in the current folder with the name given as name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	newCmd.Flags().StringP("append", "a", "", "Append to existing YAML file")
	newCmd.Flags().String("cpu-limit", "", "Set a limit for the CPU")
	newCmd.Flags().String("cpu-request", "", "Set a request value for the CPU")
	newCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL to store in YAML stack file")
	newCmd.Flags().String("handler", "", "directory the handler will be written to")
	newCmd.Flags().String("lang", "", "Language or template to use")
	newCmd.Flags().Bool("list", false, "List available languages")
	newCmd.Flags().String("memory-limit", "", "Set a limit for the memory")
	newCmd.Flags().String("memory-request", "", "Set a request or the memory")
	newCmd.Flags().StringP("prefix", "p", "", "Set prefix for the function image")
	newCmd.Flags().BoolP("quiet", "q", false, "Skip template notes")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"append":  carapace.ActionFiles(),
		"handler": carapace.ActionDirectories(),
		"lang":    action.ActionLanguageTemplatesNew(),
	})
}
