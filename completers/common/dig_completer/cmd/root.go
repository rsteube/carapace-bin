package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dig",
	Short: "DNS lookup utility",
	Long:  "https://linux.die.net/man/1/dig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "This option indicates that only IPv4 should be used.")
	rootCmd.Flags().BoolS("6", "6", false, "This option indicates that only IPv6 should be used.")
	rootCmd.Flags().StringS("b", "b", "", "This option sets the source IP address of the query.")
	rootCmd.Flags().StringS("c", "c", "", "This  option  sets the query class.")
	rootCmd.Flags().StringS("f", "f", "", "This option sets batch mode.")
	rootCmd.Flags().StringS("k", "k", "", "This option tells named to sign queries using TSIG using a key read  from  the  given file.")
	rootCmd.Flags().BoolS("m", "m", false, "This option enables memory usage debugging.")
	rootCmd.Flags().StringS("p", "p", "", "This option sends the query to a non-standard port on the server.")
	rootCmd.Flags().StringS("q", "q", "", "This option specifies the domain name to query.")
	rootCmd.Flags().BoolS("r", "r", false, "This  option  indicates  that options from ${HOME}/.digrc should not be read.")
	rootCmd.Flags().StringS("t", "t", "", "This option indicates the resource record type to query.")
	rootCmd.Flags().BoolS("u", "u", false, "This option indicates that print query times should be provided in  microseconds  instead of milliseconds.")
	rootCmd.Flags().BoolS("v", "v", false, "This option prints the version number and exits.")
	rootCmd.Flags().StringS("x", "x", "", "This option sets simplified reverse lookups, for mapping addresses to names.")
	rootCmd.Flags().StringS("y", "y", "", "This option signs queries using TSIG with the given authentication key.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionValuesDescribed(
			"IN", "",
			"HS", "Hesiod records",
			"CH", "Chaosnet records",
		),
		"f": carapace.ActionFiles(),
		"k": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ActionOptions(),
	)
}

