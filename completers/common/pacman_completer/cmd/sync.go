package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func initSyncCmd(cmd *cobra.Command) {
	cmd.Flags().String("arch", "", "set an alternate architecture")
	cmd.Flags().Bool("asdeps", false, "install packages as non-explicitly installed")
	cmd.Flags().Bool("asexplicit", false, "install packages as explicitly installed")
	cmd.Flags().String("assume-installed", "", "")
	cmd.Flags().BoolP("clean", "c", false, "remove old packages from cache directory (-cc for all)")
	cmd.Flags().String("color", "", "colorize the output")
	cmd.Flags().String("config", "", "set an alternate configuration file")
	cmd.Flags().Bool("confirm", false, "always ask for confirmation")
	cmd.Flags().Bool("dbonly", false, "only modify database entries, not package files")
	cmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	cmd.Flags().Bool("debug", false, "display debug messages")
	cmd.Flags().Bool("disable-download-timeout", false, "")
	cmd.Flags().BoolP("downloadonly", "w", false, "download packages but do not install/upgrade anything")
	cmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	cmd.Flags().BoolP("groups", "g", false, "view all members of a package group")
	cmd.Flags().String("hookdir", "", "set an alternate hook location")
	cmd.Flags().String("ignore", "", "ignore a package upgrade (can be used more than once)")
	cmd.Flags().String("ignoregroup", "", "")
	cmd.Flags().BoolP("info", "i", false, "view package information (-ii for extended information)")
	cmd.Flags().StringP("list", "l", "", "view a list of packages in a repo")
	cmd.Flags().Bool("needed", false, "do not reinstall up to date packages")
	cmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	cmd.Flags().BoolP("nodeps", "d", false, "skip dependency version checks (-dd to skip all checks)")
	cmd.Flags().Bool("noprogressbar", false, "do not show a progress bar when downloading files")
	cmd.Flags().Bool("noscriptlet", false, "do not execute the install scriptlet if one exists")
	cmd.Flags().String("overwrite", "", "")
	cmd.Flags().BoolP("print", "p", false, "print the targets instead of performing the operation")
	cmd.Flags().String("print-format", "", "")
	cmd.Flags().BoolP("quiet", "q", false, "show less information for query and search")
	cmd.Flags().BoolP("refresh", "y", false, "download fresh package databases from the server")
	cmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	cmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	cmd.Flags().BoolP("sysupgrade", "u", false, "upgrade installed packages (-uu enables downgrades)")
	cmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"arch":    carapace.ActionValues("i686", "x86_64"),
		"color":   carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"config":  carapace.ActionFiles(),
		"dbpath":  carapace.ActionFiles(),
		"gpgdir":  carapace.ActionDirectories(),
		"hookdir": carapace.ActionDirectories(),
		// TODO ignore
		// TODO ingrnoregroup
		// TODO list
		// TODO overwrite
		"root": carapace.ActionDirectories(),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pacman.ActionPackages(pacman.PackageOption{}).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
