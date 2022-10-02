package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clamdscan",
	Short: "scan files and directories for viruses using Clam AntiVirus Daemon",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("allmatch", "z", false, "Continue scanning within file after finding a match.")
	rootCmd.Flags().String("config-file", "", "Read configuration from FILE.")
	rootCmd.Flags().String("copy", "", "Copy infected files into DIRECTORY")
	rootCmd.Flags().Bool("fdpass", false, "Pass filedescriptor to clamd (useful if clamd is running as a different user)")
	rootCmd.Flags().StringP("file-list", "f", "", "Scan files from FILE")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().BoolP("infected", "i", false, "Only print infected files")
	rootCmd.Flags().StringP("log", "l", "", "Save scan report in FILE")
	rootCmd.Flags().String("move", "", "Move infected files into DIRECTORY")
	rootCmd.Flags().BoolP("multiscan", "m", false, "Force MULTISCAN mode")
	rootCmd.Flags().Bool("no-summary", false, "Disable summary at end of scanning")
	rootCmd.Flags().StringP("ping", "p", "", "Ping clamd up to [A] times at optional interval [I] until it responds.")
	rootCmd.Flags().Bool("quiet", false, "Be quiet, only output error messages")
	rootCmd.Flags().Bool("reload", false, "Request clamd to reload virus database")
	rootCmd.Flags().Bool("remove", false, "Remove infected files. Be careful!")
	rootCmd.Flags().Bool("stdout", false, "Write to stdout instead of stderr. Does not affect 'debug' messages.")
	rootCmd.Flags().Bool("stream", false, "Force streaming files to clamd (for debugging and unit testing)")
	rootCmd.Flags().BoolP("verbose", "v", false, "Be verbose")
	rootCmd.Flags().BoolP("version", "V", false, "Print version number and exit")
	rootCmd.Flags().BoolP("wait", "w", false, "Wait up to 30 seconds for clamd to start. Optionally use alongside --ping to set attempts [A] and interval [I] to check clamd.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"copy":        carapace.ActionDirectories(),
		"file-list":   carapace.ActionFiles(),
		"log":         carapace.ActionFiles(),
		"move":        carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
