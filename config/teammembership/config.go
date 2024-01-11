package teammembership

import "github.com/upbound/upjet/pkg/config"

// Configure github_team_membership resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_team_membership", func(r *config.Resource) {
		r.Kind = "TeamMembership"
		r.ShortGroup = "team"

		r.References["team_id"] = config.Reference{
			Type: "github.com/coopnorge/provider-github/apis/team/v1alpha1.Team",
		}
	})
}
