package models

type SportsResponse struct {
	Sports []Sport `json:"sports"`
}

type Sport struct {
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	HasOfferings        bool   `json:"hasOfferings"`
	LeagueSpecialsCount int64  `json:"leagueSpecialsCount"`
	EventSpecialsCount  int64  `json:"eventSpecialsCount"`
	EventCount          int64  `json:"eventCount"`
}
