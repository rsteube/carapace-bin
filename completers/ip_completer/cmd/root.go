package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ip",
	Short: "show / manipulate routing, network devices, interfaces and tunnels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("iec", "iec", false, "print human readable rates in IEC units")
	rootCmd.Flags().BoolS("4", "4", false, "shortcut for -family inet")
	rootCmd.Flags().BoolS("human-readable", "human-readable", false, "output statistics with human readable values followed by suffix")
	rootCmd.Flags().BoolS("0", "0", false, "shortcut for -family link")
	rootCmd.Flags().BoolS("M", "M", false, "shortcut for -family mpls")
	rootCmd.Flags().BoolS("B", "B", false, "shortcut for -family bridge")
	rootCmd.Flags().BoolS("echo", "echo", false, "Request the kernel to send the applied configuration back")
	rootCmd.Flags().BoolS("statistics", "statistics", false, "Output more information")
	rootCmd.Flags().BoolS("force", "force", false, "Don't terminate ip on errors in batch mode")
	rootCmd.Flags().BoolS("6", "6", false, "shortcut for -family inet6")
	rootCmd.Flags().BoolN("Numeric", "N", false, "Print the number of protocol, scope, dsfield, etc di‐ rectly instead of converting it to human readable name")
	rootCmd.Flags().BoolN("Version", "V", false, "Print the version of the ip utility and exit")
	rootCmd.Flags().BoolN("all", "a", false, "executes specified command over all objects, it depends if command supports this option")
	rootCmd.Flags().StringN("batch", "b", "", "Read commands from provided file or standard input and invoke them")
	rootCmd.Flags().BoolN("brief", "br", false, "Print only basic information in a tabular format for better readability")
	rootCmd.Flags().StringN("color", "c", "", "Configure color output")
	rootCmd.Flags().BoolN("details", "d", false, "Output more detailed information")
	rootCmd.Flags().StringN("family", "f", "", "Specifies the protocol family to use")
	rootCmd.Flags().BoolN("human", "h", false, "output statistics with human readable values followed by suffix")
	rootCmd.Flags().BoolN("json", "j", false, "Output results in JavaScript Object Notation (JSON)")
	rootCmd.Flags().StringN("loops", "l", "", "Specify maximum number of loops the 'ip address flush' logic will attempt before giving up")
	rootCmd.Flags().StringN("netns", "n", "", "switches ip to the specified network namespace NETNS")
	rootCmd.Flags().BoolN("oneline", "o", false, "output each record on a single line, replacing line feeds with the '\\' character")
	rootCmd.Flags().BoolN("pretty", "p", false, "The default JSON format is compact and more efficient to")
	rootCmd.Flags().StringN("rcvbuf", "rc", "", "Set the netlink socket receive buffer size, defaults to 1MB")
	rootCmd.Flags().BoolN("resolve", "r", false, "use the system's name resolver to print DNS names instead of host addresses")
	rootCmd.Flags().BoolN("stats", "s", false, "Output more information")
	rootCmd.Flags().BoolN("timestamp", "t", false, "display current time when using monitor option")
	rootCmd.Flags().BoolN("tshort", "ts", false, "Like -timestamp, but use shorter format")

	rootCmd.Flag("color").NoOptDefVal = " "

	rootCmd.MarkFlagsMutuallyExclusive("human-readable", "human")
	rootCmd.MarkFlagsMutuallyExclusive("statistics", "stats")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"batch":  carapace.ActionFiles(),
		"color":  carapace.ActionValues("always", "auto", "never").StyleF(style.ForKeyword),
		"family": carapace.ActionValues("inet", "inet6", "bridge", "mpls", "link"),
	})
}
