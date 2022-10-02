package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate manual pages and auto-completion files (bash, fish, zsh)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()
	generateCmd.Flags().String("bash-completion", "", "write bash completion `file`")
	generateCmd.Flags().String("fish-completion", "", "write fish completion `file`")
	generateCmd.Flags().String("man", "", "write man pages to `directory`")
	generateCmd.Flags().String("zsh-completion", "", "write zsh completion `file`")
	rootCmd.AddCommand(generateCmd)

	carapace.Gen(generateCmd).FlagCompletion(carapace.ActionMap{
		"bash-completion": carapace.ActionFiles(),
		"fish-completion": carapace.ActionFiles(),
		"man":             carapace.ActionDirectories(),
		"zsh-completion":  carapace.ActionFiles(),
	})
}
