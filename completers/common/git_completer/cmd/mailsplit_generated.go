package cmd

import (
	"github.com/spf13/cobra"
)

var mailsplitCmd = &cobra.Command{
	Use:   "mailsplit",
	Short: "Simple UNIX mbox splitter program",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(mailsplitCmd)
}
