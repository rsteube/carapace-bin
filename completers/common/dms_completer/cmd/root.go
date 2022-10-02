package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dms",
	Short: "A UPnP DLNA Digital Media Server",
	Long:  "https://github.com/anacrolix/dms",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("allowedIps", "allowedIps", "", "allowed ip of clients, separated by comma")
	rootCmd.Flags().StringS("config", "config", "", "json configuration file")
	rootCmd.Flags().StringS("deviceIcon", "deviceIcon", "", "device icon")
	rootCmd.Flags().StringS("fFprobeCachePath", "fFprobeCachePath", "", "path to FFprobe cache file (default \"~/.dms-ffprobe-cache\")")
	rootCmd.Flags().StringS("forceTranscodeTo", "forceTranscodeTo", "", "force transcoding to certain format, supported: 'chromecast', 'vp8', 'web'")
	rootCmd.Flags().StringS("friendlyName", "friendlyName", "", "server friendly name")
	rootCmd.Flags().StringS("http", "http", "", "http server port (default \":1338\")")
	rootCmd.Flags().StringS("ifname", "ifname", "", "specific SSDP network interface")
	rootCmd.Flags().BoolS("ignoreHidden", "ignoreHidden", false, "ignore hidden files and directories")
	rootCmd.Flags().BoolS("ignoreUnreadable", "ignoreUnreadable", false, "ignore unreadable files and directories")
	rootCmd.Flags().BoolS("logHeaders", "logHeaders", false, "log HTTP headers")
	rootCmd.Flags().BoolS("noProbe", "noProbe", false, "disable media probing with ffprobe")
	rootCmd.Flags().BoolS("noTranscode", "noTranscode", false, "disable transcoding")
	rootCmd.Flags().StringS("notifyInterval", "notifyInterval", "", "interval between SSPD announces (default 30s)")
	rootCmd.Flags().StringS("path", "path", "", "browse root path")
	rootCmd.Flags().BoolS("stallEventSubscribe", "stallEventSubscribe", false, "workaround for some bad event subscribers")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":           carapace.ActionFiles(".json"),
		"fFprobeCachePath": carapace.ActionFiles(),
		"forceTranscodeTo": carapace.ActionValues("chromecast", "vp8", "web"),
		"ifname":           net.ActionDevices(net.IncludedDevices{Ethernet: true, Wifi: true}),
		"path":             carapace.ActionDirectories(),
	})
}
