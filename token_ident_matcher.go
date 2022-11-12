package gh2tf

import "github.com/yoanm/go-tfsig"

var defaultUnescapedStringPrefixList []string

func init() {
	defaultUnescapedStringPrefixList = []string{
		// gitHub provider resources name - START
		"github_actions_environment_secret.",
		"github_actions_organization_permissions.",
		"github_actions_organization_secret.",
		"github_actions_organization_secret_repositories.",
		"github_actions_runner_group.",
		"github_actions_secret.",
		"github_app_installation_repository.",
		"github_branch.",
		"github_branch_default.",
		"github_branch_protection.",
		"github_branch_protection_v3.",
		"github_issue_label.",
		"github_membership.",
		"github_organization_block.",
		"github_organization_project.",
		"github_organization_webhook.",
		"github_project_card.",
		"github_project_column.",
		"github_repository.",
		"github_repository_autolink_reference.",
		"github_repository_collaborator.",
		"github_repository_deploy_key.",
		"github_repository_environment.",
		"github_repository_file.",
		"github_repository_milestone.",
		"github_repository_project.",
		"github_repository_pull_request.",
		"github_repository_webhook.",
		"github_team.",
		"github_team_membership.",
		"github_team_repository.",
		"github_team_sync_group_mapping.",
		"github_user_gpg_key.",
		"github_user_invitation_accepter.",
		"github_user_ssh_key.",
		// gitHub provider resources name - END
	}
}

// NewIdentTokenMatcher returns an instance of IdentTokenMatcher with provided list of prefix to consider as 'ident' tokens
//
// Default `tfsig` ident tokens and terraform gitHub resources will be automatically added
func NewIdentTokenMatcher(prefixList ...string) tfsig.IdentTokenMatcher {
	return tfsig.NewIdentTokenMatcher(append(prefixList, defaultUnescapedStringPrefixList...)...)
}
