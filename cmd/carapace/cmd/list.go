package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "--list",
	Short: "list completers",
	Run: func(cmd *cobra.Command, args []string) {
		switch cmd.Flag("format").Value.String() {
		case "json":
			printCompletersJson()
		default:
			printCompleters()
		}
	},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("format", "plain", "output format")

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json").StyleF(func(s string, sc style.Context) string {
			return style.ForPathExt("."+s, sc)
		}),
	})

}
func printCompleters() {
	maxlen := 0
	for _, name := range completers.Names() {
		if len := len(name); len > maxlen {
			maxlen = len
		}
	}

	for _, name := range completers.Names() {
		fmt.Printf("%-"+strconv.Itoa(maxlen)+"v %v\n", name, completers.Description(name))
	}
}

func printCompletersJson() {
	// TODO move to completers package
	type _completer struct {
		Name        string
		Description string
		Spec        string `json:",omitempty"`
		Overlay     string `json:",omitempty"`
	}

	_completers := make([]_completer, 0)
	for _, name := range completers.Names() {
		specPath, _ := completers.SpecPath(name)       // TODO handle error (log?)
		overlayPath, _ := completers.OverlayPath(name) // TODO handle error (log?)
		_completers = append(_completers, _completer{
			Name:        name,
			Description: completers.Description(name),
			Spec:        specPath,
			Overlay:     overlayPath,
		})
	}
	if m, err := json.Marshal(_completers); err == nil { // TODO handle error (log?)
		fmt.Println(string(m))
	}
}
