package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clamonacc",
	Short: "an anti-virus on-access scanning daemon and clamd client",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("allmatch", "z", false, "Continue scanning within file after finding a match.")
	rootCmd.Flags().StringP("config-file", "c", "", "Read configuration from FILE")
	rootCmd.Flags().String("copy", "", "Copy infected files into DIRECTORY")
	rootCmd.Flags().StringP("exclude-list", "e", "", "Exclude directories from FILE")
	rootCmd.Flags().Bool("fdpass", false, "Pass filedescriptor to clamd (useful if clamd is running as a different user)")
	rootCmd.Flags().BoolP("foreground", "F", false, "Output to foreground and do not daemonize")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().StringP("log", "l", "", "Save scanning output to FILE")
	rootCmd.Flags().String("move", "", "Move infected files into DIRECTORY")
	rootCmd.Flags().BoolP("ping", "p", false, "]    Ping clamd up to [A] times at optional interval [I] until it responds.")
	rootCmd.Flags().Bool("remove", false, "Remove infected files. Be careful!")
	rootCmd.Flags().Bool("stream", false, "Force streaming files to clamd (for debugging and unit testing)")
	rootCmd.Flags().BoolP("verbose", "v", false, "Be verbose")
	rootCmd.Flags().BoolP("version", "V", false, "Print version number and exit")
	rootCmd.Flags().BoolP("wait", "w", false, "Wait up to 30 seconds for clamd to start. Optionally use alongside --ping to set attempts [A] and interval [I] to check clamd.")
	rootCmd.Flags().StringP("watch-list", "W", "", "Watch directories from FILE")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file":  carapace.ActionFiles(),
		"copy":         carapace.ActionDirectories(),
		"exclude-list": carapace.ActionFiles(),
		"log":          carapace.ActionDirectories(),
		"move":         carapace.ActionDirectories(),
		"watch-list":   carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
