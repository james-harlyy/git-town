package cmdhelpers

import (
	"github.com/git-town/git-town/v22/internal/config"
)

// BranchNameWithPrefix applies the configured GitHub username prefix to the given branch name.
// If a GitHub username is configured, it prepends it with a "/" separator.
// If no GitHub username is configured, it returns the original branch name unchanged.
func BranchNameWithPrefix(baseBranchName string, config config.UnvalidatedConfig) string {
	if githubUsername, hasGitHubUsername := config.NormalConfig.GitHubUsername.Get(); hasGitHubUsername {
		return string(githubUsername) + "/" + baseBranchName
	}
	return baseBranchName
}
