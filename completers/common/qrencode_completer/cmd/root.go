package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qrencode",
	Short: "Encode input data in a QR Code and save as a PNG or EPS image",
	Long:  "https://linux.die.net/man/1/qrencode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("8bit", "8", false, "encode entire data in 8-bit mode. -k, -c and -i will be ignored.")
	rootCmd.Flags().String("background", "", "specify foreground/background color in hexadecimal notation.")
	rootCmd.Flags().BoolP("casesensitive", "c", false, "encode lower-case alphabet characters in 8-bit mode.")
	rootCmd.Flags().StringP("dpi", "d", "", "specify the DPI of the generated PNG.")
	rootCmd.Flags().String("foreground", "", "specify foreground/background color in hexadecimal notation.")
	rootCmd.Flags().BoolP("help", "h", false, "display the help message.")
	rootCmd.Flags().BoolP("ignorecase", "i", false, "ignore case distinctions and use only upper-case characters.")
	rootCmd.Flags().Bool("inline", false, "only useful for SVG output, generates an SVG without the XML tag.")
	rootCmd.Flags().BoolP("kanji", "k", false, "assume that the input text contains kanji (shift-jis).")
	rootCmd.Flags().StringP("level", "l", "", "specify error correction level from L (lowest) to H (highest).")
	rootCmd.Flags().StringP("margin", "m", "", "specify the width of the margins.")
	rootCmd.Flags().BoolP("micro", "M", false, "encode in a Micro QR Code.")
	rootCmd.Flags().StringP("output", "o", "", "write image to FILENAME.")
	rootCmd.Flags().StringP("read-from", "r", "", "read input data from FILENAME.")
	rootCmd.Flags().Bool("rle", false, "enable run-length encoding for SVG.")
	rootCmd.Flags().StringP("size", "s", "", "specify module size in dots (pixels).")
	rootCmd.Flags().Bool("strict-version", false, "disable automatic version number adjustment.")
	rootCmd.Flags().BoolP("structured", "S", false, "make structured symbols.")
	rootCmd.Flags().Bool("svg-path", false, "use single path to draw modules for SVG.")
	rootCmd.Flags().StringP("symversion", "v", "", "specify the minimum version of the symbol.")
	rootCmd.Flags().StringP("type", "t", "", "specify the type of the generated image.")
	rootCmd.Flags().Bool("verbose", false, "display verbose information to stderr.")
	rootCmd.Flags().BoolP("version", "V", false, "display the version number and copyrights of the qrencode.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"level":     carapace.ActionValues("L", "M", "Q", "H"),
		"output":    carapace.ActionFiles(),
		"read-from": carapace.ActionFiles(),
		"type":      carapace.ActionValues("PNG", "PNG32", "EPS", "SVG", "XPM", "ANSI", "ANSI256", "ASCII", "ASCIIi", "UTF8", "UTF8i", "ANSIUTF8", "ANSIUTF8i", "ANSI256UTF8"),
	})
}
