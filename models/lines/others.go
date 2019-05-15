package lines

import "go-pinnacle/models"

type LeaguesResponse struct {
	Leagues []League `json:"leagues"`
}

type League struct {
	ID                  int64           `json:"id"`
	Name                string          `json:"name"`
	HomeTeamType        models.TeamType `json:"homeTeamType"`
	HasOfferings        bool            `json:"hasOfferings"`
	Container           string          `json:"container"`
	AllowRoundRobins    bool            `json:"allowRoundRobins"`
	LeagueSpecialsCount int64           `json:"leagueSpecialsCount"`
	EventSpecialsCount  int64           `json:"eventSpecialsCount"`
	EventCount          int64           `json:"eventCount"`
}
