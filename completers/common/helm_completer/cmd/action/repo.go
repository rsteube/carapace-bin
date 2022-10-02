package action

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
	"gopkg.in/yaml.v3"
)

type repositoriesFile struct {
	Repositories []struct {
		Name string
		Url  string
	}
}

func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		dir, err := os.UserConfigDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		content, err := os.ReadFile(dir + "/helm/repositories.yaml")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var repositoriesFile repositoriesFile
		err = yaml.Unmarshal(content, &repositoriesFile)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vars := make([]string, 0)
		for _, repo := range repositoriesFile.Repositories {
			vars = append(vars, repo.Name, repo.Url)
		}
		return carapace.ActionValuesDescribed(vars...)
	})
}

type index struct {
	Entries map[string][]struct {
		Name        string
		Description string
		Version     string
	}
}

func loadIndex(repo string) (index *index, err error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(fmt.Sprintf("%v/helm/repository/%v-index.yaml", dir, repo))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &index)
	return index, err
}

func ActionCharts(repo string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		index, err := loadIndex(repo)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		uniqueCharts := make(map[string]string)
		for _, charts := range index.Entries {
			for _, chart := range charts {
				uniqueCharts[chart.Name] = chart.Description
			}
		}

		vals := make([]string, 0)
		for key, value := range uniqueCharts {
			vals = append(vals, key, value)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionChartVersions(repo string, chart string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		index, err := loadIndex(repo)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, charts := range index.Entries {
			for _, c := range charts {
				if c.Name == chart {
					vals = append(vals, c.Version)
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionRepositoryCharts() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionRepositories().Invoke(c).Suffix("/").ToA()
		case 1:
			return ActionCharts(c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}
