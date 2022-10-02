package cmd

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tshark"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tshark",
	Short: "Dump and analyze network traffic",
	Long:  "https://tshark.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	var extensionOpts []string

	rootCmd.Flags().BoolS("2", "2", false, "perform a two-pass analysis")
	rootCmd.Flags().StringS("C", "C", "", "start with specified configuration profile")
	rootCmd.Flags().StringS("E", "E", "", "set options for output")
	rootCmd.Flags().StringS("F", "F", "", "set the output file type")
	rootCmd.Flags().StringS("G", "G", "", "dump one of several available reports and exit")
	rootCmd.Flags().StringS("H", "H", "", "read a list of entries from a hosts file")
	rootCmd.Flags().StringS("J", "J", "", "top level protocol filter")
	rootCmd.Flags().StringS("K", "K", "", "keytab file to use for kerberos decryption")
	rootCmd.Flags().StringS("M", "M", "", "perform session auto reset")
	rootCmd.Flags().StringS("N", "N", "", "enable specific name resolution")
	rootCmd.Flags().StringS("O", "O", "", "Only show packet details of these protocols")
	rootCmd.Flags().BoolS("Q", "Q", false, "only log true errors to stderr")
	rootCmd.Flags().StringS("S", "S", "", "the line separator to print between packets")
	rootCmd.Flags().StringS("T", "T", "", "format of text output")
	rootCmd.Flags().StringS("U", "U", "", "PDUs export mode, see the man page for details")
	rootCmd.Flags().BoolS("V", "V", false, "add output of packet tree")
	rootCmd.Flags().StringS("W", "W", "", "Save extra information in the file, if supported.")
	rootCmd.Flags().StringArrayVarS(&extensionOpts, "X", "X", []string{}, "eXtension options, see the man page for details")
	rootCmd.Flags().StringArrayP("autostop", "a", []string{}, "specify stop criterion ")
	rootCmd.Flags().StringP("buffer-size", "B", "", "size of kernel buffer")
	rootCmd.Flags().StringS("c", "c", "", "stop after n packets (def: infinite)")
	rootCmd.Flags().String("capture-comment", "", "set the capture file comment")
	rootCmd.Flags().Bool("color", false, "color output text similarly to the Wireshark GUI")
	rootCmd.Flags().StringArrayS("d", "d", []string{}, "specify layer type dissection")
	rootCmd.Flags().String("disable-heuristic", "", "disable dissection of heuristic protocol")
	rootCmd.Flags().String("disable-protocol", "", "disable dissection of proto_name")
	rootCmd.Flags().StringS("e", "e", "", "field to print")
	rootCmd.Flags().String("elastic-mapping-filter", "", "put only the specified protocols within the mapping file")
	rootCmd.Flags().String("enable-heuristic", "", "enable dissection of heuristic protocol")
	rootCmd.Flags().Bool("enable-protocol", false, "<proto_name>   enable dissection of proto_name")
	rootCmd.Flags().String("export-objects", "", "save exported objects for a protocol to a directory")
	rootCmd.Flags().StringS("f", "f", "", "packet filter in libpcap filter syntax")
	rootCmd.Flags().BoolS("g", "g", false, "enable group read access on the output file(s)")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringP("interface", "i", "", "name or idx of interface")
	rootCmd.Flags().StringS("j", "j", "", "protocols layers filter")
	rootCmd.Flags().BoolS("l", "l", false, "flush standard output after each packet")
	rootCmd.Flags().StringP("linktype", "y", "", "link layer type")
	rootCmd.Flags().BoolP("list-data-link-types", "L", false, "print list of link-layer types of iface and exit")
	rootCmd.Flags().BoolP("list-interfaces", "D", false, "print list of interfaces and exit")
	rootCmd.Flags().Bool("list-time-stamp-types", false, "print list of timestamp types for iface and exit")
	rootCmd.Flags().BoolP("monitor-mode", "I", false, "capture in monitor mode, if available")
	rootCmd.Flags().BoolS("n", "n", false, "disable all name resolutions (def: all enabled)")
	rootCmd.Flags().Bool("no-duplicate-keys", false, "merge duplicate keys in an object into a single key")
	rootCmd.Flags().BoolP("no-promiscuous-mode", "p", false, "don't capture in promiscuous mode")
	rootCmd.Flags().StringArrayS("o", "o", []string{}, "override preference setting")
	rootCmd.Flags().BoolP("print", "P", false, "print packet summary even when writing to a file")
	rootCmd.Flags().BoolS("q", "q", false, "be more quiet on stdout")
	rootCmd.Flags().StringP("read-file", "r", "", "set the filename to read from")
	rootCmd.Flags().StringArrayP("ring-buffer", "b", []string{}, "run in 'multiple files' mode")
	rootCmd.Flags().StringP("snapshot-length", "s", "", "packet snapshot length")
	rootCmd.Flags().StringS("t", "t", "", "output format of time stamps")
	rootCmd.Flags().String("time-stamp-type", "", "timestamp method for interface")
	rootCmd.Flags().StringS("u", "u", "", "output format of seconds")
	rootCmd.Flags().BoolP("version", "v", false, "display version info and exit")
	rootCmd.Flags().StringS("w", "w", "", "write packets to a pcapng-format file named \"outfile\"")
	rootCmd.Flags().BoolS("x", "x", false, "add output of hex and ASCII dump (Packet Bytes)")
	rootCmd.Flags().StringS("z", "z", "", "various statistics, see the man page for details")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"E": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"bom", "print a UTF-8 BOM",
					"header", "switch headers on and off",
					"separator", "select tab, space, printable character as separator",
					"occurrence", "print first, last or all occurrences of each field",
					"aggregator", "select comma, space, printable character as aggregator",
					"quote", "select double, single, no quotes for values",
				).Invoke(c).Suffix("=").ToA()
			case 1:
				switch c.Parts[0] {
				case "bom":
					return carapace.ActionValues("y", "n")
				case "header":
					return carapace.ActionValues("y", "n")
				case "separator":
					return carapace.ActionValues("/t", "/s")
				case "occurrence":
					return carapace.ActionValues("f", "l", "a")
				case "aggregator":
					return carapace.ActionValues(",", "/s")
				case "quote":
					return carapace.ActionValues("d", "s", "n")
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"F": tshark.ActionFileTypes(),
		"G": carapace.ActionValuesDescribed(
			"column-formats", "dump column format codes and exit",
			"decodes", "dump \"layer type\"/\"decode as\" associations and exit",
			"dissector-tables", "dump dissector table names, types, and properties",
			"elastic-mapping", "dump ElasticSearch mapping file",
			"fieldcount", "dump count of header fields and exit",
			"fields", "dump fields glossary and exit",
			"ftypes", "dump field type basic and descriptive names",
			"heuristic-decodes", "dump heuristic dissector tables",
			"plugins", "dump installed plugins and exit",
			"protocols", "dump protocols in registration database and exit",
			"values", "dump value, range, true/false strings and exit",
			"currentprefs", "dump current preferences and exit",
			"defaultprefs", "dump default preferences and exit",
			"folders", "dump about:folders",
		),
		"K": carapace.ActionFiles(),
		"T": carapace.ActionValues("pdml", "ps", "psml", "json", "jsonraw", "ek", "tabs", "text", "fields", "?"),
		"X": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				keys := []string{
					"lua_script", "tells TShark to load the given script in addition to the default",
					"read_format", "tells TShark to use the given file format to read in the file",
				}
				for index, opt := range extensionOpts {
					if strings.HasPrefix(opt, "lua_script:") {
						keys = append(keys, fmt.Sprintf("lua_script%v", index+1), "pass argument to "+strings.SplitN(opt, ":", 2)[1])
					}
				}
				return carapace.ActionValuesDescribed(keys...).Invoke(c).Suffix(":").ToA()
			case 1:
				switch c.Parts[0] {
				case "lua_script":
					return carapace.ActionFiles()
				case "read_format":
					return tshark.ActionReadFormats()
				default:
					return carapace.ActionValues()
				}
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
		"elastic-mapping-filter": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return tshark.ActionProtocols().Invoke(c).Filter(c.Parts).ToA()
		}),
		"enable-protocol": tshark.ActionProtocols(),
		"export-objects": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("dicom", "http", "imf", "smb", "tftp").Invoke(c).Suffix(",").ToA()
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"interface": tshark.ActionInterfaces(),
		"linktype": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return tshark.ActionLinkTypes(rootCmd.Flag("interface").Value.String())
		}),
		"read-file": carapace.ActionFiles(),
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
