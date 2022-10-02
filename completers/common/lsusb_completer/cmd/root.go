package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os/usb"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsusb",
	Short: "list USB devices",
	Long:  "https://linux.die.net/man/8/lsusb",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("D", "D", "", "Selects which device lsusb will examine")
	rootCmd.Flags().StringS("d", "d", "", "Show only devices with the specified vendor and product ID numbers")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage and help")
	rootCmd.Flags().StringS("s", "s", "", "Show only devices with specified device and/or bus numbers")
	rootCmd.Flags().BoolP("tree", "t", false, "Dump the physical USB device hierarchy as a tree")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase verbosity")
	rootCmd.Flags().BoolP("version", "V", false, "Show version of program")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D": carapace.ActionFiles(),
		"d": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return usb.ActionProductNumbers().Invoke(c).ToMultiPartsA(":")
		}),
		"s": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return usb.ActionDeviceNumbers().Invoke(c).ToMultiPartsA(":")
		}),
	})
}
