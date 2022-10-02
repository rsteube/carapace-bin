package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var otp_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Test an URI string vor validity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otp_validateCmd).Standalone()
	otp_validateCmd.Flags().BoolP("help", "h", false, "show usage help")

	otpCmd.AddCommand(otp_validateCmd)
}
