package cmd

import (
	"github.com/spf13/cobra"
)

var instawebCmd = &cobra.Command{
	Use:   "instaweb",
	Short: "Instantly browse your working repository in gitweb",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	instawebCmd.Flags().BoolP("browser", "b", false, "...     the browser to launch")
	instawebCmd.Flags().BoolP("httpd", "d", false, "...       the command to launch")
	instawebCmd.Flags().BoolP("local", "l", false, "only bind on 127.0.0.1")
	instawebCmd.Flags().BoolP("module-path", "m", false, "...    the module path (only needed for apache2)")
	instawebCmd.Flags().BoolP("port", "p", false, "...        the port to bind to")
	instawebCmd.Flags().Bool("restart", false, "restart the web server")
	instawebCmd.Flags().Bool("start", false, "start the web server")
	instawebCmd.Flags().Bool("stop", false, "stop the web server")
	rootCmd.AddCommand(instawebCmd)
}
