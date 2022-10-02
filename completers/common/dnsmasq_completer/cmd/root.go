package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnsmasq",
	Short: "A lightweight DHCP and caching DNS server",
	Long:  "https://en.wikipedia.org/wiki/Dnsmasq",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("add-cpe-id", "", "Add client identification to forwarded DNS queries.")
	rootCmd.Flags().String("add-mac", "", "Add requestor's MAC address to forwarded DNS queries.")
	rootCmd.Flags().StringP("addn-hosts", "H", "", "Specify a hosts file to be read in addition to /etc/hosts.")
	rootCmd.Flags().StringP("address", "A", "", "Return ipaddr for all hosts in specified domains.")
	rootCmd.Flags().StringP("alias", "V", "", "Translate IPv4 addresses from upstream servers.")
	rootCmd.Flags().Bool("all-servers", false, "Always perform DNS queries to all servers.")
	rootCmd.Flags().String("auth-peer", "", "Peers which are allowed to do zone transfer")
	rootCmd.Flags().String("auth-sec-servers", "", "Secondary authoritative nameservers for forward domains")
	rootCmd.Flags().String("auth-server", "", "Export local names to global DNS")
	rootCmd.Flags().String("auth-soa", "", "Set authoritative zone information")
	rootCmd.Flags().String("auth-ttl", "", "Set TTL for authoritative replies")
	rootCmd.Flags().String("auth-zone", "", "Domain to export to global DNS")
	rootCmd.Flags().Bool("bind-dynamic", false, "Bind to interfaces in use - check for new interfaces")
	rootCmd.Flags().BoolP("bind-interfaces", "z", false, "Bind only to interfaces in use.")
	rootCmd.Flags().StringP("bogus-nxdomain", "B", "", "Treat ipaddr as NXDOMAIN (defeats Verisign wildcard).")
	rootCmd.Flags().BoolP("bogus-priv", "b", false, "Fake reverse lookups for RFC1918 private address ranges.")
	rootCmd.Flags().StringP("bootp-dynamic", "3", "", "Enable dynamic address allocation for bootp.")
	rootCmd.Flags().String("bridge-interface", "", "Treat DHCP requests on aliases as arriving from interface.")
	rootCmd.Flags().String("caa-record", "", "Specify certification authority authorization record")
	rootCmd.Flags().StringP("cache-size", "c", "", "Specify the size of the cache in entries (defaults to 150).")
	rootCmd.Flags().Bool("clear-on-reload", false, "Clear DNS cache when reloading /etc/resolv.conf.")
	rootCmd.Flags().String("cname", "", "Specify alias name for LOCAL DNS name.")
	rootCmd.Flags().StringP("conf-dir", "7", "", "Read configuration from all the files in this directory.")
	rootCmd.Flags().StringP("conf-file", "C", "", "Specify configuration file (defaults to /etc/dnsmasq.conf).")
	rootCmd.Flags().Bool("conntrack", false, "Copy connection-track mark from queries to upstream connections.")
	rootCmd.Flags().String("dhcp-alternate-port", "", "Use alternative ports for DHCP.")
	rootCmd.Flags().BoolP("dhcp-authoritative", "K", false, "Assume we are the only DHCP server on the local network.")
	rootCmd.Flags().String("dhcp-broadcast", "", "Force broadcast replies for hosts with tag set.")
	rootCmd.Flags().String("dhcp-circuitid", "", "Map RFC3046 circuit-id to tag.")
	rootCmd.Flags().Bool("dhcp-client-update", false, "Allow DHCP clients to do their own DDNS updates.")
	rootCmd.Flags().String("dhcp-duid", "", "Specify DUID_EN-type DHCPv6 server DUID")
	rootCmd.Flags().Bool("dhcp-fqdn", false, "Use only fully qualified domain names for DHCP clients.")
	rootCmd.Flags().String("dhcp-generate-names", "", "Generate hostnames based on MAC address for nameless clients.")
	rootCmd.Flags().StringP("dhcp-host", "G", "", "Set address or hostname for a specified machine.")
	rootCmd.Flags().String("dhcp-hostsdir", "", "Read DHCP host specs from a directory.")
	rootCmd.Flags().String("dhcp-hostsfile", "", "Read DHCP host specs from file.")
	rootCmd.Flags().StringP("dhcp-ignore", "J", "", "Don't do DHCP for hosts with tag set.")
	rootCmd.Flags().Bool("dhcp-ignore-clid", false, "Ignore client identifier option sent by DHCP clients.")
	rootCmd.Flags().String("dhcp-ignore-names", "", "Ignore hostnames provided by DHCP clients.")
	rootCmd.Flags().StringP("dhcp-lease-max", "X", "", "Specify maximum number of DHCP leases (defaults to 1000).")
	rootCmd.Flags().StringP("dhcp-leasefile", "l", "", "Specify where to store DHCP leases (defaults to /var/lib/misc/dnsmasq.leases).")
	rootCmd.Flags().String("dhcp-luascript", "", "Lua script to run on DHCP lease creation and destruction.")
	rootCmd.Flags().String("dhcp-match", "", "Set tag if client includes matching option in request.")
	rootCmd.Flags().String("dhcp-name-match", "", "Set tag if client provides given name.")
	rootCmd.Flags().Bool("dhcp-no-override", false, "Do NOT reuse filename and server fields for extra DHCP options.")
	rootCmd.Flags().StringP("dhcp-option", "O", "", "Specify options to be sent to DHCP clients.")
	rootCmd.Flags().String("dhcp-option-force", "", "DHCP option sent even if the client does not request it.")
	rootCmd.Flags().String("dhcp-optsdir", "", "Read DHCP options from a directory.")
	rootCmd.Flags().String("dhcp-optsfile", "", "Read DHCP option specs from file.")
	rootCmd.Flags().String("dhcp-proxy", "", "Use these DHCP relays as full proxies.")
	rootCmd.Flags().String("dhcp-pxe-vendor", "", "Specify vendor class to match for PXE requests.")
	rootCmd.Flags().StringP("dhcp-range", "F", "", "Enable DHCP in the range given with lease duration.")
	rootCmd.Flags().Bool("dhcp-rapid-commit", false, "Enables DHCPv4 Rapid Commit option.")
	rootCmd.Flags().String("dhcp-relay", "", "Relay DHCP requests to a remote server")
	rootCmd.Flags().String("dhcp-remoteid", "", "Map RFC3046 remote-id to tag.")
	rootCmd.Flags().String("dhcp-reply-delay", "", "Delay DHCP replies for at least number of seconds.")
	rootCmd.Flags().StringP("dhcp-script", "6", "", "Shell script to run on DHCP lease creation and destruction.")
	rootCmd.Flags().String("dhcp-scriptuser", "", "Run lease-change scripts as this user.")
	rootCmd.Flags().Bool("dhcp-sequential-ip", false, "Attempt to allocate sequential IP addresses to DHCP clients.")
	rootCmd.Flags().String("dhcp-subscrid", "", "Map RFC3993 subscriber-id to tag.")
	rootCmd.Flags().String("dhcp-ttl", "", "Set TTL in DNS responses with DHCP-derived addresses.")
	rootCmd.Flags().StringP("dhcp-userclass", "j", "", "Map DHCP user class to tag.")
	rootCmd.Flags().StringP("dhcp-vendorclass", "U", "", "Map DHCP vendor class to tag.")
	rootCmd.Flags().StringP("dns-forward-max", "0", "", "Maximum number of concurrent DNS queries. (defaults to 150)")
	rootCmd.Flags().Bool("dns-loop-detect", false, "Detect and remove DNS forwarding loops.")
	rootCmd.Flags().String("dns-rr", "", "Specify arbitrary DNS resource record")
	rootCmd.Flags().Bool("dnssec", false, "Activate DNSSEC validation")
	rootCmd.Flags().Bool("dnssec-check-unsigned", false, "Ensure answers without DNSSEC are in unsigned zones.")
	rootCmd.Flags().Bool("dnssec-debug", false, "Disable upstream checking for DNSSEC debugging.")
	rootCmd.Flags().Bool("dnssec-no-timecheck", false, "Don't check DNSSEC signature timestamps until first cache-reload")
	rootCmd.Flags().String("dnssec-timestamp", "", "Timestamp file to verify system clock for DNSSEC")
	rootCmd.Flags().StringP("domain", "s", "", "Specify the domain to be assigned in DHCP leases.")
	rootCmd.Flags().BoolP("domain-needed", "D", false, "Do NOT forward queries with no domain part.")
	rootCmd.Flags().String("dumpfile", "", "Path to debug packet dump file")
	rootCmd.Flags().String("dumpmask", "", "Mask which packets to dump")
	rootCmd.Flags().StringP("edns-packet-max", "P", "", "Maximum supported UDP packet size for EDNS.0 (defaults to 4096).")
	rootCmd.Flags().StringP("enable-dbus", "1", "", "Enable the DBus interface for setting upstream servers, etc.")
	rootCmd.Flags().Bool("enable-ra", false, "Send router-advertisements for interfaces doing DHCPv6")
	rootCmd.Flags().String("enable-tftp", "", "Enable integrated read-only TFTP server.")
	rootCmd.Flags().String("enable-ubus", "", "Enable the UBus interface.")
	rootCmd.Flags().StringP("except-interface", "I", "", "Specify interface(s) NOT to listen on.")
	rootCmd.Flags().BoolP("expand-hosts", "E", false, "Expand simple names in /etc/hosts with domain-suffix.")
	rootCmd.Flags().BoolP("filterwin2k", "f", false, "Don't forward spurious DNS requests from Windows hosts.")
	rootCmd.Flags().StringP("group", "g", "", "Change to this group after startup (defaults to dip).")
	rootCmd.Flags().BoolP("help", "w", false, "Display this message. Use --help dhcp or --help dhcp6 for known DHCP options.")
	rootCmd.Flags().String("host-record", "", "Specify host (A/AAAA and PTR) records")
	rootCmd.Flags().String("hostsdir", "", "Read hosts files from a directory.")
	rootCmd.Flags().String("ignore-address", "", "Ignore DNS responses containing ipaddr.")
	rootCmd.Flags().StringP("interface", "i", "", "Specify interface(s) to listen on.")
	rootCmd.Flags().String("interface-name", "", "Give DNS name to IPv4 address of interface.")
	rootCmd.Flags().String("ipset", "", "Specify ipsets to which matching domains should be added")
	rootCmd.Flags().BoolP("keep-in-foreground", "k", false, "Do NOT fork into the background, do NOT run in debug mode.")
	rootCmd.Flags().BoolP("leasefile-ro", "9", false, "Do not use leasefile.")
	rootCmd.Flags().StringP("listen-address", "a", "", "Specify local address(es) to listen on.")
	rootCmd.Flags().String("local", "", "Never forward queries to specified domains.")
	rootCmd.Flags().Bool("local-service", false, "Accept queries only from directly-connected networks.")
	rootCmd.Flags().StringP("local-ttl", "T", "", "Specify time-to-live in seconds for replies from /etc/hosts.")
	rootCmd.Flags().BoolP("localise-queries", "y", false, "Answer DNS queries based on the interface a query was sent to.")
	rootCmd.Flags().BoolP("localmx", "L", false, "Return MX records for local hosts.")
	rootCmd.Flags().String("log-async", "", "Enable async. logging; optionally set queue length.")
	rootCmd.Flags().Bool("log-debug", false, "Log debugging information.")
	rootCmd.Flags().Bool("log-dhcp", false, "Extra logging for DHCP.")
	rootCmd.Flags().StringP("log-facility", "8", "", "Log to this syslog facility or file. (defaults to DAEMON)")
	rootCmd.Flags().BoolP("log-queries", "q", false, "Log DNS queries.")
	rootCmd.Flags().String("max-cache-ttl", "", "Specify time-to-live ceiling for cache.")
	rootCmd.Flags().String("max-port", "", "Specify highest port available for DNS query transmission.")
	rootCmd.Flags().String("max-ttl", "", "Specify time-to-live in seconds for maximum TTL to send to clients.")
	rootCmd.Flags().String("min-cache-ttl", "", "Specify time-to-live floor for cache.")
	rootCmd.Flags().String("min-port", "", "Specify lowest port available for DNS query transmission.")
	rootCmd.Flags().StringP("mx-host", "m", "", "Specify an MX record.")
	rootCmd.Flags().StringP("mx-target", "t", "", "Specify default target in an MX record.")
	rootCmd.Flags().String("naptr-record", "", "Specify NAPTR DNS record.")
	rootCmd.Flags().String("neg-ttl", "", "Specify time-to-live in seconds for negative caching.")
	rootCmd.Flags().BoolP("no-daemon", "d", false, "Do NOT fork into the background: run in debug mode.")
	rootCmd.Flags().StringP("no-dhcp-interface", "2", "", "Do not provide DHCP on this interface, only provide DNS.")
	rootCmd.Flags().BoolP("no-hosts", "h", false, "Do NOT load /etc/hosts file.")
	rootCmd.Flags().BoolP("no-negcache", "N", false, "Do NOT cache failed search results.")
	rootCmd.Flags().BoolP("no-ping", "5", false, "Disable ICMP echo address checking in the DHCP server.")
	rootCmd.Flags().BoolP("no-poll", "n", false, "Do NOT poll /etc/resolv.conf file, reload only on SIGHUP.")
	rootCmd.Flags().BoolP("no-resolv", "R", false, "Do NOT read resolv.conf.")
	rootCmd.Flags().StringP("pid-file", "x", "", "Specify path of PID file (defaults to /var/run/dnsmasq.pid).")
	rootCmd.Flags().StringP("port", "p", "", "Specify port to listen for DNS requests on (defaults to 53).")
	rootCmd.Flags().Bool("proxy-dnssec", false, "Proxy DNSSEC validation results from upstream nameservers.")
	rootCmd.Flags().String("ptr-record", "", "Specify PTR DNS record.")
	rootCmd.Flags().String("pxe-prompt", "", "Prompt to send to PXE clients.")
	rootCmd.Flags().String("pxe-service", "", "Boot service for PXE menu.")
	rootCmd.Flags().StringP("query-port", "Q", "", "Force the originating port for upstream DNS queries.")
	rootCmd.Flags().Bool("quiet-dhcp", false, "Do not log routine DHCP.")
	rootCmd.Flags().Bool("quiet-dhcp6", false, "Do not log routine DHCPv6.")
	rootCmd.Flags().Bool("quiet-ra", false, "Do not log RA.")
	rootCmd.Flags().BoolP("read-ethers", "Z", false, "Read DHCP static host information from /etc/ethers.")
	rootCmd.Flags().String("rebind-domain-ok", "", "Inhibit DNS-rebind protection on this domain.")
	rootCmd.Flags().Bool("rebind-localhost-ok", false, "Allow rebinding of 127.0.0.0/8, for RBL servers.")
	rootCmd.Flags().StringP("resolv-file", "r", "", "Specify path to resolv.conf (defaults to /etc/resolv.conf).")
	rootCmd.Flags().String("rev-server", "", "Specify address of upstream servers for reverse address queries")
	rootCmd.Flags().Bool("script-arp", false, "Call dhcp-script with changes to local ARP table.")
	rootCmd.Flags().Bool("script-on-renewal", false, "Call dhcp-script when lease expiry changes.")
	rootCmd.Flags().BoolP("selfmx", "e", false, "Return self-pointing MX records for local hosts.")
	rootCmd.Flags().StringP("server", "S", "", "Specify address(es) of upstream servers with optional domains.")
	rootCmd.Flags().String("servers-file", "", "Specify path to file with server= options")
	rootCmd.Flags().String("shared-network", "", "Specify extra networks sharing a broadcast domain for DHCP")
	rootCmd.Flags().StringP("srv-host", "W", "", "Specify a SRV record.")
	rootCmd.Flags().Bool("stop-dns-rebind", false, "Stop DNS rebinding. Filter private IP ranges when resolving.")
	rootCmd.Flags().BoolP("strict-order", "o", false, "Use nameservers strictly in the order given in /etc/resolv.conf.")
	rootCmd.Flags().String("synth-domain", "", "Specify a domain and address range for synthesised names")
	rootCmd.Flags().String("tag-if", "", "Evaluate conditional tag expression.")
	rootCmd.Flags().Bool("test", false, "Check configuration syntax.")
	rootCmd.Flags().Bool("tftp-lowercase", false, "Convert TFTP filenames to lowercase")
	rootCmd.Flags().String("tftp-max", "", "Maximum number of concurrent TFTP transfers (defaults to 50).")
	rootCmd.Flags().String("tftp-mtu", "", "Maximum MTU to use for TFTP transfers.")
	rootCmd.Flags().Bool("tftp-no-blocksize", false, "Disable the TFTP blocksize extension.")
	rootCmd.Flags().Bool("tftp-no-fail", false, "Do not terminate the service if TFTP directories are inaccessible.")
	rootCmd.Flags().String("tftp-port-range", "", "Ephemeral port range for use by TFTP transfers.")
	rootCmd.Flags().String("tftp-root", "", "Export files by TFTP only from the specified subtree.")
	rootCmd.Flags().Bool("tftp-secure", false, "Allow access only to files owned by the user running dnsmasq.")
	rootCmd.Flags().Bool("tftp-single-port", false, "Use only one port for TFTP server.")
	rootCmd.Flags().String("tftp-unique-root", "", "Add client IP or hardware address to tftp-root.")
	rootCmd.Flags().String("trust-anchor", "", "Specify trust anchor key digest.")
	rootCmd.Flags().StringP("txt-record", "Y", "", "Specify TXT DNS record.")
	rootCmd.Flags().StringP("user", "u", "", "Change to this user after startup. (defaults to nobody).")
	rootCmd.Flags().BoolP("version", "v", false, "Display dnsmasq version and copyright information.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"addn-hosts":       carapace.ActionFiles(),
		"conf-dir":         carapace.ActionDirectories(),
		"conf-file":        carapace.ActionFiles(),
		"dhcp-hostsdir":    carapace.ActionDirectories(),
		"dhcp-hostsfile":   carapace.ActionFiles(),
		"dhcp-leasefile":   carapace.ActionDirectories(),
		"dhcp-luascript":   carapace.ActionDirectories(),
		"dhcp-optsdir":     carapace.ActionDirectories(),
		"dhcp-optsfile":    carapace.ActionFiles(),
		"dhcp-script":      carapace.ActionDirectories(),
		"dhcp-scriptuser":  os.ActionUsers(),
		"dumpfile":         carapace.ActionFiles(),
		"except-interface": net.ActionDevices(net.AllDevices),
		"group":            os.ActionGroups(),
		"hostsdir":         carapace.ActionDirectories(),
		"interface":        net.ActionDevices(net.AllDevices),
		"max-port":         net.ActionPorts(),
		"min-port":         net.ActionPorts(),
		"pid-file":         carapace.ActionFiles(),
		"port":             net.ActionPorts(),
		"query-port":       net.ActionPorts(),
		"resolv-file":      carapace.ActionFiles(),
		"user":             os.ActionUsers(),
	})
}
