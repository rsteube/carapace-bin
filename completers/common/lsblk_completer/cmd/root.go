package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsblk",
	Short: "list block devices",
	Long:  "https://linux.die.net/man/8/lsblk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "print all devices")
	rootCmd.Flags().BoolP("ascii", "i", false, "use ascii characters only")
	rootCmd.Flags().BoolP("bytes", "b", false, "print SIZE in bytes rather than in human readable format")
	rootCmd.Flags().StringP("dedup", "E", "", "de-duplicate output by <column>")
	rootCmd.Flags().BoolP("discard", "D", false, "print discard capabilities")
	rootCmd.Flags().StringP("exclude", "e", "", "exclude devices by major number (default: RAM disks)")
	rootCmd.Flags().BoolP("fs", "f", false, "output info about filesystems")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("include", "I", "", "show only devices with specified major numbers")
	rootCmd.Flags().BoolP("inverse", "s", false, "inverse dependencies")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().BoolP("list", "l", false, "use list format output")
	rootCmd.Flags().BoolP("merge", "M", false, "group parents of sub-trees (usable for RAIDs, Multi-path)")
	rootCmd.Flags().BoolP("nodeps", "d", false, "don't print slaves or holders")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().StringP("output", "o", "", "output columns")
	rootCmd.Flags().BoolP("output-all", "O", false, "output all columns")
	rootCmd.Flags().BoolP("pairs", "P", false, "use key=\"value\" output format")
	rootCmd.Flags().BoolP("paths", "p", false, "print complete device path")
	rootCmd.Flags().BoolP("perms", "m", false, "output info about permissions")
	rootCmd.Flags().BoolP("raw", "r", false, "use raw output format")
	rootCmd.Flags().BoolP("scsi", "S", false, "output info about SCSI devices")
	rootCmd.Flags().StringP("sort", "x", "", "sort output by <column>")
	rootCmd.Flags().String("sysroot", "", "use specified directory as system root")
	rootCmd.Flags().BoolP("topology", "t", false, "output info about topology")
	rootCmd.Flags().StringP("tree", "T", "", "use tree format output")
	rootCmd.Flags().BoolP("version", "V", false, "display version")
	rootCmd.Flags().BoolP("zoned", "z", false, "print zone model")

	rootCmd.Flag("tree").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dedup": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"output": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"sort":    ActionColumns(),
		"sysroot": carapace.ActionDirectories(),
		"tree":    ActionColumns(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		fs.ActionBlockDevices(),
	)
}

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"NAME", "device name",
		"KNAME", "internal kernel device name",
		"PATH", "path to the device node",
		"MAJ:MIN", "major:minor device number",
		"FSAVAIL", "filesystem size available",
		"FSSIZE", "filesystem size",
		"FSTYPE", "filesystem type",
		"FSUSED", "filesystem size used",
		"FSUSE%", "filesystem use percentage",
		"FSVER", "filesystem version",
		"MOUNTPOINT", "where the device is mounted",
		"LABEL", "filesystem LABEL",
		"UUID", "filesystem UUID",
		"PTUUID", "partition table identifier (usually UUID)",
		"PTTYPE", "partition table type",
		"PARTTYPE", "partition type code or UUID",
		"PARTTYPENAME", "partition type name",
		"PARTLABEL", "partition LABEL",
		"PARTUUID", "partition UUID",
		"PARTFLAGS", "partition flags",
		"RA", "read-ahead of the device",
		"RO", "read-only device",
		"RM", "removable device",
		"HOTPLUG", "removable or hotplug device (usb, pcmcia, ...)",
		"MODEL", "device identifier",
		"SERIAL", "disk serial number",
		"SIZE", "size of the device",
		"STATE", "state of the device",
		"OWNER", "user name",
		"GROUP", "group name",
		"MODE", "device node permissions",
		"ALIGNMENT", "alignment offset",
		"MIN-IO", "minimum I/O size",
		"OPT-IO", "optimal I/O size",
		"PHY-SEC", "physical sector size",
		"LOG-SEC", "logical sector size",
		"ROTA", "rotational device",
		"SCHED", "I/O scheduler name",
		"RQ-SIZE", "request queue size",
		"TYPE", "device type",
		"DISC-ALN", "discard alignment offset",
		"DISC-GRAN", "discard granularity",
		"DISC-MAX", "discard max bytes",
		"DISC-ZERO", "discard zeroes data",
		"WSAME", "write same max bytes",
		"WWN", "unique storage identifier",
		"RAND", "adds randomness",
		"PKNAME", "internal parent kernel device name",
		"HCTL", "Host:Channel:Target:Lun for SCSI",
		"TRAN", "device transport type",
		"SUBSYSTEMS", "de-duplicated chain of subsystems",
		"REV", "device revision",
		"VENDOR", "device vendor",
		"ZONED", "zone model",
		"DAX", "dax-capable device",
	)
}
