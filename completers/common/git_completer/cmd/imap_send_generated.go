package cmd

import (
	"github.com/spf13/cobra"
)

var imap_sendCmd = &cobra.Command{
	Use:   "imap-send",
	Short: "Send a collection of patches from stdin to an IMAP folder",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	imap_sendCmd.Flags().Bool("curl", false, "use libcurl to communicate with the IMAP server")
	imap_sendCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	imap_sendCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	rootCmd.AddCommand(imap_sendCmd)
}
