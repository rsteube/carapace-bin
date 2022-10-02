package cmd

import (
	"os"
	"path/filepath"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "caragen",
	Short: "internal tool for go:generate",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd)
}

func rootDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path, err := util.FindReverse(wd, "go.mod")
	if err != nil {
		return "", err
	}

	return filepath.Dir(path), nil
}
