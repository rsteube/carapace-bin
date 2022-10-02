package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "i3",
	Short: "an improved dynamic, tiling window manager",
	Long:  "https://i3wm.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "validate configuration file and exit")
	rootCmd.Flags().StringS("L", "L", "", "path to the serialized layout during restarts")
	rootCmd.Flags().BoolS("V", "V", false, "enable verbose mode")
	rootCmd.Flags().BoolS("a", "a", false, "disable autostart ('exec' lines in config)")
	rootCmd.Flags().StringS("c", "c", "", "use the provided configfile instead")
	rootCmd.Flags().StringS("d", "d", "", "enable debug output")
	rootCmd.Flags().Bool("force-xinerama", false, "Use Xinerama instead of RandR.")
	rootCmd.Flags().Bool("get-socketpath", false, "Retrieve the i3 IPC socket path from X11, print it, then exit.")
	rootCmd.Flags().String("shmlog-size", "", "Limits the size of the i3 SHM log to <limit> bytes.")
	rootCmd.Flags().BoolS("v", "v", false, "display version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"L": carapace.ActionFiles(),
		"c": carapace.ActionFiles(),
		"d": carapace.ActionValues("all"),
	})
}
