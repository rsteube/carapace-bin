package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "upx",
	Short: "compress or expand executable files",
	Long:  "https://upx.github.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("1", "1", false, "compress faster")
	rootCmd.Flags().BoolS("2", "2", false, "")
	rootCmd.Flags().BoolS("3", "3", false, "")
	rootCmd.Flags().BoolS("4", "4", false, "")
	rootCmd.Flags().BoolS("5", "5", false, "")
	rootCmd.Flags().BoolS("6", "6", false, "")
	rootCmd.Flags().BoolS("7", "7", false, "")
	rootCmd.Flags().BoolS("8", "8", false, "")
	rootCmd.Flags().Bool("8-bit", false, "uses 8 bit size compression [default: 32 bit]")
	rootCmd.Flags().Bool("8086", false, "make compressed exe/sys/com work on any 8086")
	rootCmd.Flags().Bool("8mib-ram", false, "8 megabyte memory limit [default: 2 MiB]")
	rootCmd.Flags().BoolS("9", "9", false, "compress better")
	rootCmd.Flags().BoolS("L", "L", false, "display software license")
	rootCmd.Flags().BoolS("V", "V", false, "display version number")
	rootCmd.Flags().BoolP("backup", "k", false, "keep backup files")
	rootCmd.Flags().Bool("boot-only", false, "disables client/host transfer compatibility")
	rootCmd.Flags().Bool("brute", false, "try all available compression methods & filters [slow]")
	rootCmd.Flags().Bool("coff", false, "produce COFF output")
	rootCmd.Flags().Bool("color", false, "change look")
	rootCmd.Flags().String("compress-exports", "", "")
	rootCmd.Flags().String("compress-icons", "", "")
	rootCmd.Flags().String("compress-resources", "", "")
	rootCmd.Flags().BoolS("d", "d", false, "decompress")
	rootCmd.Flags().BoolS("f", "f", false, "force compression of suspicious files")
	rootCmd.Flags().BoolS("h", "h", false, "give this help")
	rootCmd.Flags().String("keep-resource", "", "do not compress resources specified by list")
	rootCmd.Flags().BoolS("l", "l", false, "list compressed file")
	rootCmd.Flags().Bool("le", false, "produce LE output [default: EXE]")
	rootCmd.Flags().Bool("mono", false, "change look")
	rootCmd.Flags().Bool("no-align", false, "don't align to 2048 bytes [enables: --console-run]")
	rootCmd.Flags().Bool("no-backup", false, "no backup files [default]")
	rootCmd.Flags().Bool("no-color", false, "change look")
	rootCmd.Flags().Bool("no-progress", false, "change look")
	rootCmd.Flags().Bool("no-reloc", false, "put no relocations in to the exe header")
	rootCmd.Flags().StringS("o", "o", "", "write output to file")
	rootCmd.Flags().String("overlay", "", "")
	rootCmd.Flags().Bool("preserve-build-id", false, "copy .gnu.note.build-id to compressed output")
	rootCmd.Flags().BoolS("q", "q", false, "be quiet")
	rootCmd.Flags().String("strip-relocs", "", "")
	rootCmd.Flags().BoolS("t", "t", false, "test compressed file")
	rootCmd.Flags().Bool("ultra-brute", false, "try even more compression variants [very slow]")
	rootCmd.Flags().BoolS("v", "v", false, "be verbose")

	rootCmd.Flag("overlay").NoOptDefVal = "copy"
	rootCmd.Flag("compress-exports").NoOptDefVal = "1"
	rootCmd.Flag("compress-icons").NoOptDefVal = "2"
	rootCmd.Flag("compress-resources").NoOptDefVal = "0"
	rootCmd.Flag("keep-resource").NoOptDefVal = " "
	rootCmd.Flag("strip-relocs").NoOptDefVal = "1"

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"compress-exports": carapace.ActionValuesDescribed(
			"0", "do not compress the export section",
			"1", "compress the export section [default]",
		),
		"compress-icons": carapace.ActionValuesDescribed(
			"0", "do not compress any icons",
			"1", "compress all but the first icon",
			"2", "compress all but the first icon directory [default]",
			"3", "compress all icons",
		),
		"compress-resources": carapace.ActionValuesDescribed(
			"0", "do not compress any resources at all",
		),
		"o": carapace.ActionFiles(),
		"overlay": carapace.ActionValuesDescribed(
			"copy", "copy any extra data attached to the file [default]",
			"strip", "strip any extra data attached to the file [DANGEROUS]",
			"skip", "don't compress a file with an overlay",
		),
		"strip-relocs": carapace.ActionValuesDescribed(
			"0", "do not strip relocations",
			"1", "strip relocations [default]",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
