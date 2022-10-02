package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "acpid",
	Short: "Advanced Configuration and Power Interface event daemon",
	Long:  "https://linux.die.net/man/8/acpid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("confdir", "c", "", "This option changes the directory in which acpid looks for rule configuration files.")
	rootCmd.Flags().StringP("dropaction", "r", "", "This  option  defines the pseudo-action which tells acpid to abort all processing of an event, including client notifications.")
	rootCmd.Flags().StringP("eventfile", "e", "", "This option changes the event file  from  which  acpid  reads  events.")
	rootCmd.Flags().BoolP("foreground", "f", false, "This option keeps acpid in the foreground by not forking at startup.")
	rootCmd.Flags().BoolP("help", "h", false, "Show help and exit.")
	rootCmd.Flags().StringP("lockfile", "L", "", "This  option  changes  the  lock  file  used  to  stop event processing.")
	rootCmd.Flags().BoolP("logevents", "l", false, "This option tells acpid to log information about all events and actions.")
	rootCmd.Flags().BoolP("netlink", "n", false, "This  option  forces acpid to use the Linux kernel input layer and netlink interface for ACPI events.")
	rootCmd.Flags().BoolP("nosocket", "S", false, "This option tells acpid not to open a UNIX domain socket.")
	rootCmd.Flags().StringP("pidfile", "p", "", "This option tells acpid to use the specified file as its pidfile.")
	rootCmd.Flags().StringP("socketfile", "s", "", "This option changes the name of the UNIX domain socket which acpid  opens.")
	rootCmd.Flags().StringP("socketgroup", "g", "", "This option changes the group ownership of the UNIX domain  socket  to  which  acpid publishes events.")
	rootCmd.Flags().StringP("socketmode", "m", "", "This  option  changes  the permissions of the UNIX domain socket to which acpid pub‐ lishes events.")
	rootCmd.Flags().BoolP("tpmutefix", "t", false, "This option enables special handling of the mute button for certain ThinkPad models.")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and exit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"confdir":     carapace.ActionDirectories(),
		"eventfile":   carapace.ActionFiles(),
		"lockfile":    carapace.ActionFiles(),
		"pidfile":     carapace.ActionFiles(),
		"socketfile":  carapace.ActionFiles(),
		"socketgroup": os.ActionGroups(),
		"socketmode":  fs.ActionFileModesNumeric(),
	})
}
