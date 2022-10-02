package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Pulumi project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()
	newCmd.PersistentFlags().StringArrayP("config", "c", []string{}, "Config to save")
	newCmd.PersistentFlags().Bool("config-path", false, "Config keys contain a path to a property in a map or list to set")
	newCmd.PersistentFlags().StringP("description", "d", "", "The project description; if not specified, a prompt will request it")
	newCmd.PersistentFlags().String("dir", "", "The location to place the generated project; if not specified, the current directory is used")
	newCmd.PersistentFlags().BoolP("force", "f", false, "Forces content to be generated even if it would change existing files")
	newCmd.PersistentFlags().BoolP("generate-only", "g", false, "Generate the project only; do not create a stack, save config, or install dependencies")
	newCmd.PersistentFlags().StringP("name", "n", "", "The project name; if not specified, a prompt will request it")
	newCmd.PersistentFlags().BoolP("offline", "o", false, "Use locally cached templates without making any network requests")
	newCmd.PersistentFlags().String("secrets-provider", "default", "The type of the provider that should be used to encrypt and decrypt secrets (possible choices: default, passphrase, awskms, azurekeyvault, gcpkms, hashivault)")
	newCmd.PersistentFlags().StringP("stack", "s", "", "The stack name; either an existing stack or stack to create; if not specified, a prompt will request it")
	newCmd.PersistentFlags().BoolP("yes", "y", false, "Skip prompts and proceed with default values")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"dir":              carapace.ActionDirectories(),
		"secrets-provider": action.ActionSecretsProvider(),
		"stack":            action.ActionStacks(newCmd, action.StackOpts{}),
	})

	carapace.Gen(newCmd).PositionalCompletion(
		action.ActionTemplates(),
	)
}