func ActionOptions() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"+aaflag", "This option is a synonym for +[no]aaonly.",
				"+aaonly", "This option sets the aa flag in the query.",
				"+additional", "This option displays [or does not display] the additional section of a reply.",
				"+adflag", "This option sets [or does not set] the AD (authentic data) bit in the query.",
				"+all", "This option sets or clears all display flags.",
				"+answer", "This option displays [or does not display] the answer section of a reply.",
				"+authority", "This option displays [or does not display] the authority section of a reply.",
				"+badcookie", "This option retries the lookup with a new server cookie if a BADCOOKIE response is received.",
				"+besteffort", "This option attempts to display the contents of messages which are malformed.",
				"+bufsize=", "This option sets the UDP message buffer size advertised using EDNS0 to B bytes.",
				"+bufsize=", "This option sets the UDP message buffer size advertised using EDNS0 to B bytes.",
				"+cdflag", "This option sets [or does not set] the CD (checking disabled) bit in the query.",
				"+class", "This option displays [or does not display] the CLASS when printing the record.",
				"+cmd", "This  option  toggles the printing of the initial comment in the output.",
				"+comments", "This option toggles the display of some comment lines in the output.",
				"+cookie=", "This option sends [or does not send] a COOKIE EDNS option.",
				"+crypto", "This  option  toggles the display of cryptographic fields in DNSSEC records.",
				"+defname", "This option, which is deprecated, is treated as a synonym for +[no]search.",
				"+dnssec", "This option requests that DNSSEC records be sent by setting the DNSSEC OK (DO) bit in the OPT record.",
				"+domain=", "This  option  sets the search list to contain the single domain somename.",
				"+domain=", "This  option  sets the search list to contain the single domain somename.",
				"+dscp=", "This option sets the DSCP code point to be used when sending the query.",
				"+dscp=", "This option sets the DSCP code point to be used when sending the query.",
				"+edns=", "This option specifies the EDNS version to query with. Valid values are 0 to 255.",
				"+ednsflags=", "This  option sets the must-be-zero EDNS flags bits (Z bits) to the specified value.",
				"+ednsnegotiation", "This option enables/disables EDNS version negotiation.",
				"+ednsopt=", "This option specifies the EDNS option with code point code and an optional payload of value as a hexadecimal string.",
				"+expire", "This option sends an EDNS Expire option.",
				"+fail", "This  option  indicates  that named should try [or not try] the next server if a SERVFAIL is received.",
				"+header-only", "This option sends a query with a DNS header without a question section.",
				"+identify", "This option shows [or does not show] the IP address and port number that supplied the answer.",
				"+idnin", "This option processes [or does not process] IDN domain names on input.",
				"+idnout", "This option converts [or does not convert] puny code on output.",
				"+ignore", "This option ignores [or does not ignore] truncation in UDP responses instead of retrying with TCP.",
				"+keepalive", "This option sends [or does not send] an EDNS Keepalive option.",
				"+keepopen", "This option keeps [or does not keep] the TCP socket open between queries.",
				"+mapped", "This option allows [or does not allow] mapped IPv4-over-IPv6 addresses to be used.",
				"+multiline", "This  option  prints [or does not print] records.",
				"+ndots=", "This option sets the number of dots (D) that must appear in name for it to be considered absolute.",
				"+ndots=", "This option sets the number of dots (D) that must appear in name for it to be considered absolute.",
				"+noaaflag", "This option is a synonym for +[no]aaonly.",
				"+noaaonly", "This option sets the aa flag in the query.",
				"+noadditional", "This option displays [or does not display] the additional section of a reply.",
				"+noadflag", "This option sets [or does not set] the AD (authentic data) bit in the query.",
				"+noall", "This option sets or clears all display flags.",
				"+noanswer", "This option displays [or does not display] the answer section of a reply.",
				"+noauthority", "This option displays [or does not display] the authority section of a reply.",
				"+nobadcookie", "This option retries the lookup with a new server cookie if a BADCOOKIE response is received.",
				"+nobesteffort", "This option attempts to display the contents of messages which are malformed.",
				"+nocdflag", "This option sets [or does not set] the CD (checking disabled) bit in the query.",
				"+noclass", "This option displays [or does not display] the CLASS when printing the record.",
				"+nocmd", "This  option  toggles the printing of the initial comment in the output.",
				"+nocomments", "This option toggles the display of some comment lines in the output.",
				"+nocookie=", "This option sends [or does not send] a COOKIE EDNS option.",
				"+nocrypto", "This  option  toggles the display of cryptographic fields in DNSSEC records.",
				"+nodefname", "This option, which is deprecated, is treated as a synonym for +[no]search.",
				"+nodnssec", "This option requests that DNSSEC records be sent by setting the DNSSEC OK (DO) bit in the OPT record.",
				"+noedns=", "This option specifies the EDNS version to query with. Valid values are 0 to 255.",
				"+noednsflags=", "This  option sets the must-be-zero EDNS flags bits (Z bits) to the specified value.",
				"+noednsnegotiation", "This option enables/disables EDNS version negotiation.",
				"+noednsopt=", "This option specifies the EDNS option with code point code and an optional payload of value as a hexadecimal string.",
				"+noexpire", "This option sends an EDNS Expire option.",
				"+nofail", "This  option  indicates  that named should try [or not try] the next server if a SERVFAIL is received.",
				"+noheader-only", "This option sends a query with a DNS header without a question section.",
				"+noidentify", "This option shows [or does not show] the IP address and port number that supplied the answer.",
				"+noidnin", "This option processes [or does not process] IDN domain names on input.",
				"+noidnout", "This option converts [or does not convert] puny code on output.",
				"+noignore", "This option ignores [or does not ignore] truncation in UDP responses instead of retrying with TCP.",
				"+nokeepalive", "This option sends [or does not send] an EDNS Keepalive option.",
				"+nokeepopen", "This option keeps [or does not keep] the TCP socket open between queries.",
				"+nomapped", "This option allows [or does not allow] mapped IPv4-over-IPv6 addresses to be used.",
				"+nomultiline", "This  option  prints [or does not print] records.",
				"+nonsid", "When enabled, this option includes an EDNS name server ID request when sending a query.",
				"+nonssearch", "When this option is set, dig attempts to find the authoritative name servers.",
				"+noonesoa", "When enabled, this option prints only one (starting) SOA record when performing an AXFR.",
				"+noopcode=", "When enabled, this option sets (restores) the DNS message opcode to the specified value.",
				"+noqr", "This option toggles the display of the query message as it is sent.",
				"+noquestion", "This option toggles the display of the question section of a query when an answer is returned.",
				"+noraflag", "This option sets [or does not set] the RA (Recursion Available) bit in the query.",
				"+nordflag", "This option is a synonym for +[no]recurse.",
				"+norecurse", "This option toggles the setting of the RD (recursion desired) bit in the query.",
				"+norrcomments", "This option toggles the display of per-record comments in the output.",
				"+nosearch", "This option uses [or does not use] the search list defined by the searchlist or domain directive in resolv.conf, if any.",
				"+noshort", "This  option toggles whether a terse answer is provided.",
				"+noshowsearch", "This option performs [or does not perform] a search showing intermediate results.",
				"+nosigchase", "This feature is now obsolete and has been removed; use delv instead.",
				"+nostats", "This  option  toggles  the  printing  of  statistics.",
				"+nosubnet=", "This option sends [or does not send] an EDNS CLIENT-SUBNET option with the specified IP address or network prefix.",
				"+notcflag", "This option sets [or does not set] the TC (TrunCation) bit in the query.",
				"+notcp", "This  option  uses [or does not use] TCP when querying name servers.",
				"+notopdown", "This feature is related to dig +sigchase, which is obsolete and has been removed.",
				"+notrace", "This option toggles tracing of the delegation path from the root name servers for the name being looked up.",
				"+nottlid", "This option displays [or does not display] the TTL when printing the record.",
				"+nottlunits", "This  option  displays  [or  does  not  display]  the TTL in friendly human-readable time units",
				"+nounexpected", "This option accepts [or does not accept] answers from unexpected sources.",
				"+nounknownformat", "This option prints all RDATA in unknown RR type presentation format (RFC 3597).",
				"+novc", "This option uses [or does not use] TCP when querying name servers.",
				"+noyaml", "When enabled, this option prints the responses in a detailed YAML format.",
				"+nozflag", "This option sets [or does not set] the last unassigned DNS header flag in a DNS query.",
				"+nsid", "When enabled, this option includes an EDNS name server ID request when sending a query.",
				"+nssearch", "When this option is set, dig attempts to find the authoritative name servers.",
				"+onesoa", "When enabled, this option prints only one (starting) SOA record when performing an AXFR.",
				"+opcode=", "When enabled, this option sets (restores) the DNS message opcode to the specified value.",
				"+padding=", "This  option  pads the size of the query packet using the EDNS Padding option to blocks of value bytes.",
				"+padding=", "This  option  pads the size of the query packet using the EDNS Padding option to blocks of value bytes.",
				"+qr", "This option toggles the display of the query message as it is sent.",
				"+question", "This option toggles the display of the question section of a query when an answer is returned.",
				"+raflag", "This option sets [or does not set] the RA (Recursion Available) bit in the query.",
				"+rdflag", "This option is a synonym for +[no]recurse.",
				"+recurse", "This option toggles the setting of the RD (recursion desired) bit in the query.",
				"+retry=", "This option sets the number of times to retry UDP queries to server to T instead of the default, 2.",
				"+retry=", "This option sets the number of times to retry UDP queries to server to T instead of the default, 2.",
				"+rrcomments", "This option toggles the display of per-record comments in the output.",
				"+search", "This option uses [or does not use] the search list defined by the searchlist or domain directive in resolv.conf, if any.",
				"+short", "This  option toggles whether a terse answer is provided.",
				"+showsearch", "This option performs [or does not perform] a search showing intermediate results.",
				"+sigchase", "This feature is now obsolete and has been removed; use delv instead.",
				"+split=", "This option splits long hex- or base64-formatted fields in resource records into chunks of W characters.",
				"+split=", "This option splits long hex- or base64-formatted fields in resource records into chunks of W characters.",
				"+stats", "This  option  toggles  the  printing  of  statistics.",
				"+subnet=", "This option sends [or does not send] an EDNS CLIENT-SUBNET option with the specified IP address or network prefix.",
				"+tcflag", "This option sets [or does not set] the TC (TrunCation) bit in the query.",
				"+tcp", "This  option  uses [or does not use] TCP when querying name servers.",
				"+timeout=", "This option sets the timeout for a query to T seconds.",
				"+timeout=", "This option sets the timeout for a query to T seconds.",
				"+topdown", "This feature is related to dig +sigchase, which is obsolete and has been removed.",
				"+trace", "This option toggles tracing of the delegation path from the root name servers for the name being looked up.",
				"+tries=", "This option sets the number of times to try UDP queries to server to T instead of the default, 3.",
				"+tries=", "This option sets the number of times to try UDP queries to server to T instead of the default, 3.",
				"+trusted-key=", "This option formerly specified trusted keys for use with dig +sigchase.",
				"+trusted-key=", "This option formerly specified trusted keys for use with dig +sigchase.",
				"+ttlid", "This option displays [or does not display] the TTL when printing the record.",
				"+ttlunits", "This  option  displays  [or  does  not  display]  the TTL in friendly human-readable time units",
				"+unexpected", "This option accepts [or does not accept] answers from unexpected sources.",
				"+unknownformat", "This option prints all RDATA in unknown RR type presentation format (RFC 3597).",
				"+vc", "This option uses [or does not use] TCP when querying name servers.",
				"+yaml", "When enabled, this option prints the responses in a detailed YAML format.",
				"+zflag", "This option sets [or does not set] the last unassigned DNS header flag in a DNS query.",
			)
		default:
			return carapace.ActionValues()
		}
	})
}
