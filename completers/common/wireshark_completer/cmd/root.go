package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tshark"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wireshark",
	Short: "Interactively dump and analyze network traffic",
	Long:  "https://www.wireshark.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("C", "C", "", "start with specified configuration profile")
	rootCmd.Flags().BoolS("H", "H", false, "hide the capture info dialog during packet capture")
	rootCmd.Flags().StringS("J", "J", "", "jump to the first packet matching the (display) filter")
	rootCmd.Flags().StringS("K", "K", "", "keytab file to use for kerberos decryption")
	rootCmd.Flags().StringS("N", "N", "", "enable specific name resolution")
	rootCmd.Flags().StringS("P", "P", "", "personal files")
	rootCmd.Flags().BoolS("S", "S", false, "update packet display when new packets are captured")
	rootCmd.Flags().StringS("X", "X", "", "eXtension options")
	rootCmd.Flags().StringArrayP("autostop", "a", []string{}, "specify stop criterion")
	rootCmd.Flags().String("capture-comment", "", "set the capture file comment, if supported")
	rootCmd.Flags().StringArrayS("d", "d", []string{}, "specify layer type dissection")
	rootCmd.Flags().String("disable-heuristic", "", "disable dissection of heuristic protocol")
	rootCmd.Flags().String("disable-protocol", "", "disable dissection of proto_name")
	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().StringP("display-filter", "Y", "", "start with the given display filter")
	rootCmd.Flags().String("enable-heuristic", "", "enable dissection of heuristic protocol")
	rootCmd.Flags().String("enable-protocol", "", "enable dissection of proto_name")
	rootCmd.Flags().Bool("fullscreen", false, "start Wireshark in full screen")
	rootCmd.Flags().StringS("g", "g", "", "go to specified packet number after \"-r\"")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringP("interface", "i", "", "name or idx of interface")
	rootCmd.Flags().BoolS("j", "j", false, "search backwards for a matching packet after \"-J\"")
	rootCmd.Flags().BoolS("k", "k", false, "start capturing immediately")
	rootCmd.Flags().BoolS("l", "l", false, "turn on automatic scrolling while -S is in use")
	rootCmd.Flags().BoolP("list-data-link-types", "L", false, "print list of link-layer types of iface and exit")
	rootCmd.Flags().BoolP("list-interfaces", "D", false, "print list of interfaces and exit")
	rootCmd.Flags().Bool("list-time-stamp-types", false, "print list of timestamp types for iface and exit")
	rootCmd.Flags().BoolP("monitor-mode", "I", false, "capture in monitor mode, if available")
	rootCmd.Flags().BoolS("n", "n", false, "disable all name resolutions")
	rootCmd.Flags().BoolP("no-promiscuous-mode", "p", false, "don't capture in promiscuous mode")
	rootCmd.Flags().StringArrayS("o", "o", []string{}, "override preference or recent setting")
	rootCmd.Flags().StringP("read-file", "r", "", "set the filename to read from")
	rootCmd.Flags().StringP("read-filter", "R", "", "packet filter in Wireshark display filter syntax")
	rootCmd.Flags().StringArrayP("ring-buffer", "b", []string{}, "run in 'multiple files' mode")
	rootCmd.Flags().StringP("snapshot-length", "s", "", "packet snapshot length")
	rootCmd.Flags().StringS("t", "t", "", "format of time stamps")
	rootCmd.Flags().String("time-stamp-type", "", "timestamp method for interface")
	rootCmd.Flags().StringS("u", "u", "", "output format of seconds")
	rootCmd.Flags().BoolP("version", "v", false, "display version info and exit")
	rootCmd.Flags().StringS("w", "w", "", "set the output filename")
	rootCmd.Flags().StringS("z", "z", "", "show various statistics")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"K": carapace.ActionFiles(),
		"P": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues(
					"persconf", "personal configuration files",
					"persdata", "personal data files",
				).Invoke(c).Suffix(":").ToA()
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"autostop": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"duration", "stop after NUM seconds",
					"filesize", "stop this file after NUM KB",
					"files", "stop after NUM files",
					"packets", "stop after NUM packets",
				).Invoke(c).Suffix(":").ToA()
			default:
				return carapace.ActionValues()
			}
		}),
		"d": carapace.ActionMultiParts("==", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return tshark.ActionSelectors().Invoke(c).Suffix("==").ToA()
			case 1:
				return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
							switch len(c.Parts) {
							case 0:
								return net.ActionPorts()
							case 1:
								return net.ActionPorts()
							default:
								return carapace.ActionValues()
							}
						})
					case 1:
						return tshark.ActionProtocols()
					default:
						return carapace.ActionValues()
					}
				})
			default:
				return carapace.ActionValues()
			}
		}),
		"disable-protocol": tshark.ActionProtocols(),
		"display":          os.ActionDisplays(),
		"enable-protocol":  tshark.ActionProtocols(),
		"interface":        tshark.ActionInterfaces(),
		"read-file":        carapace.ActionFiles(),
		"ring-buffer": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"duration", "switch to next file after NUM secs",
					"filesize", "switch to next file after NUM KB",
					"files", "ringbuffer: replace after NUM files",
					"packets", "switch to next file after NUM packets",
					"interval", "switch to next file when the time is an exact multiple of NUM secs",
				).Invoke(c).Suffix(":").ToA()
			default:
				return carapace.ActionValues()
			}
		}),
		"t": carapace.ActionValues("a", "ad", "adoy", "d", "dd", "e", "r", "u", "ud", "udoy"),
		"u": carapace.ActionValues("s", "hms"),
		"w": carapace.ActionFiles(),
		"z": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return tshark.ActionStatistics().Invoke(c).ToMultiPartsA(",")
		}),
	})
}
