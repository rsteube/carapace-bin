package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pass_completer/cmd/action"
	"github.com/spf13/cobra"
)

var otp_appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append an OTP secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otp_appendCmd).Standalone()
	otp_appendCmd.Flags().StringP("account", "a", "", "specify account")
	otp_appendCmd.Flags().BoolP("echo", "e", false, "echo input")
	otp_appendCmd.Flags().BoolP("force", "f", false, "overwrite existing")
	otp_appendCmd.Flags().StringP("issuer", "i", "", "specify issuer")
	otp_appendCmd.Flags().BoolP("secret", "s", false, "prompt for the secret value")

	otpCmd.AddCommand(otp_appendCmd)

	carapace.Gen(otp_appendCmd).PositionalCompletion(
		action.ActionPassNames(),
	)
}
