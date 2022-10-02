package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "od",
	Short: "dump files in octal and other formats",
	Long:  "https://linux.die.net/man/1/od",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("address-radix", "A", "", "output format for file offsets; RADIX is one")
	rootCmd.Flags().String("endian", "", "swap input bytes according the specified order")
	rootCmd.Flags().StringP("format", "t", "", "select output format or formats")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("output-duplicates", "v", false, "do not use * to mark line suppression")
	rootCmd.Flags().StringP("read-bytes", "N", "", "limit dump to BYTES input bytes")
	rootCmd.Flags().StringP("skip-bytes", "j", "", "skip BYTES input bytes first")
	rootCmd.Flags().Bool("traditional", false, "accept arguments in third form above")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"address-radix": carapace.ActionValuesDescribed(
			"d", "decimal",
			"n", "none",
			"o", "octal",
			"x", "hexadecimal",
		),
		"endian": carapace.ActionValues("big", "little"),
		// TODO complete format (this is incomplete, compare with zsh completions)
		"format": carapace.ActionValuesDescribed(
			"a", "named character",
			"c", "character (C-style escape)",
			"d", "decimal",
			"f", "floating-point number",
			"o", "octal",
			"u", "unsigned decimal",
			"x", "hexadecimal",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
