package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func initDeptestCmd(cmd *cobra.Command) {
	cmd.Flags().String("arch", "", "set an alternate architecture")
	cmd.Flags().String("color", "", "colorize the output")
	cmd.Flags().String("config", "", "set an alternate configuration file")
	cmd.Flags().Bool("confirm", false, "always ask for confirmation")
	cmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	cmd.Flags().Bool("debug", false, "display debug messages")
	cmd.Flags().Bool("disable-download-timeout", false, "")
	cmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	cmd.Flags().String("hookdir", "", "set an alternate hook location")
	cmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	cmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	cmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	cmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"arch":    carapace.ActionValues("i686", "x86_64"),
		"color":   carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"config":  carapace.ActionFiles(),
		"dbpath":  carapace.ActionFiles(),
		"gpgdir":  carapace.ActionDirectories(),
		"hookdir": carapace.ActionDirectories(),
		"root":    carapace.ActionDirectories(),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pacman.ActionPackages(pacman.PackageOption{}).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
