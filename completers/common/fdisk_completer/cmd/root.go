package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fdisk",
	Short: "manipulate disk partition table",
	Long:  "https://en.wikipedia.org/wiki/Fdisk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("bytes", false, "print SIZE in bytes rather than in human readable format")
	rootCmd.Flags().StringP("color", "L", "", "colorize output")
	rootCmd.Flags().StringP("compatibility", "c", "", "mode is 'dos' or 'nondos' (default)")
	rootCmd.Flags().StringP("cylinders", "C", "", "specify the number of cylinders")
	rootCmd.Flags().BoolP("getsz", "s", false, "display device size in 512-byte sectors [DEPRECATED]")
	rootCmd.Flags().StringP("heads", "H", "", "specify the number of heads")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("list", "l", false, "display partitions and exit")
	rootCmd.Flags().BoolP("list-details", "x", false, "like --list but with more details")
	rootCmd.Flags().String("lock", "", "use exclusive device lock (yes, no or nonblock)")
	rootCmd.Flags().BoolP("noauto-pt", "n", false, "don't create default partition table on empty devices")
	rootCmd.Flags().StringP("output", "o", "", "output columns")
	rootCmd.Flags().BoolP("protect-boot", "B", false, "don't erase bootbits when creating a new label")
	rootCmd.Flags().StringP("sector-size", "b", "", "physical and logical sector size")
	rootCmd.Flags().StringP("sectors", "S", "", "specify the number of sectors per track")
	rootCmd.Flags().StringP("type", "t", "", "recognize specified partition table type only")
	rootCmd.Flags().StringP("units", "u", "", "display units: 'cylinders' or 'sectors' (default)")
	rootCmd.Flags().BoolP("version", "V", false, "display version")
	rootCmd.Flags().StringP("wipe", "w", "", "wipe signatures (auto, always or never)")
	rootCmd.Flags().StringP("wipe-partitions", "W", "", "wipe signatures from new partitions (auto, always or never)")

	rootCmd.Flag("compatibility").NoOptDefVal = "nondos"
	rootCmd.Flag("color").NoOptDefVal = "auto"
	rootCmd.Flag("units").NoOptDefVal = "sectors"
	rootCmd.Flag("lock").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":           carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"compatibility":   carapace.ActionValues("dos", "nondos"),
		"lock":            carapace.ActionValues("yes", "no", "nonblock").StyleF(style.ForKeyword),
		"units":           carapace.ActionValues("cylinders", "sectors"),
		"wipe-partitions": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			fs.ActionBlockDevices(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("list").Changed {
				return carapace.Batch(
					fs.ActionBlockDevices(),
					carapace.ActionFiles(),
				).ToA()
			}
			return carapace.ActionValues()
		}),
	)
}
