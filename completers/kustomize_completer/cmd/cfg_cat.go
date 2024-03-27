package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_catCmd = &cobra.Command{
	Use:   "cat",
	Short: "[Alpha] Print Resource Config from a local directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_catCmd).Standalone()
	cfg_catCmd.Flags().Bool("annotate", false, "annotate resources with their file origins.")
	cfg_catCmd.Flags().String("dest", "", "if specified, write output to a file rather than stdout")
	cfg_catCmd.Flags().Bool("exclude-non-local", false, "if true, exclude non-local-config in the output.")
	cfg_catCmd.Flags().Bool("format", true, "format resource config yaml before printing.")
	cfg_catCmd.Flags().String("function-config", "", "path to function config to put in ResourceList -- only if wrapped in a ResourceList.")
	cfg_catCmd.Flags().Bool("include-local", false, "if true, include local-config in the output.")
	cfg_catCmd.Flags().BoolP("recurse-subpackages", "R", true, "print resources recursively in all the nested subpackages")
	cfg_catCmd.Flags().Bool("strip-comments", false, "remove comments from yaml.")
	cfg_catCmd.Flags().StringSlice("style", []string{}, "yaml styles to apply.  may be 'TaggedStyle', 'DoubleQuotedStyle', 'LiteralStyle', 'FoldedStyle', 'FlowStyle'.")
	cfg_catCmd.Flags().String("wrap-kind", "", "if set, wrap the output in this list type kind.")
	cfg_catCmd.Flags().String("wrap-version", "", "if set, wrap the output in this list type apiVersion.")
	cfgCmd.AddCommand(cfg_catCmd)

	carapace.Gen(cfg_catCmd).FlagCompletion(carapace.ActionMap{
		"dest":            carapace.ActionFiles(),
		"function-config": carapace.ActionFiles(),
		"style": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("TaggedStyle", "DoubleQuotedStyle", "LiteralStyle", "FoldedStyle", "FlowStyle").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(cfg_catCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
