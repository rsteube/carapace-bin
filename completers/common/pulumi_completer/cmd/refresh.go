package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh the resources in a stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refreshCmd).Standalone()
	refreshCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	refreshCmd.PersistentFlags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	refreshCmd.PersistentFlags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
	refreshCmd.PersistentFlags().String("exec-agent", "", "")
	refreshCmd.PersistentFlags().String("exec-kind", "", "")
	refreshCmd.PersistentFlags().Bool("expect-no-changes", false, "Return an error if any changes occur during this update")
	refreshCmd.Flags().BoolP("json", "j", false, "Serialize the refresh diffs, operations, and overall output as JSON")
	refreshCmd.PersistentFlags().StringP("message", "m", "", "Optional message to associate with the update operation")
	refreshCmd.PersistentFlags().IntP("parallel", "p", 2147483647, "Allow P resource operations to run in parallel at once (1 for no parallelism). Defaults to unbounded.")
	refreshCmd.PersistentFlags().Bool("show-replacement-steps", false, "Show detailed resource replacement creates and deletes instead of a single step")
	refreshCmd.PersistentFlags().Bool("show-sames", false, "Show resources that needn't be updated because they haven't changed, alongside those that do")
	refreshCmd.PersistentFlags().BoolP("skip-preview", "f", false, "Do not perform a preview before performing the refresh")
	refreshCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	refreshCmd.PersistentFlags().Bool("suppress-outputs", false, "Suppress display of stack outputs (in case they contain sensitive values)")
	refreshCmd.PersistentFlags().String("suppress-permalink", "", "Suppress display of the state permalink")
	refreshCmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Specify a single resource URN to refresh. Multiple resource can be specified using: --target urn1 --target urn2")
	refreshCmd.PersistentFlags().BoolP("yes", "y", false, "Automatically approve and perform the refresh after previewing it")
	refreshCmd.Flag("suppress-permalink").NoOptDefVal = "false"
	rootCmd.AddCommand(refreshCmd)
}
