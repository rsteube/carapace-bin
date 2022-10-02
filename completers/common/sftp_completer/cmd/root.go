package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sftp",
	Short: "OpenSSH secure file transfer",
	Long:  "https://linux.die.net/man/1/sftp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "Forces sftp to use IPv4 addresses only.")
	rootCmd.Flags().BoolS("6", "6", false, "Forces sftp to use IPv6 addresses only.")
	rootCmd.Flags().BoolS("A", "A", false, "Allows forwarding of ssh-agent(1) to the remote system.")
	rootCmd.Flags().StringS("B", "B", "", "Specify the size of the buffer that sftp uses when transferring files.")
	rootCmd.Flags().BoolS("C", "C", false, "Enables compression (via ssh's -C flag).")
	rootCmd.Flags().StringS("D", "D", "", "Connect directly to a local sftp server (rather than via ssh(1)).")
	rootCmd.Flags().StringS("F", "F", "", "Specifies an alternative per-user configuration file for ssh(1).")
	rootCmd.Flags().StringS("J", "J", "", "Connect to the target host by first making an sftp connection to the jump host described by destination.")
	rootCmd.Flags().BoolS("N", "N", false, "Disables quiet mode, e.g. to override the implicit quiet mode set by the -b flag.")
	rootCmd.Flags().StringS("P", "P", "", "Specifies the port to connect to on the remote host.")
	rootCmd.Flags().StringS("R", "R", "", "Specify how many requests may be outstanding at any one time.")
	rootCmd.Flags().StringS("S", "S", "", "Name of the program to use for the encrypted connection.")
	rootCmd.Flags().BoolS("a", "a", false, "Attempt to continue interrupted transfers rather than overwriting existing partial or complete copies of files.")
	rootCmd.Flags().StringS("b", "b", "", "Batch mode reads a series of commands from an input batchfile instead of stdin.")
	rootCmd.Flags().StringS("c", "c", "", "Selects the cipher to use for encrypting the data transfers.")
	rootCmd.Flags().BoolS("f", "f", false, "Requests that files be flushed to disk immediately after transfer.")
	rootCmd.Flags().StringS("i", "i", "", "Selects the file from which the identity (private key) for public key authentication is read.")
	rootCmd.Flags().StringS("l", "l", "", "Limits the used bandwidth, specified in Kbit/s.")
	rootCmd.Flags().StringS("o", "o", "", "Can be used to pass options to ssh in the format used in ssh_config(5).")
	rootCmd.Flags().BoolS("p", "p", false, "Preserves modification times, access times, and modes from the original files transferred.")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode: disables the progress meter as well as warning and diagnostic messages from ssh(1).")
	rootCmd.Flags().BoolS("r", "r", false, "Recursively copy entire directories when uploading and downloading.")
	rootCmd.Flags().StringS("s", "s", "", "Specifies the SSH2 subsystem or the path for an sftp server on the remote host.")
	rootCmd.Flags().BoolS("v", "v", false, "Raise logging level.  This option is also passed to ssh.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D": carapace.ActionDirectories(),
		"F": carapace.ActionFiles(),
		"J": net.ActionHosts(),
		"b": carapace.ActionFiles(),
		"c": ssh.ActionCiphers(),
		"i": carapace.ActionFiles(),
		"o": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ssh.ActionOptions()
		}),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		net.ActionHosts(),
	)
}
