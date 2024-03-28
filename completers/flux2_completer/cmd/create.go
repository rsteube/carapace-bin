package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create or update sources and resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()
	createCmd.PersistentFlags().Bool("export", false, "export in YAML format to stdout")
	createCmd.PersistentFlags().String("interval", "", "source sync interval")
	createCmd.PersistentFlags().StringSlice("label", []string{}, "set labels on the resource (can specify multiple labels with commas: label1=value1,label2=value2)")
	rootCmd.AddCommand(createCmd)
}
