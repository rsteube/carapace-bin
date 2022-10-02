package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tesseract",
	Short: "command-line OCR engine",
	Long:  "https://github.com/tesseract-ocr/tessdoc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("c", "c", "", "Set value for parameter CONFIGVAR to VALUE.")
	rootCmd.Flags().String("dpi", "", "Specify the resolution N in DPI for the input image(s).")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message.")
	rootCmd.Flags().Bool("help-extra", false, "Show extra help for advanced users.")
	rootCmd.Flags().Bool("help-oem", false, "Show OCR Engine modes.")
	rootCmd.Flags().Bool("help-psm", false, "Show page segmentation modes.")
	rootCmd.Flags().StringS("l", "l", "", "The language or script to use.")
	rootCmd.Flags().Bool("list-langs", false, "List available languages for tesseract engine.")
	rootCmd.Flags().String("oem", "", "Specify OCR Engine mode.")
	rootCmd.Flags().Bool("print-parameters", false, "Print tesseract parameters.")
	rootCmd.Flags().String("psm", "", "Set Tesseract to only run a subset of layout analysis and assume a certain form of image.")
	rootCmd.Flags().String("tessdata-dir", "", "Specify the location of tessdata path.")
	rootCmd.Flags().String("user-patterns", "", "Specify the location of user patterns file.")
	rootCmd.Flags().String("user-words", "", "Specify the location of user words file.")
	rootCmd.Flags().BoolP("version", "v", false, "Returns the current version of the tesseract(1) executable.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"l": carapace.ActionMultiParts("+", func(c carapace.Context) carapace.Action {
			return ActionLanguages().Invoke(c).Filter(c.Parts).ToA()
		}),
		"oem":           ActionOcrEngineModes(),
		"psm":           ActionAnalysisSubsets(),
		"tessdata-dir":  carapace.ActionDirectories(),
		"user-patterns": carapace.ActionFiles(),
		"user-words":    carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}

func ActionLanguages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("tesseract", "--list-langs")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[1 : len(lines)-1]...)
		})
	})
}

func ActionAnalysisSubsets() carapace.Action {
	return carapace.ActionValuesDescribed(
		"0", "Orientation and script detection (OSD) only.",
		"1", "Automatic page segmentation with OSD.",
		"2", "Automatic page segmentation, but no OSD, or OCR. (not implemented)",
		"3", "Fully automatic page segmentation, but no OSD. (Default)",
		"4", "Assume a single column of text of variable sizes.",
		"5", "Assume a single uniform block of vertically aligned text.",
		"6", "Assume a single uniform block of text.",
		"7", "Treat the image as a single text line.",
		"8", "Treat the image as a single word.",
		"9", "Treat the image as a single word in a circle.",
		"10", "Treat the image as a single character.",
		"11", "Sparse text. Find as much text as possible in no particular order.",
		"12", "Sparse text with OSD.",
		"13", "Raw line.",
	)
}

func ActionOcrEngineModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"0", "Original Tesseract only.",
		"1", "Neural nets LSTM only.",
		"2", "Tesseract + LSTM.",
		"3", "Default, based on what is available.",
	)
}
