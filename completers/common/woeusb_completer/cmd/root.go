package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "woeusb",
	Short: "A Linux program to create a Windows USB stick installer",
	Long:  "https://github.com/WoeUSB/WoeUSB-ng",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("about", false, "Show info about this application")
	rootCmd.Flags().Bool("debug", false, "Enable script debugging")
	rootCmd.Flags().BoolP("device", "d", false, "Completely WIPE the entire USB storage device")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().StringP("label", "l", "", "Specify label for the newly created file system")
	rootCmd.Flags().Bool("no-color", false, "Disable message coloring")
	rootCmd.Flags().BoolP("partition", "p", false, "Copy Windows files to an existing partition")
	rootCmd.Flags().String("target-filesystem", "", "Specify the filesystem to use as the target")
	rootCmd.Flags().String("tgt-fs", "", "Specify the filesystem to use as the target")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")
	rootCmd.Flags().BoolP("version", "V", false, "Print application version")
	rootCmd.Flags().Bool("workaround-bios-boot-flag", false, "Workaround BIOS bug that won't include the device in boot menu")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"target-filesystem": carapace.ActionValues("FAT", "NTFS"),
		"tgt-fs":            carapace.ActionValues("FAT", "NTFS"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.Batch(
			fs.ActionBlockDevices(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
