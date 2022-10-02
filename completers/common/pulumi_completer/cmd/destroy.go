package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:     "destroy",
	Short:   "Destroy an existing stack and its resources",
	Aliases: []string{"down"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()
	destroyCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	destroyCmd.PersistentFlags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	destroyCmd.PersistentFlags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
	destroyCmd.PersistentFlags().Bool("exclude-protected", false, "Do not destroy protected resources. Destroy all other resources.")
	destroyCmd.PersistentFlags().String("exec-agent", "", "")
	destroyCmd.PersistentFlags().String("exec-kind", "", "")
	destroyCmd.Flags().BoolP("json", "j", false, "Serialize the destroy diffs, operations, and overall output as JSON")
	destroyCmd.PersistentFlags().StringP("message", "m", "", "Optional message to associate with the destroy operation")
	destroyCmd.PersistentFlags().IntP("parallel", "p", 2147483647, "Allow P resource operations to run in parallel at once (1 for no parallelism). Defaults to unbounded.")
	destroyCmd.PersistentFlags().StringP("refresh", "r", "", "Refresh the state of the stack's resources before this update")
	destroyCmd.PersistentFlags().Bool("show-config", false, "Show configuration keys and variables")
	destroyCmd.PersistentFlags().Bool("show-replacement-steps", false, "Show detailed resource replacement creates and deletes instead of a single step")
	destroyCmd.PersistentFlags().Bool("show-sames", false, "Show resources that don't need to be updated because they haven't changed, alongside those that do")
	destroyCmd.PersistentFlags().BoolP("skip-preview", "f", false, "Do not perform a preview before performing the destroy")
	destroyCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	destroyCmd.PersistentFlags().Bool("suppress-outputs", false, "Suppress display of stack outputs (in case they contain sensitive values)")
	destroyCmd.PersistentFlags().String("suppress-permalink", "", "Suppress display of the state permalink")
	destroyCmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Specify a single resource URN to destroy. All resources necessary to destroy this target will also be destroyed. Multiple resources can be specified using: --target urn1 --target urn2. Wildcards (*, **) are also supported")
	destroyCmd.PersistentFlags().Bool("target-dependents", false, "Allows destroying of dependent targets discovered but not specified in --target list")
	destroyCmd.PersistentFlags().BoolP("yes", "y", false, "Automatically approve and perform the destroy after previewing it")
	destroyCmd.Flag("refresh").NoOptDefVal = "true"
	destroyCmd.Flag("suppress-permalink").NoOptDefVal = "false"
	rootCmd.AddCommand(destroyCmd)

	carapace.Gen(destroyCmd).FlagCompletion(carapace.ActionMap{
		//"event-log": carapace.ActionFiles(),
		"config-file": carapace.ActionFiles(),
		"exec-kind":   action.ActionExecKinds(),
		"stack":       action.ActionStacks(destroyCmd, action.StackOpts{}),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(destroyCmd)
		}),
	})
}
