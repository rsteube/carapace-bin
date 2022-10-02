package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ntpd",
	Short: "NTP daemon program",
	Long:  "https://linux.die.net/man/8/ntpd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("!", "!", false, "extended usage information passed thru pager")
	rootCmd.Flags().BoolS("4", "4", false, "Force IPv4 DNS name resolution")
	rootCmd.Flags().BoolS("6", "6", false, "Force IPv6 DNS name resolution")
	rootCmd.Flags().BoolS("?", "?", false, "display extended usage information and exit")
	rootCmd.Flags().BoolS("A", "A", false, "Do not require crypto authentication")
	rootCmd.Flags().StringS("D", "D", "", "Set the debug verbosity level")
	rootCmd.Flags().BoolS("G", "G", false, "Step any initial offset correction.")
	rootCmd.Flags().StringS("I", "I", "", "Listen on an interface name or address")
	rootCmd.Flags().BoolS("L", "L", false, "Do not listen to virtual interfaces")
	rootCmd.Flags().BoolS("N", "N", false, "Run at high priority")
	rootCmd.Flags().StringS("P", "P", "", "Process priority")
	rootCmd.Flags().StringS("U", "U", "", "interval in seconds between scans for new or dropped interfaces")
	rootCmd.Flags().BoolS("a", "a", false, "Require crypto authentication")
	rootCmd.Flags().BoolS("b", "b", false, "Allow us to sync to broadcast servers")
	rootCmd.Flags().StringS("c", "c", "", "configuration file name")
	rootCmd.Flags().BoolS("d", "d", false, "Increase debug verbosity level")
	rootCmd.Flags().StringS("f", "f", "", "frequency drift file name")
	rootCmd.Flags().BoolS("g", "g", false, "Allow the first adjustment to be Big")
	rootCmd.Flags().StringS("i", "i", "", "Jail directory")
	rootCmd.Flags().StringS("k", "k", "", "path to symmetric keys")
	rootCmd.Flags().StringS("l", "l", "", "path to the log file")
	rootCmd.Flags().BoolS("n", "n", false, "Do not fork")
	rootCmd.Flags().StringS("p", "p", "", "path to the PID file")
	rootCmd.Flags().BoolS("q", "q", false, "Set the time and quit")
	rootCmd.Flags().StringS("r", "r", "", "Broadcast/propagation delay")
	rootCmd.Flags().StringS("s", "s", "", "Statistics file location")
	rootCmd.Flags().StringS("t", "t", "", "Trusted key number")
	rootCmd.Flags().StringS("u", "u", "", "Run as userid (or userid:groupid)")
	rootCmd.Flags().StringS("w", "w", "", "Seconds to wait for first clock sync")
	rootCmd.Flags().BoolS("x", "x", false, "Slew up to 600 seconds")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"I": net.ActionDevices(net.AllDevices),
		"c": carapace.ActionFiles(),
		"f": carapace.ActionFiles(),
		"i": carapace.ActionDirectories(),
		"l": carapace.ActionFiles(),
		"p": carapace.ActionFiles(),
		"s": carapace.ActionDirectories(),
		"u": os.ActionUserGroup(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return ActionPoolZones().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}

func ActionPoolZones() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ad.pool.ntp.org", "Andorra",
		"al.pool.ntp.org", "Albania",
		"at.pool.ntp.org", "Austria",
		"ax.pool.ntp.org", "Aland Islands",
		"ba.pool.ntp.org", "Bosnia and Herzegovina",
		"be.pool.ntp.org", "Belgium",
		"bg.pool.ntp.org", "Bulgaria",
		"by.pool.ntp.org", "Belarus",
		"ch.pool.ntp.org", "Switzerland",
		"cz.pool.ntp.org", "Czech Republic",
		"de.pool.ntp.org", "Germany",
		"dk.pool.ntp.org", "Denmark",
		"ee.pool.ntp.org", "Estonia",
		"es.pool.ntp.org", "Spain",
		"fi.pool.ntp.org", "Finland",
		"fo.pool.ntp.org", "Faroe Islands",
		"fr.pool.ntp.org", "France",
		"gg.pool.ntp.org", "Guernsey",
		"gi.pool.ntp.org", "Gibraltar",
		"gr.pool.ntp.org", "Greece",
		"hr.pool.ntp.org", "Croatia",
		"hu.pool.ntp.org", "Hungary",
		"ie.pool.ntp.org", "Ireland",
		"im.pool.ntp.org", "Isle of Man",
		"is.pool.ntp.org", "Iceland",
		"it.pool.ntp.org", "Italy",
		"je.pool.ntp.org", "Jersey",
		"li.pool.ntp.org", "Liechtenstein",
		"lt.pool.ntp.org", "Lithuania",
		"lu.pool.ntp.org", "Luxembourg",
		"lv.pool.ntp.org", "Latvia",
		"mc.pool.ntp.org", "Monaco",
		"md.pool.ntp.org", "Moldova",
		"me.pool.ntp.org", "Republic of Montenegro",
		"mk.pool.ntp.org", "Macedonia",
		"mt.pool.ntp.org", "Malta",
		"nl.pool.ntp.org", "Netherlands",
		"no.pool.ntp.org", "Norway",
		"pl.pool.ntp.org", "Poland",
		"pt.pool.ntp.org", "Portugal",
		"ro.pool.ntp.org", "Romania",
		"rs.pool.ntp.org", "Republic of Serbia",
		"ru.pool.ntp.org", "Russian Federation",
		"se.pool.ntp.org", "Sweden",
		"si.pool.ntp.org", "Slovenia",
		"sj.pool.ntp.org", "Svalbard and Jan Mayen",
		"sk.pool.ntp.org", "Slovakia",
		"sm.pool.ntp.org", "San Marino",
		"tr.pool.ntp.org", "Turkey",
		"ua.pool.ntp.org", "Ukraine",
		"uk.pool.ntp.org", "United Kingdom",
		"va.pool.ntp.org", "Holy See (Vatican City State)",
		"yu.pool.ntp.org", "Yugoslavia",
	)
}
