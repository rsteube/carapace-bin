package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:     "preview",
	Short:   "Show a preview of updates to a stack's resources",
	Aliases: []string{"pre"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(previewCmd).Standalone()
	previewCmd.PersistentFlags().String("client", "", "The address of an existing language runtime host to connect to")
	previewCmd.PersistentFlags().StringArrayP("config", "c", []string{}, "Config to use during the preview")
	previewCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	previewCmd.PersistentFlags().Bool("config-path", false, "Config keys contain a path to a property in a map or list to set")
	previewCmd.PersistentFlags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	previewCmd.PersistentFlags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
	previewCmd.PersistentFlags().String("exec-agent", "", "")
	previewCmd.PersistentFlags().String("exec-kind", "", "")
	previewCmd.PersistentFlags().Bool("expect-no-changes", false, "Return an error if any changes are proposed by this preview")
	previewCmd.Flags().BoolP("json", "j", false, "Serialize the preview diffs, operations, and overall output as JSON")
	previewCmd.PersistentFlags().StringP("message", "m", "", "Optional message to associate with the preview operation")
	previewCmd.PersistentFlags().IntP("parallel", "p", 2147483647, "Allow P resource operations to run in parallel at once (1 for no parallelism). Defaults to unbounded.")
	previewCmd.PersistentFlags().StringSlice("policy-pack", []string{}, "Run one or more policy packs as part of this update")
	previewCmd.PersistentFlags().StringSlice("policy-pack-config", []string{}, "Path to JSON file containing the config for the policy pack of the corresponding \"--policy-pack\" flag")
	previewCmd.PersistentFlags().StringP("refresh", "r", "", "Refresh the state of the stack's resources before this update")
	previewCmd.PersistentFlags().StringArray("replace", []string{}, "Specify resources to replace. Multiple resources can be specified using --replace urn1 --replace urn2")
	previewCmd.PersistentFlags().String("save-plan", "", "[EXPERIMENTAL] Save the operations proposed by the preview to a plan file at the given path")
	previewCmd.PersistentFlags().Bool("show-config", false, "Show configuration keys and variables")
	previewCmd.PersistentFlags().Bool("show-reads", false, "Show resources that are being read in, alongside those being managed directly in the stack")
	previewCmd.PersistentFlags().Bool("show-replacement-steps", false, "Show detailed resource replacement creates and deletes instead of a single step")
	previewCmd.PersistentFlags().Bool("show-sames", false, "Show resources that needn't be updated because they haven't changed, alongside those that do")
	previewCmd.Flags().Bool("show-secrets", false, "Emit secrets in plaintext in the plan file. Defaults to `false`")
	previewCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	previewCmd.PersistentFlags().Bool("suppress-outputs", false, "Suppress display of stack outputs (in case they contain sensitive values)")
	previewCmd.PersistentFlags().String("suppress-permalink", "", "Suppress display of the state permalink")
	previewCmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Specify a single resource URN to update. Other resources will not be updated. Multiple resources can be specified using --target urn1 --target urn2")
	previewCmd.PersistentFlags().Bool("target-dependents", false, "Allows updating of dependent targets discovered but not specified in --target list")
	previewCmd.PersistentFlags().StringArray("target-replace", []string{}, "Specify a single resource URN to replace. Other resources will not be updated. Shorthand for --target urn --replace urn.")
	previewCmd.Flag("refresh").NoOptDefVal = "true"
	previewCmd.Flag("suppress-permalink").NoOptDefVal = "false"
	rootCmd.AddCommand(previewCmd)

	carapace.Gen(previewCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"replace": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(previewCmd)
		}),
		"save-plan": carapace.ActionFiles(),
		"stack":     action.ActionStacks(previewCmd, action.StackOpts{}),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(previewCmd)
		}),
		"target-replace": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(previewCmd)
		}),
	})
}
