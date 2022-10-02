package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/dmesg_completer/cmd/action"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dmesg",
	Short: "Display or control the kernel ring buffer",
	Long:  "https://linux.die.net/man/8/dmesg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("buffer-size", "s", "", "buffer size to query the kernel ring buffer")
	rootCmd.Flags().BoolP("clear", "C", false, "clear the kernel ring buffer")
	rootCmd.Flags().StringP("color", "L", "", "colorize messages(auto, always or never)")
	rootCmd.Flags().StringP("console-level", "n", "", "set level of messages printed to console")
	rootCmd.Flags().BoolP("console-off", "D", false, "disable printing messages to console")
	rootCmd.Flags().BoolP("console-on", "E", false, "enable printing messages to console")
	rootCmd.Flags().BoolP("ctime", "T", false, "show human-readable timestamp (may be inaccurate!)")
	rootCmd.Flags().BoolP("decode", "x", false, "decode facility and level to readable string")
	rootCmd.Flags().StringP("facility", "f", "", "restrict output to defined facilities")
	rootCmd.Flags().StringP("file", "F", "", "use the file instead of the kernel log buffer")
	rootCmd.Flags().BoolP("follow", "w", false, "wait for new messages")
	rootCmd.Flags().BoolP("follow-new", "W", false, "wait and print only new messages")
	rootCmd.Flags().BoolP("force-prefix", "p", false, "force timestamp output on each line of multi-line messages")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("human", "H", false, "human readable output")
	rootCmd.Flags().BoolP("kernel", "k", false, "display kernel messages")
	rootCmd.Flags().StringP("level", "l", "", "restrict output to defined levels")
	rootCmd.Flags().Bool("noescape", false, "don't escape unprintable character")
	rootCmd.Flags().BoolP("nopager", "P", false, "do not pipe output into a pager")
	rootCmd.Flags().BoolP("notime", "t", false, "don't show any timestamp with messages")
	rootCmd.Flags().BoolP("raw", "r", false, "print the raw message buffer")
	rootCmd.Flags().BoolP("read-clear", "c", false, "read and clear all messages")
	rootCmd.Flags().BoolP("reltime", "e", false, "show local time and time delta in readable format")
	rootCmd.Flags().BoolP("show-delta", "d", false, "show time delta between printed messages")
	rootCmd.Flags().String("since", "", "display the lines since the specified time")
	rootCmd.Flags().BoolP("syslog", "S", false, "force to use syslog(2) rather than /dev/kmsg")
	rootCmd.Flags().String("time-format", "", "show timestamp using the given format:")
	rootCmd.Flags().String("until", "", "display the lines until the specified time")
	rootCmd.Flags().BoolP("userspace", "u", false, "display userspace messages")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("color").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":         carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"console-level": action.ActionLogLevels(),
		"facility": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionFacilities().Invoke(c).Filter(c.Parts).ToA()
		}),
		"file": carapace.ActionFiles(),
		"level": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLogLevels().Invoke(c).Filter(c.Parts).ToA()
		}),
		"time-format": carapace.ActionValues("delta", "reltime", "ctime", "notime", "iso"),
	})
}
