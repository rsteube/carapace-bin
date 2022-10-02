package cmd

import (
	"github.com/spf13/cobra"
)

var store_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available OpenFaaS functions in a store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	store_listCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output to see the full description of each function in the store")
	storeCmd.AddCommand(store_listCmd)
}
