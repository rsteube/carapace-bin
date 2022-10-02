package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fzf",
	Short: "a command-line fuzzy finder",
	Long:  "https://github.com/junegunn/fzf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("algo", "", "Fuzzy matching algorithm")
	rootCmd.Flags().Bool("ansi", false, "Enable processing of ANSI color codes")
	rootCmd.Flags().String("bind", "", "Custom key bindings")
	rootCmd.Flags().String("border", "", "Draw border around the finder")
	rootCmd.Flags().String("color", "", "Base scheme")
	rootCmd.Flags().Bool("cycle", false, "Enable cyclic scroll")
	rootCmd.Flags().StringP("delimiter", "d", "", "Field delimiter regex")
	rootCmd.Flags().BoolP("exact", "e", false, "Enable Exact-match")
	rootCmd.Flags().BoolP("exit-0", "0", false, "Exit immediately when there's no match")
	rootCmd.Flags().String("expect", "", "Comma-separated list of keys to complete fzf")
	rootCmd.Flags().BoolP("extended", "x", false, "Extended-search mode")
	rootCmd.Flags().Bool("filepath-word", false, "Make word-wise movements respect path separators")
	rootCmd.Flags().StringP("filter", "f", "", "Filter mode")
	rootCmd.Flags().String("header", "", "String to print as header")
	rootCmd.Flags().String("header-lines", "", "The first N lines of the input are treated as header")
	rootCmd.Flags().String("height", "", "Display fzf window below the cursor with the given height")
	rootCmd.Flags().String("history", "", "History file")
	rootCmd.Flags().String("history-size", "", "Maximum number of history entries")
	rootCmd.Flags().String("hscroll-off", "", "Number of screen columns to keep to the right")
	rootCmd.Flags().BoolS("i", "i", false, "Case-insensitive match")
	rootCmd.Flags().String("info", "", "Finder info style [default|inline|hidden]")
	rootCmd.Flags().String("jump-labels", "", "Label characters for jump and jump-accept")
	rootCmd.Flags().Bool("keep-right", false, "Keep the right end of the line visible on overflow")
	rootCmd.Flags().String("layout", "", "Choose layout")
	rootCmd.Flags().Bool("literal", false, "Do not normalize latin script letters before matching")
	rootCmd.Flags().String("margin", "", "Screen margin")
	rootCmd.Flags().String("marker", "", "Multi-select marker")
	rootCmd.Flags().String("min-height", "", "Minimum height when --height is given in percent")
	rootCmd.Flags().StringP("multi", "m", "", "Enable multi-select with tab/shift-tab")
	rootCmd.Flags().Bool("no-bold", false, "Do not use bold text")
	rootCmd.Flags().Bool("no-hscroll", false, "Disable horizontal scroll")
	rootCmd.Flags().Bool("no-mouse", false, "Disable mouse")
	rootCmd.Flags().StringP("nth", "n", "", "Comma-separated list of field index expressions")
	rootCmd.Flags().String("padding", "", "Padding inside border")
	rootCmd.Flags().Bool("phony", false, "Do not perform search")
	rootCmd.Flags().String("pointer", "", "Pointer to the current line")
	rootCmd.Flags().String("preview", "", "Command to preview highlighted line")
	rootCmd.Flags().String("preview-window", "", "Preview window layout")
	rootCmd.Flags().Bool("print-query", false, "Print query as the first line")
	rootCmd.Flags().Bool("print0", false, "Print output delimited by ASCII NUL characters")
	rootCmd.Flags().String("prompt", "", "Input prompt")
	rootCmd.Flags().StringP("query", "q", "", "Start the finder with the given query")
	rootCmd.Flags().Bool("read0", false, "Read input delimited by ASCII NUL characters")
	rootCmd.Flags().BoolP("select-1", "1", false, "Automatically select the only match")
	rootCmd.Flags().Bool("sync", false, "Synchronous search for multi-staged filtering")
	rootCmd.Flags().String("tabstop", "", "Number of spaces for a tab character")
	rootCmd.Flags().Bool("tac", false, "Reverse the order of the input")
	rootCmd.Flags().String("tiebreak", "", "Comma-separated list of sort criteria to apply")
	rootCmd.Flags().Bool("version", false, "Display version information and exit")
	rootCmd.Flags().String("with-nth", "", "Transform the presentation of each line using field index expressions")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"algo":   carapace.ActionValues("v1", "v2"),
		"border": carapace.ActionValues("rounded", "sharp", "horizontal", "vertical", "top", "bottom", "left", "right"),
		"color":  carapace.ActionValues("dark", "light", "16", "bw"),
		"info":   carapace.ActionValues("default", "inline", "hidden"),
		"layout": carapace.ActionValues("default", "reverse", "reverse-list"),
		"tiebreak": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("length", "begin", "end", "index").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
