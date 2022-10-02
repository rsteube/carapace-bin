package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "date",
	Short: "print or set the system date and time",
	Long:  "https://linux.die.net/man/1/date",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("date", "d", "", "display time described by STRING, not 'now'")
	rootCmd.Flags().Bool("debug", false, "annotate the parsed date,")
	rootCmd.Flags().StringP("file", "f", "", "like --date; once for each line of DATEFILE")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("iso-8601", "I", "", "output date/time in ISO 8601 format")
	rootCmd.Flags().StringP("reference", "r", "", "display the last modification time of FILE")
	rootCmd.Flags().String("rfc-3339", "", "output date/time in RFC 3339 format.")
	rootCmd.Flags().BoolP("rfc-email", "R", false, "output date and time in RFC 5322 format.")
	rootCmd.Flags().StringP("set", "s", "", "set time described by STRING")
	rootCmd.Flags().String("universal", "", "print or set Coordinated Universal Time (UTC)")
	rootCmd.Flags().StringP("utc,", "u", "", "print or set Coordinated Universal Time (UTC)")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"date":      time.ActionDateTime(time.DateTimeOpts{}),
		"file":      carapace.ActionFiles(),
		"reference": carapace.ActionFiles(),
	})
}
