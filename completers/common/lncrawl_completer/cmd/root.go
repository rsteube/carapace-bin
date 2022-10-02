package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lncrawl",
	Short: "Generate and download e-books from online sources",
	Long:  "https://github.com/dipu-bd/lightnovel-crawler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("add-source-url", false, "Add source url at the end of each chapter.")
	rootCmd.Flags().Bool("all", false, "Download all chapters.")
	rootCmd.Flags().String("bot", "", "Select a bot.")
	rootCmd.Flags().String("chapters", "", "A list of specific chapter urls.")
	rootCmd.Flags().String("crawler", "", "Load additional crawler files.")
	rootCmd.Flags().String("filename", "", "Set the output file name")
	rootCmd.Flags().Bool("filename-only", false, "Skip appending chapter range with file name")
	rootCmd.Flags().String("first", "", "Download first few chapters.")
	rootCmd.Flags().BoolP("force", "f", false, "Force replace any existing folder.")
	rootCmd.Flags().String("format", "", "Define which formats to output. Default: all.")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().BoolP("ignore", "i", false, "Ignore any existing folder (do not replace).")
	rootCmd.Flags().BoolS("l", "l", false, "Set log levels.")
	rootCmd.Flags().String("last", "", "Download last few chapters.")
	rootCmd.Flags().Bool("list-sources", false, "Display a list of available sources.")
	rootCmd.Flags().String("login", "", "User name/email address and password for login.")
	rootCmd.Flags().Bool("multi", false, "Build separate books by volumes.")
	rootCmd.Flags().StringP("output", "o", "", "Path where the downloads to be stored.")
	rootCmd.Flags().String("page", "", "The start and final chapter urls.")
	rootCmd.Flags().StringP("query", "q", "", "Novel query followed by list of source sites.")
	rootCmd.Flags().String("range", "", "The start and final chapter indexes.")
	rootCmd.Flags().String("resume", "", "Resume download of a novel")
	rootCmd.Flags().String("shard-count", "", "Discord bot shard counts.")
	rootCmd.Flags().String("shard-id", "", "Discord bot shard id.")
	rootCmd.Flags().Bool("single", false, "Put everything in a single book.")
	rootCmd.Flags().StringP("source", "s", "", "Profile page url of the novel.")
	rootCmd.Flags().StringP("sources", "x", "", "Filter out the sources to search for novels.")
	rootCmd.Flags().Bool("suppress", false, "Suppress all input prompts and use defaults.")
	rootCmd.Flags().BoolP("version", "v", false, "show program's version number and exit")
	rootCmd.Flags().String("volumes", "", "The list of volume numbers to download.")

	// TODO incomplete and currently does not support flags consuming multiple arguments (argparse)
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"crawler":  carapace.ActionFiles(),
		"filename": carapace.ActionFiles(),
		"format": carapace.ActionValues("json", "epub", "text", "web", "docx", "mobi", "pdf", "rtf", "txt", "azw3", "fb2", "lit", "lrf", "oeb", "pdb", "rb", "snb", "tcr").StyleF(func(s string) string {
			return style.ForPathExt("." + s)
		}),
		"output": carapace.ActionDirectories(),
	})
}
