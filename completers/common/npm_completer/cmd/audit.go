package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Run a security audit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditCmd).Standalone()
	auditCmd.Flags().String("audit-level", "", "mnimum level of vulnerability level to exit with non-zero exit code")
	auditCmd.Flags().Bool("dry-run", false, "only report what would be done")
	auditCmd.Flags().BoolP("force", "f", false, "remove various protections against unfortunate side effects")
	auditCmd.Flags().Bool("json", false, "output as json")
	auditCmd.Flags().String("omit", "", "omit dependency type")
	auditCmd.Flags().Bool("package-lock-only", false, "ignore node-modules")
	addWorkspaceFlags(auditCmd)

	carapace.Gen(auditCmd).FlagCompletion(carapace.ActionMap{
		"audit-level": carapace.ActionValues("info", "low", "moderate", "high", "critical", "none"),
		"omit":        carapace.ActionValues("dev", "optional", "peer"),
	})

	rootCmd.AddCommand(auditCmd)
}
