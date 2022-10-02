package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/hwinfo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hwinfo",
	Short: "Probe for hardware",
	Long:  "https://github.com/openSUSE/hwinfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("arch", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("bios", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("block", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("bluetooth", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("braille", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("bridge", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("camera", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("cdrom", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("chipcard", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("cpu", false, "probe for particular hardware item")
	rootCmd.Flags().String("debug", "", "Set debug level to N.")
	rootCmd.Flags().Bool("disk", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("dsl", false, "probe for particular hardware item")
	rootCmd.Flags().String("dump-db", "", "Dump hardware data base.")
	rootCmd.Flags().Bool("dvb", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("fingerprint", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("floppy", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("framebuffer", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("gfxcard", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("help", false, "Print usage.")
	rootCmd.Flags().Bool("hub", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("ide", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("isapnp", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("isdn", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("joystick", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("keyboard", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("listmd", false, "Normally hwinfo does not report RAID devices.")
	rootCmd.Flags().String("log", "", "Write log info to FILE.")
	rootCmd.Flags().Bool("map", false, "If disk names have  changed this prints a list of disk name mappings.")
	rootCmd.Flags().Bool("memory", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("mmc-ctrl", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("modem", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("monitor", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("mouse", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("netcard", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("network", false, "probe for particular hardware item")
	rootCmd.Flags().String("only", "", "This option can be given more than once.")
	rootCmd.Flags().Bool("partition", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("pci", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("pcmcia", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("pcmcia-ctrl", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("pppoe", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("printer", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("reallyall", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("redasd", false, "probe for particular hardware item")
	rootCmd.Flags().String("save-config", "", "Store config for a particular device below /var/lib/hardware.")
	rootCmd.Flags().Bool("scanner", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("scsi", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("short", false, "Show only a summary.")
	rootCmd.Flags().String("show-config", "", "Show saved config data for a particular device.")
	rootCmd.Flags().Bool("smp", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("sound", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("storage-ctrl", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("sys", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("tape", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("tv", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("uml", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("usb", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("usb-ctrl", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("vbe", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("verbose", false, "Increase verbosity.")
	rootCmd.Flags().Bool("version", false, "Print libhd version.")
	rootCmd.Flags().Bool("wlan", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("xen", false, "probe for particular hardware item")
	rootCmd.Flags().Bool("zip", false, "probe for particular hardware item")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dump-db": carapace.ActionValuesDescribed(
			"0", "external data base",
			"1", "internal data base",
		),
		"log": carapace.ActionFiles(),
		"save-config": carapace.Batch(
			action.ActionUniqueIds(),
			carapace.ActionValues("all"),
		).ToA(),
		"show-config": action.ActionUniqueIds(),
	})
}
