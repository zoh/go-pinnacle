package odds

import "go-pinnacle/models"

type OddsSpecialRequest struct {
	OddsFormat models.OddsFormat `url:"oddsFormat,omitempty"`
	SportID    int               `url:"sportId,omitempty"`
	LeagueIDs  []int             `url:"leagueIds,omitempty"`
	Since      int               `url:"since,omitempty"`
	SpecialID  []int             `url:"specialId"`
}

type OddsSpecialResponse struct {
	SportID int64           `json:"sportId"`
	Last    int64           `json:"last"`
	Leagues []LeagueSpecial `json:"leagues"`
}

type LeagueSpecial struct {
	ID       int64     `json:"id"`
	Specials []Special `json:"specials"`
}

type Special struct {
	ID              int64            `json:"id"`
	MaxBet          int64            `json:"maxBet"`
	ContestantLines []ContestantLine `json:"contestantLines"`
}

type ContestantLine struct {
	ID       int64       `json:"id"`
	LineID   int64       `json:"lineId"`
	Price    int64       `json:"price"`
	Handicap interface{} `json:"handicap"`
}
