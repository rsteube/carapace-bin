package action

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

type device struct {
	Id             string
	Name           string
	TargetPlatform string
	Sdk            string
}

func ActionDevices() carapace.Action {
	return carapace.ActionExecCommand("flutter", "--suppress-analytics", "devices", "--machine")(func(output []byte) carapace.Action {
		var devices []device
		if err := json.Unmarshal(output, &devices); err != nil {
			if strings.Contains(string(output), "Waiting for another") {
				return carapace.ActionMessage(string(output))
			}
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, device := range devices {
			vals = append(vals, device.Id, fmt.Sprintf("%v • %v • %v", device.Name, device.TargetPlatform, device.Sdk))
		}
		return carapace.ActionValuesDescribed(vals...)
	})

}
