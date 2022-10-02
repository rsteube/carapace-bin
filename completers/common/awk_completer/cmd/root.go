package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "awk",
	Short: "pattern scanning and processing language",
	Long:  "https://en.wikipedia.org/wiki/AWK",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("assign", "v", "", "")
	rootCmd.Flags().BoolP("bignum", "M", false, "")
	rootCmd.Flags().BoolP("characters-as-bytes", "b", false, "")
	rootCmd.Flags().BoolP("copyright", "C", false, "")
	rootCmd.Flags().StringP("debug", "D", "", "")
	rootCmd.Flags().StringP("dump-variables", "d", "", "")
	rootCmd.Flags().StringP("exec", "E", "", "")
	rootCmd.Flags().StringP("field-separator", "F", "", "")
	rootCmd.Flags().StringP("file", "f", "", "")
	rootCmd.Flags().BoolP("gen-pot", "g", false, "")
	rootCmd.Flags().BoolP("help", "h", false, "")
	rootCmd.Flags().StringP("include", "i", "", "")
	rootCmd.Flags().StringP("lint", "L", "", "")
	rootCmd.Flags().BoolP("lint-old", "t", false, "")
	rootCmd.Flags().StringP("load", "l", "", "")
	rootCmd.Flags().BoolP("no-optimize", "s", false, "")
	rootCmd.Flags().BoolP("non-decimal-data", "n", false, "")
	rootCmd.Flags().BoolP("optimize", "O", false, "")
	rootCmd.Flags().BoolP("posix", "P", false, "")
	rootCmd.Flags().StringP("pretty-print", "o", "", "")
	rootCmd.Flags().StringP("profile", "p", "", "")
	rootCmd.Flags().BoolP("re-interval", "r", false, "")
	rootCmd.Flags().BoolP("sandbox", "S", false, "")
	rootCmd.Flags().StringP("source", "e", "", "")
	rootCmd.Flags().BoolP("traditional", "c", false, "")
	rootCmd.Flags().BoolP("use-lc-numeric", "N", false, "")
	rootCmd.Flags().BoolP("version", "V", false, "")

	rootCmd.Flag("dump-variables").NoOptDefVal = " "
	rootCmd.Flag("debug").NoOptDefVal = " "
	rootCmd.Flag("lint").NoOptDefVal = " "
	rootCmd.Flag("pretty-print").NoOptDefVal = " "
	rootCmd.Flag("profile").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"debug":          carapace.ActionFiles(),
		"dump-variables": carapace.ActionFiles(),
		"exec":           carapace.ActionFiles(),
		"file":           carapace.ActionFiles(),
		"include":        carapace.ActionFiles(),
		"lint":           carapace.ActionValues("fatal", "invalid", "no-ext"),
		"pretty-print":   carapace.ActionFiles(),
		"profile":        carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("file").Changed {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
