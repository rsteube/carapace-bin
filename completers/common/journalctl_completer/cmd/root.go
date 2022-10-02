package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/journalctl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/journalctl"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "journalctl",
	Short: "Query the journal",
	Long:  "https://man7.org/linux/man-pages/man1/journalctl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("after-cursor", "", "Show entries after the specified cursor")
	rootCmd.Flags().BoolP("all", "a", false, "Show all fields, including long and unprintable")
	rootCmd.Flags().StringP("boot", "b", "", "Show current boot or the specified boot")
	rootCmd.Flags().String("case-sensitive", "", "Force case sensitive or insensitive matching")
	rootCmd.Flags().BoolP("catalog", "x", false, "Add message explanations where available")
	rootCmd.Flags().StringP("cursor", "c", "", "Show entries starting at the specified cursor")
	rootCmd.Flags().String("cursor-file", "", "Show entries after cursor in FILE and update FILE")
	rootCmd.Flags().StringP("directory", "D", "", "Show journal files from directory")
	rootCmd.Flags().Bool("disk-usage", false, "Show total disk usage of all journal files")
	rootCmd.Flags().BoolP("dmesg", "k", false, "Show kernel message log from the current boot")
	rootCmd.Flags().Bool("dump-catalog", false, "Show entries in the message catalog")
	rootCmd.Flags().String("facility", "", "Show entries with the specified facilities")
	rootCmd.Flags().StringP("field", "F", "", "List all values that a specified field takes")
	rootCmd.Flags().BoolP("fields", "N", false, "List all field names currently used")
	rootCmd.Flags().String("file", "", "Show journal file")
	rootCmd.Flags().Bool("flush", false, "Flush all journal data from /run into /var")
	rootCmd.Flags().BoolP("follow", "f", false, "Follow the journal")
	rootCmd.Flags().Bool("force", false, "Override of the FSS key pair with --setup-keys")
	rootCmd.Flags().StringP("grep", "g", "", "Show entries with MESSAGE matching PATTERN")
	rootCmd.Flags().Bool("header", false, "Show journal header information")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help text")
	rootCmd.Flags().StringArrayP("identifier", "t", []string{}, "Show entries with the specified syslog identifier")
	rootCmd.Flags().String("image", "", "Operate on files in filesystem image")
	rootCmd.Flags().String("interval", "", "Time interval for changing the FSS sealing key")
	rootCmd.Flags().StringP("lines", "n", "", "Number of journal entries to show")
	rootCmd.Flags().Bool("list-boots", false, "Show terse information about recorded boots")
	rootCmd.Flags().Bool("list-catalog", false, "Show all message IDs in the catalog")
	rootCmd.Flags().StringP("machine", "M", "", "Operate on local container")
	rootCmd.Flags().BoolP("merge", "m", false, "Show entries from all available journals")
	rootCmd.Flags().String("namespace", "", "Show journal data from specified namespace")
	rootCmd.Flags().Bool("no-full", false, "Ellipsize fields")
	rootCmd.Flags().Bool("no-hostname", false, "Suppress output of hostname field")
	rootCmd.Flags().Bool("no-pager", false, "Do not pipe output into a pager")
	rootCmd.Flags().Bool("no-tail", false, "Show all lines, even in follow mode")
	rootCmd.Flags().StringP("output", "o", "", "Change journal output mode")
	rootCmd.Flags().String("output-fields", "", "Select fields to print in verbose/export/json modes")
	rootCmd.Flags().BoolP("pager-end", "e", false, "Immediately jump to the end in the pager")
	rootCmd.Flags().StringP("priority", "p", "", "Show entries with the specified priority")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not show info messages and privilege warning")
	rootCmd.Flags().Bool("relinquish-var", false, "Stop logging to disk, log to temporary file system")
	rootCmd.Flags().BoolP("reverse", "r", false, "Show the newest entries first")
	rootCmd.Flags().String("root", "", "Operate on files below a root directory")
	rootCmd.Flags().Bool("rotate", false, "Request immediate rotation of the journal files")
	rootCmd.Flags().Bool("setup-keys", false, "Generate a new FSS key pair")
	rootCmd.Flags().Bool("show-cursor", false, "Print the cursor after all the entries")
	rootCmd.Flags().StringP("since", "S", "", "Show entries not older than the specified date")
	rootCmd.Flags().Bool("smart-relinquish-var", false, "Similar, but NOP if log directory is on root mount")
	rootCmd.Flags().Bool("sync", false, "Synchronize unwritten journal messages to disk")
	rootCmd.Flags().Bool("system", false, "Show the system journal")
	rootCmd.Flags().StringP("unit", "u", "", "Show logs from the specified unit")
	rootCmd.Flags().StringP("until", "U", "", "Show entries not newer than the specified date")
	rootCmd.Flags().Bool("update-catalog", false, "Update the message catalog database")
	rootCmd.Flags().Bool("user", false, "Show the user journal for the current user")
	rootCmd.Flags().String("user-unit", "", "Show logs from the specified user unit")
	rootCmd.Flags().Bool("utc", false, "Express time in Coordinated Universal Time (UTC)")
	rootCmd.Flags().String("vacuum-files", "", "Leave only the specified number of journal files")
	rootCmd.Flags().String("vacuum-size", "", "Reduce disk usage below specified size")
	rootCmd.Flags().String("vacuum-time", "", "Remove journal files older than specified time")
	rootCmd.Flags().Bool("verify", false, "Verify journal file consistency")
	rootCmd.Flags().String("verify-key", "", "Specify FSS verification key")
	rootCmd.Flags().Bool("version", false, "Show package version")

	rootCmd.Flag("case-sensitive").NoOptDefVal = " "

	// TODO identifier, namespace, machine, unit
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"boot":           action.ActionBoots(),
		"case-sensitive": carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
		"cursor-file":    carapace.ActionFiles(),
		"directory":      carapace.ActionDirectories(),
		"facility": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionFacilities().Invoke(c).Filter(c.Parts).ToA()
		}),
		"field":  action.ActionFields(),
		"file":   carapace.ActionFiles(),
		"image":  carapace.ActionFiles(),
		"output": journalctl.ActionOutputs(),
		"output-fields": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionFields().Invoke(c).Filter(c.Parts).ToA()
		}),
		"priority": carapace.ActionMultiParts("..", func(c carapace.Context) carapace.Action {
			if len(c.Parts) < 2 {
				return action.ActionPriorities()
			}
			return carapace.ActionValues()
		}),
		"root": carapace.ActionDirectories(),
		"since": carapace.Batch(
			carapace.ActionValues("yesterday", "today", "tomorrow"),
			time.ActionDateTime(time.DateTimeOpts{}),
		).ToA(),
		"until": carapace.Batch(
			carapace.ActionValues("yesterday", "today", "tomorrow"),
			time.ActionDateTime(time.DateTimeOpts{}),
		).ToA(),
	})
}
