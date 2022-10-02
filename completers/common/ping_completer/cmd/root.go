package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ping",
	Short: "send ICMP ECHO_REQUEST to network hosts",
	Long:  "https://linux.die.net/man/8/ping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "use IPv4")
	rootCmd.Flags().BoolS("6", "6", false, "use IPv6")
	rootCmd.Flags().BoolS("A", "A", false, "use adaptive ping")
	rootCmd.Flags().BoolS("B", "B", false, "sticky source address")
	rootCmd.Flags().BoolS("D", "D", false, "print timestamps")
	rootCmd.Flags().StringS("F", "F", "", "define flow label, default is random")
	rootCmd.Flags().StringS("I", "I", "", "either interface name or address")
	rootCmd.Flags().BoolS("L", "L", false, "suppress loopback of multicast packets")
	rootCmd.Flags().StringS("M", "M", "", "define mtu discovery")
	rootCmd.Flags().BoolS("O", "O", false, "report outstanding replies")
	rootCmd.Flags().StringS("Q", "Q", "", "use quality of service <tclass> bits")
	rootCmd.Flags().BoolS("R", "R", false, "record route")
	rootCmd.Flags().StringS("S", "S", "", "use <size> as SO_SNDBUF socket option value")
	rootCmd.Flags().StringS("T", "T", "", "define timestamp, can be one of <tsonly|tsandaddr|tsprespec>")
	rootCmd.Flags().BoolS("U", "U", false, "print user-to-user latency")
	rootCmd.Flags().BoolS("V", "V", false, "print version and exit")
	rootCmd.Flags().StringS("W", "W", "", "time to wait for response")
	rootCmd.Flags().BoolS("a", "a", false, "use audible ping")
	rootCmd.Flags().BoolS("b", "b", false, "allow pinging broadcast")
	rootCmd.Flags().StringS("c", "c", "", "stop after <count> replies")
	rootCmd.Flags().BoolS("d", "d", false, "use SO_DEBUG socket option")
	rootCmd.Flags().BoolS("f", "f", false, "flood ping")
	rootCmd.Flags().BoolS("h", "h", false, "print help and exit")
	rootCmd.Flags().StringS("i", "i", "", "seconds between sending each packet")
	rootCmd.Flags().StringS("l", "l", "", "send <preload> number of packages while waiting replies")
	rootCmd.Flags().StringS("m", "m", "", "tag the packets going out")
	rootCmd.Flags().BoolS("n", "n", false, "no dns name resolution")
	rootCmd.Flags().StringS("p", "p", "", "contents of padding byte")
	rootCmd.Flags().BoolS("q", "q", false, "quiet output")
	rootCmd.Flags().StringS("s", "s", "", "use <size> as number of data bytes to be sent")
	rootCmd.Flags().StringS("t", "t", "", "define time to live")
	rootCmd.Flags().BoolS("v", "v", false, "verbose output")
	rootCmd.Flags().StringS("w", "w", "", "reply wait <deadline> in seconds")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"M": carapace.ActionValues("do", "dont", "want"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		net.ActionHosts(),
	)
}
