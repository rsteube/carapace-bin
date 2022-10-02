package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:     "up",
	Short:   "Create or update the resources in a stack",
	Aliases: []string{"update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()
	upCmd.PersistentFlags().String("client", "", "The address of an existing language runtime host to connect to")
	upCmd.PersistentFlags().StringArrayP("config", "c", []string{}, "Config to use during the update")
	upCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	upCmd.PersistentFlags().Bool("config-path", false, "Config keys contain a path to a property in a map or list to set")
	upCmd.PersistentFlags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	upCmd.PersistentFlags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
	upCmd.PersistentFlags().String("exec-agent", "", "")
	upCmd.PersistentFlags().String("exec-kind", "", "")
	upCmd.PersistentFlags().Bool("expect-no-changes", false, "Return an error if any changes occur during this update")
	upCmd.Flags().BoolP("json", "j", false, "Serialize the update diffs, operations, and overall output as JSON")
	upCmd.PersistentFlags().StringP("message", "m", "", "Optional message to associate with the update operation")
	upCmd.PersistentFlags().IntP("parallel", "p", 2147483647, "Allow P resource operations to run in parallel at once (1 for no parallelism). Defaults to unbounded.")
	upCmd.PersistentFlags().String("plan", "", "[EXPERIMENTAL] Path to a plan file to use for the update. The update will not perform operations that exceed its plan (e.g. replacements instead of updates, or updates insteadof sames).")
	upCmd.PersistentFlags().StringSlice("policy-pack", []string{}, "Run one or more policy packs as part of this update")
	upCmd.PersistentFlags().StringSlice("policy-pack-config", []string{}, "Path to JSON file containing the config for the policy pack of the corresponding \"--policy-pack\" flag")
	upCmd.PersistentFlags().StringP("refresh", "r", "", "Refresh the state of the stack's resources before this update")
	upCmd.PersistentFlags().StringArray("replace", []string{}, "Specify resources to replace. Multiple resources can be specified using --replace urn1 --replace urn2. Wildcards (*, **) are also supported")
	upCmd.PersistentFlags().String("secrets-provider", "default", "The type of the provider that should be used to encrypt and decrypt secrets (possible choices: default, passphrase, awskms, azurekeyvault, gcpkms, hashivault). Onlyused when creating a new stack from an existing template")
	upCmd.PersistentFlags().Bool("show-config", false, "Show configuration keys and variables")
	upCmd.PersistentFlags().Bool("show-reads", false, "Show resources that are being read in, alongside those being managed directly in the stack")
	upCmd.PersistentFlags().Bool("show-replacement-steps", false, "Show detailed resource replacement creates and deletes instead of a single step")
	upCmd.PersistentFlags().Bool("show-sames", false, "Show resources that don't need be updated because they haven't changed, alongside those that do")
	upCmd.PersistentFlags().BoolP("skip-preview", "f", false, "Do not perform a preview before performing the update")
	upCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	upCmd.PersistentFlags().Bool("suppress-outputs", false, "Suppress display of stack outputs (in case they contain sensitive values)")
	upCmd.PersistentFlags().String("suppress-permalink", "", "Suppress display of the state permalink")
	upCmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Specify a single resource URN to update. Other resources will not be updated. Multiple resources can be specified using --target urn1 --target urn2. Wildcards (*, **) are also supported")
	upCmd.PersistentFlags().Bool("target-dependents", false, "Allows updating of dependent targets discovered but not specified in --target list")
	upCmd.PersistentFlags().StringArray("target-replace", []string{}, "Specify a single resource URN to replace. Other resources will not be updated. Shorthand for --target urn --replace urn.")
	upCmd.PersistentFlags().BoolP("yes", "y", false, "Automatically approve and perform the update after previewing it")
	upCmd.Flag("refresh").NoOptDefVal = "true"
	upCmd.Flag("suppress-permalink").NoOptDefVal = "false"
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"replace": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(upCmd)
		}),
		"secrets-provider": action.ActionSecretsProvider(),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(upCmd)
		}),
		"target-replace": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(upCmd)
		}),
	})

	carapace.Gen(upCmd).PositionalCompletion(
		action.ActionTemplates(),
	)
}
