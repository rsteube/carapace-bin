package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var otpCmd = &cobra.Command{
	Use:   "otp",
	Short: "extension for managing one-time-password (OTP) tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otpCmd).Standalone()

	rootCmd.AddCommand(otpCmd)
}
