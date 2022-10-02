package cmd

import (
	"github.com/spf13/cobra"
)

var checkout_indexCmd = &cobra.Command{
	Use:   "checkout-index",
	Short: "Copy files from the index to the working tree",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	checkout_indexCmd.Flags().BoolP("all", "a", false, "check out all files in the index")
	checkout_indexCmd.Flags().BoolP("force", "f", false, "force overwrite of existing files")
	checkout_indexCmd.Flags().BoolP("index", "u", false, "update stat information in the index file")
	checkout_indexCmd.Flags().BoolP("no-create", "n", false, "don't checkout new files")
	checkout_indexCmd.Flags().String("prefix", "", "when creating files, prepend <string>")
	checkout_indexCmd.Flags().BoolP("quiet", "q", false, "no warning for existing files and files not in index")
	checkout_indexCmd.Flags().String("stage", "", "copy out the files from named stage")
	checkout_indexCmd.Flags().Bool("stdin", false, "read list of paths from the standard input")
	checkout_indexCmd.Flags().Bool("temp", false, "write the content to temporary files")
	checkout_indexCmd.Flags().BoolS("z", "z", false, "paths are separated with NUL character")
	rootCmd.AddCommand(checkout_indexCmd)
}
