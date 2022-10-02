package cmd

import (
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/third_party/github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tail",
	Short: "output the last part of files",
	Long:  "https://linux.die.net/man/1/tail",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("F", "F", false, "same as --follow=name --retry")
	rootCmd.Flags().StringP("bytes", "c", "", "output the last NUM bytes; or use -c +NUM to")
	rootCmd.Flags().StringP("follow", "f", "", "")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("lines", "n", "", "output the last NUM lines, instead of the last 10;")
	rootCmd.Flags().String("max-unchanged-stats", "", "")
	rootCmd.Flags().String("pid", "", "with -f, terminate after process ID, PID dies")
	rootCmd.Flags().StringP("quiet", "q", "", "never output headers giving file names")
	rootCmd.Flags().Bool("retry", false, "keep trying to open a file if it is inaccessible")
	rootCmd.Flags().StringP("sleep-interval", "s", "", "with -f, sleep for approximately N seconds")
	rootCmd.Flags().BoolP("verbose", "v", false, "always output headers giving file names")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	rootCmd.Flag("follow").NoOptDefVal = "descriptor"

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"follow": carapace.ActionValues("descriptor", "name"),
		"pid":    ActionPids(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}

func ActionPids() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if processes, err := ps.Processes(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			pids := make([]string, 0)
			for _, process := range processes {
				pids = append(pids, strconv.Itoa(process.Pid()), process.Executable())
			}
			return carapace.ActionValuesDescribed(pids...)
		}
	})
}
