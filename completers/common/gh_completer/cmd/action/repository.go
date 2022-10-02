package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

func ActionOwnerRepositories(cmd *cobra.Command) carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		// TODO hack to enable completion outside git repo - this needs to be fixed in GraphQlAction/repooverride though
		if cmd.Flag("repo") == nil {
			cmd.Flags().String("repo", "", "")
		}

		switch len(c.Parts) {
		case 0:
			return gh.ActionOwners(gh.HostOpts{}).Invoke(c).Suffix("/").ToA() // TODO host
		case 1:
			return gh.ActionRepositories(gh.OwnerOpts{Owner: c.Parts[0]})
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionRepositoryFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"name",
		"nameWithOwner",
		"owner",
		"parent",
		"templateRepository",
		"description",
		"homepageUrl",
		"openGraphImageUrl",
		"usesCustomOpenGraphImage",
		"url",
		"sshUrl",
		"mirrorUrl",
		"securityPolicyUrl",

		"createdAt",
		"pushedAt",
		"updatedAt",

		"isBlankIssuesEnabled",
		"isSecurityPolicyEnabled",
		"hasIssuesEnabled",
		"hasProjectsEnabled",
		"hasWikiEnabled",
		"mergeCommitAllowed",
		"squashMergeAllowed",
		"rebaseMergeAllowed",

		"forkCount",
		"stargazerCount",
		"watchers",
		"issues",
		"pullRequests",

		"codeOfConduct",
		"contactLinks",
		"defaultBranchRef",
		"deleteBranchOnMerge",
		"diskUsage",
		"fundingLinks",
		"isArchived",
		"isEmpty",
		"isFork",
		"isInOrganization",
		"isMirror",
		"isPrivate",
		"isTemplate",
		"isUserConfigurationRepository",
		"licenseInfo",
		"viewerCanAdminister",
		"viewerDefaultCommitEmail",
		"viewerDefaultMergeMethod",
		"viewerHasStarred",
		"viewerPermission",
		"viewerPossibleCommitEmails",
		"viewerSubscription",

		"repositoryTopics",
		"primaryLanguage",
		"languages",
		"issueTemplates",
		"pullRequestTemplates",
		"labels",
		"milestones",
		"latestRelease",

		"assignableUsers",
		"mentionableUsers",
		"projects",
	)
}

func ActionRepositorySearchFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"defaultBranch",
		"description",
		"forksCount",
		"fullName",
		"hasDownloads",
		"hasIssues",
		"hasPages",
		"hasProjects",
		"hasWiki",
		"homepage",
		"id",
		"isArchived",
		"isDisabled",
		"isFork",
		"isPrivate",
		"language",
		"license",
		"name",
		"openIssuesCount",
		"owner",
		"pushedAt",
		"size",
		"stargazersCount",
		"updatedAt",
		"visibility",
		"watchersCount",
	)
}
