package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pass_completer/cmd/action"
	"github.com/spf13/cobra"
)

var otp_insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Prompt for and insert a new OTP secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otp_insertCmd).Standalone()
	otp_insertCmd.Flags().StringP("account", "a", "", "specify account")
	otp_insertCmd.Flags().BoolP("echo", "e", false, "echo input")
	otp_insertCmd.Flags().BoolP("force", "f", false, "overwrite existing")
	otp_insertCmd.Flags().StringP("issuer", "i", "", "specify issuer")
	otp_insertCmd.Flags().BoolP("secret", "s", false, "prompt for the secret value")

	otpCmd.AddCommand(otp_insertCmd)

	carapace.Gen(otp_insertCmd).PositionalCompletion(
		action.ActionPassNames(),
	)
}
