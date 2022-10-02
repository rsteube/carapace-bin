package action

import (
	"encoding/json"
	"os"
	"time"

	"github.com/rsteube/carapace"
)

type sample struct {
	Id          string
	Description string
}

func ActionSamples() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		tmpfile, err := os.CreateTemp(os.TempDir(), "carapace-flutter-samples.json")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		err = tmpfile.Close() // just using ioutil.Tempfile for the random name
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		err = os.Remove(tmpfile.Name()) // will be written by flutter and must not exist
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		defer os.Remove(tmpfile.Name()) // remove after parsing content
		return carapace.ActionExecCommand("flutter", "create", "--suppress-analytics", "--list-samples", tmpfile.Name())(func(output []byte) carapace.Action {
			content, err := os.ReadFile(tmpfile.Name())
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			var samples []sample
			if err := json.Unmarshal(content, &samples); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, sample := range samples {
				vals = append(vals, sample.Id, sample.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Invoke(c).ToA() // ensure this is executed before `defer os.Remove`
	}).Cache(1 * time.Hour)
}
