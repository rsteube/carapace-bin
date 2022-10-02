package action

import (
	"encoding/xml"
	"os"

	"github.com/rsteube/carapace"
)

type project struct {
	XMLName xml.Name `xml:"project"`
	Target  []struct {
		Name string `xml:"name,attr"`
	} `xml:"target"`
}

func ActionTargets(path string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path == "" {
			path = "build.xml"
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var p project
		if err := xml.Unmarshal(content, &p); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, target := range p.Target {
			vals = append(vals, target.Name)
		}
		return carapace.ActionValues(vals...)
	})
}
