package cmd

import (
	"github.com/spf13/cobra"
)

var send_emailCmd = &cobra.Command{
	Use:   "send-email",
	Short: "Send a collection of patches as emails",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	send_emailCmd.Flags().Bool("8bit-encoding", false, "<str>  * Encoding to assume 8bit mails if undeclared")
	send_emailCmd.Flags().Bool("annotate", false, "* Review each patch that will be sent in an editor.")
	send_emailCmd.Flags().Bool("batch-size", false, "<int>  * send max <int> message per connection.")
	send_emailCmd.Flags().Bool("bcc", false, "<str>  * Email Bcc:")
	send_emailCmd.Flags().Bool("cc", false, "<str>  * Email Cc:")
	send_emailCmd.Flags().Bool("cc-cmd", false, "<str>  * Email Cc: via `<str> $patch_path`")
	send_emailCmd.Flags().Bool("cc-cover", false, "* Email Cc: addresses in the cover letter.")
	send_emailCmd.Flags().Bool("chain-reply-to", false, "* Chain In-Reply-To: fields. Default off.")
	send_emailCmd.Flags().Bool("compose", false, "* Open an editor for introduction.")
	send_emailCmd.Flags().Bool("compose-encoding", false, "<str>  * Encoding to assume for introduction.")
	send_emailCmd.Flags().Bool("confirm", false, "<str>  * Confirm recipients before sending;               auto, cc, compose, always, or never.")
	send_emailCmd.Flags().Bool("dry-run", false, "* Don't actually send the emails.")
	send_emailCmd.Flags().Bool("dump-aliases", false, "* Dump configured aliases and exit.")
	send_emailCmd.Flags().Bool("envelope-sender", false, "<str>  * Email envelope sender.")
	send_emailCmd.Flags().Bool("force", false, "* Send even if safety checks would prevent it.")
	send_emailCmd.Flags().Bool("format-patch", false, "* understand any non optional arguments as               `git format-patch` ones.")
	send_emailCmd.Flags().Bool("from", false, "<str>  * Email From:")
	send_emailCmd.Flags().Bool("identity", false, "<str>  * Use the sendemail.<id> options.")
	send_emailCmd.Flags().Bool("in-reply-to", false, "<str>  * Email \"In-Reply-To:\"")
	send_emailCmd.Flags().Bool("no-annotate", false, "* Review each patch that will be sent in an editor.")
	send_emailCmd.Flags().Bool("no-bcc", false, "<str>  * Email Bcc:")
	send_emailCmd.Flags().Bool("no-cc", false, "<str>  * Email Cc:")
	send_emailCmd.Flags().Bool("no-cc-cover", false, "* Email Cc: addresses in the cover letter.")
	send_emailCmd.Flags().Bool("no-chain-reply-to", false, "* Chain In-Reply-To: fields. Default off.")
	send_emailCmd.Flags().Bool("no-format-patch", false, "* understand any non optional arguments as               `git format-patch` ones.")
	send_emailCmd.Flags().Bool("no-signed-off-by-cc", false, "* Send to Signed-off-by: addresses. Default on.")
	send_emailCmd.Flags().Bool("no-smtp-auth", false, "Disable SMTP authentication. Shorthand for               `--smtp-auth=none`")
	send_emailCmd.Flags().Bool("no-suppress-from", false, "* Send to self. Default off.")
	send_emailCmd.Flags().Bool("no-thread", false, "* Use In-Reply-To: field. Default on.")
	send_emailCmd.Flags().Bool("no-to", false, "<str>  * Email To:")
	send_emailCmd.Flags().Bool("no-to-cover", false, "* Email To: addresses in the cover letter.")
	send_emailCmd.Flags().Bool("no-validate", false, "* Perform patch sanity checks. Default on.")
	send_emailCmd.Flags().Bool("no-xmailer", false, "* Add \"X-Mailer:\" header (default).")
	send_emailCmd.Flags().Bool("quiet", false, "* Output one line of info per email.")
	send_emailCmd.Flags().Bool("relogin-delay", false, "<int>  * delay <int> seconds between two successive login.               This option can only be used with --batch-size")
	send_emailCmd.Flags().Bool("reply-to", false, "<str>  * Email \"Reply-To:\"")
	send_emailCmd.Flags().Bool("signed-off-by-cc", false, "* Send to Signed-off-by: addresses. Default on.")
	send_emailCmd.Flags().Bool("smtp-auth", false, "<str>  * Space-separated list of allowed AUTH mechanisms, or               \"none\" to disable authentication.               This setting forces to use one of the listed mechanisms.")
	send_emailCmd.Flags().Bool("smtp-debug", false, "<0|1>  * Disable, enable Net::SMTP debug.")
	send_emailCmd.Flags().Bool("smtp-domain", false, "<str>  * The domain name sent to HELO/EHLO handshake")
	send_emailCmd.Flags().Bool("smtp-encryption", false, "<str>  * tls or ssl; anything else disables.")
	send_emailCmd.Flags().Bool("smtp-pass", false, "<str>  * Password for SMTP-AUTH; not necessary.")
	send_emailCmd.Flags().Bool("smtp-server", false, "<str:int>  * Outgoing SMTP server to use. The port               is optional. Default 'localhost'.")
	send_emailCmd.Flags().Bool("smtp-server-option", false, "<str>  * Outgoing SMTP server option to use.")
	send_emailCmd.Flags().Bool("smtp-server-port", false, "<int>  * Outgoing SMTP server port.")
	send_emailCmd.Flags().Bool("smtp-ssl", false, "* Deprecated. Use '--smtp-encryption ssl'.")
	send_emailCmd.Flags().Bool("smtp-ssl-cert-path", false, "<str>  * Path to ca-certificates (either directory or file).               Pass an empty string to disable certificate               verification.")
	send_emailCmd.Flags().Bool("smtp-user", false, "<str>  * Username for SMTP-AUTH.")
	send_emailCmd.Flags().Bool("subject", false, "<str>  * Email \"Subject:\"")
	send_emailCmd.Flags().Bool("suppress-cc", false, "<str>  * author, self, sob, cc, cccmd, body, bodycc, misc-by, all.")
	send_emailCmd.Flags().Bool("suppress-from", false, "* Send to self. Default off.")
	send_emailCmd.Flags().Bool("thread", false, "* Use In-Reply-To: field. Default on.")
	send_emailCmd.Flags().Bool("to", false, "<str>  * Email To:")
	send_emailCmd.Flags().Bool("to-cmd", false, "<str>  * Email To: via `<str> $patch_path`")
	send_emailCmd.Flags().Bool("to-cover", false, "* Email To: addresses in the cover letter.")
	send_emailCmd.Flags().Bool("transfer-encoding", false, "<str>  * Transfer encoding to use (quoted-printable, 8bit, base64)")
	send_emailCmd.Flags().Bool("validate", false, "* Perform patch sanity checks. Default on.")
	send_emailCmd.Flags().Bool("xmailer", false, "* Add \"X-Mailer:\" header (default).")
	rootCmd.AddCommand(send_emailCmd)
}
