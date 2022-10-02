package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "[PREVIEW] Continuously update the resources in a stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()
	watchCmd.PersistentFlags().StringArrayP("config", "c", []string{}, "Config to use during the update")
	watchCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	watchCmd.PersistentFlags().Bool("config-path", false, "Config keys contain a path to a property in a map or list to set")
	watchCmd.PersistentFlags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	watchCmd.PersistentFlags().String("exec-kind", "", "")
	watchCmd.PersistentFlags().StringP("message", "m", "", "Optional message to associate with each update operation")
	watchCmd.PersistentFlags().IntP("parallel", "p", 2147483647, "Allow P resource operations to run in parallel at once (1 for no parallelism). Defaults to unbounded.")
	watchCmd.PersistentFlags().StringArray("path", []string{}, "Specify one or more relative or absolute paths that need to be watched. A path can point to a folder or a file. Defaults to working directory")
	watchCmd.PersistentFlags().StringSlice("policy-pack", []string{}, "Run one or more policy packs as part of each update")
	watchCmd.PersistentFlags().StringSlice("policy-pack-config", []string{}, "Path to JSON file containing the config for the policy pack of the corresponding \"--policy-pack\" flag")
	watchCmd.PersistentFlags().BoolP("refresh", "r", false, "Refresh the state of the stack's resources before each update")
	watchCmd.PersistentFlags().String("secrets-provider", "default", "The type of the provider that should be used to encrypt and decrypt secrets (possible choices: default, passphrase, awskms, azurekeyvault, gcpkms, hashivault). Onlyused when creating a new stack from an existing template")
	watchCmd.PersistentFlags().Bool("show-config", false, "Show configuration keys and variables")
	watchCmd.PersistentFlags().Bool("show-replacement-steps", false, "Show detailed resource replacement creates and deletes instead of a single step")
	watchCmd.PersistentFlags().Bool("show-sames", false, "Show resources that don't need be updated because they haven't changed, alongside those that do")
	watchCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	rootCmd.AddCommand(watchCmd)

	carapace.Gen(watchCmd).FlagCompletion(carapace.ActionMap{
		"config-file":        carapace.ActionFiles(),
		"path":               carapace.ActionFiles(),
		"policy-pack-config": carapace.ActionFiles(),
		"secrets-provider":   action.ActionSecretsProvider(),
		"stack":              action.ActionStacks(watchCmd, action.StackOpts{}),
	})
}
