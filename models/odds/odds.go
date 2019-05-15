package odds

import (
	"go-pinnacle/models"
	"time"
)

type OddsResponse struct {
	SportID int64    `json:"sportId"`
	Last    int64    `json:"last"`
	Leagues []League `json:"leagues"`
}

type League struct {
	ID     int64   `json:"id"`
	Events []Event `json:"events"`
}

type Event struct {
	ID      int64    `json:"id"`
	Periods []Period `json:"periods"`
}

type Period struct {
	LineID       int64      `json:"lineId"`
	Number       int64      `json:"number"`
	Cutoff       time.Time  `json:"cutoff"`
	MaxMoneyline float64    `json:"maxMoneyline"`
	Status       int64      `json:"status"`
	Moneyline    Moneyline  `json:"moneyline"`
	MaxSpread    *float64   `json:"maxSpread,omitempty"`
	MaxTotal     *float64   `json:"maxTotal,omitempty"`
	MaxTeamTotal *float64   `json:"maxTeamTotal,omitempty"`
	Spreads      []Spread   `json:"spreads"`
	Totals       []Away     `json:"totals"`
	TeamTotal    *TeamTotal `json:"teamTotal,omitempty"`
}

type Moneyline struct {
	Home float64 `json:"home"`
	Away float64 `json:"away"`
}

type Spread struct {
	Hdp       float64 `json:"hdp"`
	Home      float64 `json:"home"`
	Away      float64 `json:"away"`
	AltLineID *int64  `json:"altLineId,omitempty"`
}

type TeamTotal struct {
	Home Away `json:"home"`
	Away Away `json:"away"`
}

type Away struct {
	Points    float64 `json:"points"`
	Over      float64 `json:"over"`
	Under     float64 `json:"under"`
	AltLineID *int64  `json:"altLineId,omitempty"`
}

type OddsRequest struct {
	SportID    int32             `url:"sportId"`
	LeagueIds  []int             `url:"leagueIds"`
	OddsFormat models.OddsFormat `url:"oddsFormat,omitempty"`
	Since      int64             `url:"since,omitempty"`
	IsLive     bool              `url:"isLive,omitempty"`
	EventIds   []int             `url:"eventIds,omitempty"`

	// Default USD.
	ToCurrencyCode string `url:"toCurrencyCode,omitempty"`
}
