package cmd

import (
	"github.com/rsteube/carapace-bin/completers/xz_completer/cmd"
)

/**
Description for go:generate
	Use: "unlzma",
	Short: "Compress or decompress .xz and .lzma files",
	Long: "https://linux.die.net/man/1/xz",
*/

func Execute() error {
	return cmd.ExecuteAlias("unlzma")
}
