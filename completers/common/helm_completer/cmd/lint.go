package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "examine a chart for possible issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	lintCmd.Flags().StringArray("set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	lintCmd.Flags().StringArray("set-file", []string{}, "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)")
	lintCmd.Flags().StringArray("set-string", []string{}, "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	lintCmd.Flags().Bool("strict", false, "fail on lint warnings")
	lintCmd.Flags().StringSliceP("values", "f", []string{}, "specify values in a YAML file or a URL (can specify multiple)")
	lintCmd.Flags().Bool("with-subcharts", false, "lint dependent charts")
	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).FlagCompletion(carapace.ActionMap{
		"set-file": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 1:
					return carapace.ActionFiles()
				default:
					return carapace.ActionValues()
				}
			})
		}),
	})

	carapace.Gen(lintCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
