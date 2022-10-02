package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import resources into an existing stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()
	importCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	importCmd.PersistentFlags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	importCmd.PersistentFlags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
	importCmd.PersistentFlags().String("exec-agent", "", "")
	importCmd.PersistentFlags().String("exec-kind", "", "")
	importCmd.PersistentFlags().StringP("file", "f", "", "The path to a JSON-encoded file containing a list of resources to import")
	importCmd.PersistentFlags().Bool("generate-code", true, "Generate resource declaration code for the imported resources")
	importCmd.PersistentFlags().StringP("message", "m", "", "Optional message to associate with the update operation")
	importCmd.PersistentFlags().StringP("out", "o", "", "The path to the file that will contain the generated resource declarations")
	importCmd.PersistentFlags().IntP("parallel", "p", 2147483647, "Allow P resource operations to run in parallel at once (1 for no parallelism). Defaults to unbounded.")
	importCmd.PersistentFlags().String("parent", "", "The name and URN of the parent resource in the format name=urn, where name is the variable name of the parent resource")
	importCmd.PersistentFlags().StringSlice("properties", []string{}, "The property names to use for the import in the format name1,name2")
	importCmd.PersistentFlags().Bool("protect", true, "Allow resources to be imported with protection from deletion enabled")
	importCmd.PersistentFlags().String("provider", "", "The name and URN of the provider to use for the import in the format name=urn, where name is the variable name for the provider resource")
	importCmd.PersistentFlags().Bool("skip-preview", false, "Do not perform a preview before performing the refresh")
	importCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	importCmd.PersistentFlags().Bool("suppress-outputs", false, "Suppress display of stack outputs (in case they contain sensitive values)")
	importCmd.PersistentFlags().String("suppress-permalink", "", "Suppress display of the state permalink")
	importCmd.PersistentFlags().BoolP("yes", "y", false, "Automatically approve and perform the refresh after previewing it")
	importCmd.Flag("suppress-permalink").NoOptDefVal = "false"
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"file":        carapace.ActionFiles(),
		"out":         carapace.ActionFiles(),
		"parent": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 1:
				return action.ActionUrns(importCmd)
			default:
				return carapace.ActionValues()
			}
		}),
		"provider": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 1:
				return action.ActionUrns(importCmd)
			default:
				return carapace.ActionValues()
			}
		}),
		"stack": action.ActionStacks(importCmd, action.StackOpts{}),
	})
}
