package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scp",
	Short: "OpenSSH secure file copy",
	Long:  "https://linux.die.net/man/1/scp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("3", "3", false, "Copies between two remote hosts are transferred through the local host.")
	rootCmd.Flags().BoolS("4", "4", false, "Forces scp to use IPv4 addresses only.")
	rootCmd.Flags().BoolS("6", "6", false, "Forces scp to use IPv6 addresses only.")
	rootCmd.Flags().BoolS("A", "A", false, "Allows forwarding of ssh-agent(1) to the remote system.")
	rootCmd.Flags().BoolS("B", "B", false, "Selects batch mode (prevents asking for passwords or passphrases).")
	rootCmd.Flags().BoolS("C", "C", false, "Compression enable.")
	rootCmd.Flags().StringS("F", "F", "", "Specifies an alternative per-user configuration file for ssh.")
	rootCmd.Flags().StringS("J", "J", "", "Connect to the target host by first making an scp connection to the jump host.")
	rootCmd.Flags().StringS("P", "P", "", "Specifies the port to connect to on the remote host.")
	rootCmd.Flags().StringS("S", "S", "", "Name of program to use for the encrypted connection.")
	rootCmd.Flags().BoolS("T", "T", false, "Disable strict filename checking.")
	rootCmd.Flags().StringS("c", "c", "", "Selects the cipher to use for encrypting the data transfer.")
	rootCmd.Flags().StringS("i", "i", "", "Selects the file from which the identity (private key) for public key authentication is read.")
	rootCmd.Flags().StringS("l", "l", "", "Limits the used bandwidth, specified in Kbit/s.")
	rootCmd.Flags().StringS("o", "o", "", "Can be used to pass options to ssh in the format used in ssh_config.")
	rootCmd.Flags().BoolS("p", "p", false, "Preserves modification times, access times, and modes from the original file.")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode.")
	rootCmd.Flags().BoolS("r", "r", false, "Recursively copy entire directories.")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"F": carapace.ActionFiles(),
		"J": net.ActionHosts(),
		"S": carapace.ActionFiles(),
		"c": ssh.ActionCiphers(),
		"i": carapace.ActionFiles(),
		"o": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ssh.ActionOptions()
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().Invoke(c).Merge(carapace.ActionFiles().Invoke(c)).ToA()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
