package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()
	initCmd.Flags().Bool("copy-chunker-params", false, "copy chunker parameters from the secondary repository (useful with the copy command)")
	initCmd.Flags().String("key-hint2", "", "key ID of key to try decrypting the secondary repository first (default: $RESTIC_KEY_HINT2)")
	initCmd.Flags().String("password-command2", "", "shell `command` to obtain the secondary repository password from (default: $RESTIC_PASSWORD_COMMAND2)")
	initCmd.Flags().String("password-file2", "", "`file` to read the secondary repository password from (default: $RESTIC_PASSWORD_FILE2)")
	initCmd.Flags().String("repo2", "", "secondary `repository` to copy chunker parameters from (default: $RESTIC_REPOSITORY2)")
	initCmd.Flags().String("repository-file2", "", "`file` from which to read the secondary repository location to copy chunker parameters from (default: $RESTIC_REPOSITORY_FILE2)")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"password-command2": carapace.ActionFiles(),
		"password-file2":    carapace.ActionFiles(),
		"repo2":             action.ActionRepo(),
		"repository-file2":  carapace.ActionFiles(),
	})
}
