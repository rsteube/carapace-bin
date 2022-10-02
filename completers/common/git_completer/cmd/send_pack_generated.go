package cmd

import (
	"github.com/spf13/cobra"
)

var send_packCmd = &cobra.Command{
	Use:   "send-pack",
	Short: "Push objects over Git protocol to another repository",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	send_packCmd.Flags().Bool("all", false, "push all refs")
	send_packCmd.Flags().Bool("atomic", false, "request atomic transaction on remote side")
	send_packCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	send_packCmd.Flags().String("exec", "", "receive pack program")
	send_packCmd.Flags().BoolP("force", "f", false, "force updates")
	send_packCmd.Flags().String("force-with-lease", "", "require old value of ref to be at this value")
	send_packCmd.Flags().Bool("helper-status", false, "print status from remote helper")
	send_packCmd.Flags().Bool("mirror", false, "mirror all refs")
	send_packCmd.Flags().Bool("progress", false, "force progress reporting")
	send_packCmd.Flags().String("push-option", "", "option to transmit")
	send_packCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	send_packCmd.Flags().String("receive-pack", "", "receive pack program")
	send_packCmd.Flags().String("remote", "", "remote name")
	send_packCmd.Flags().String("signed", "", "GPG sign the push")
	send_packCmd.Flags().Bool("stateless-rpc", false, "use stateless RPC protocol")
	send_packCmd.Flags().Bool("stdin", false, "read refs from stdin")
	send_packCmd.Flags().Bool("thin", false, "use thin pack")
	send_packCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	rootCmd.AddCommand(send_packCmd)
}
