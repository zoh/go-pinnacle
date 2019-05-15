package lines

import "go-pinnacle/models"

type LineResponse struct {
	// NOT_EXISTS or SUCCESS
	Status        string  `json:"status"`
	Price         float64 `json:"price"`
	LineID        int64   `json:"lineId"`
	AltLineID     int64   `json:"altLineId"`
	Team1Score    int32   `json:"team1Score"`
	Team2Score    int32   `json:"team2Score"`
	Team1RedCards int32   `json:"team1RedCards"`
	Team2RedCards int32   `json:"team2RedCards"`
	MaxRiskStake  float64 `json:"maxRiskStake"`
	MinRiskStake  float64 `json:"minRiskStake"`
	MaxWinStake   float64 `json:"maxWinStake"`
	MinWinStake   float64 `json:"minWinStake"`
	EffectiveAsOf string  `json:"effectiveAsOf"`
}

type LineRequest struct {
	// Team Chosen Team type. This is needed only for SPREAD, MONEYLINE and TEAM_TOTAL_POINTS bet types
	Team models.TeamType `url:"team,omitempty"`
	// Chosen Side. This is needed only for TOTAL_POINTS and TEAM_TOTAL_POINTS
	Side models.Side `url:"side,omitempty"`
}

// ** odds parlay
type LineParlayRequest struct {
	OddsFormat models.OddsFormat `url:"oddsFormat,omitempty"`
	Legs       []LineParlayLeg   `url:"legs,omitempty"`
}

type LineParlayLeg struct {
	UniqueLegId  string          `url:"uniqueLegId,omitempty"`
	EventId      int64           `url:"eventId,omitempty"`
	PeriodNumber int32           `url:"periodNumber,omitempty"` // 0 (Game), 1 (1st Half), 2 (2nd Half)
	LegBetType   models.BetType  `url:"legBetType,omitempty"`
	Team         models.TeamType `url:"team,omitempty"`
	Side         models.Side     `url:"side,omitempty"`
	Handicap     float32         `url:"handicap,omitempty"` //This is needed for SPREAD and TOTAL_POINTS bet type
}

// resp
type LineParlayResponse struct {
	Status                   string                    `json:"status"`
	Error                    string                    `json:"error"`
	MaxRiskStake             int64                     `json:"maxRiskStake"`
	MinRiskStake             int64                     `json:"minRiskStake"`
	RoundRobinOptionWithOdds []RoundRobinOptionWithOdd `json:"roundRobinOptionWithOdds"`
	Legs                     []Leg                     `json:"legs"`
}

type Leg struct {
	Status         string   `json:"status"`
	ErrorCode      string   `json:"errorCode"`
	LegID          string   `json:"legId"`
	LineID         int64    `json:"lineId"`
	AltLineID      int64    `json:"altLineId"`
	Price          int64    `json:"price"`
	CorrelatedLegs []string `json:"correlatedLegs"`
}

type RoundRobinOptionWithOdd struct {
	RoundRobinOption     string `json:"roundRobinOption"`
	Odds                 int64  `json:"odds"`
	UnroundedDecimalOdds int64  `json:"unroundedDecimalOdds"`
}

type LineTeaserRequest struct {
	TeaserId   int64             `url:"teaserId,omitempty"`
	OddsFormat models.OddsFormat `url:"oddsFormat,omitempty"`

	Legs []LineParlayLeg
}

type LineTeaserResponse struct {
	Status       string      `json:"status"`
	ErrorCode    string      `json:"errorCode"`
	Price        int64       `json:"price"`
	MinRiskStake int64       `json:"minRiskStake"`
	MaxRiskStake int64       `json:"maxRiskStake"`
	MinWinStake  int64       `json:"minWinStake"`
	MaxWinStake  int64       `json:"maxWinStake"`
	Legs         []LegTeaser `json:"legs"`
}

type LegTeaser struct {
	Status    string `json:"status"`
	ErrorCode string `json:"errorCode"`
	LegID     string `json:"legId"`
	LineID    int64  `json:"lineId"`
}

type LineSpecialResponse struct {
	Status       string  `json:"status"`
	SpecialID    int64   `json:"specialId"`
	ContestantID int64   `json:"contestantId"`
	MinRiskStake float64 `json:"minRiskStake"`
	MaxRiskStake float64 `json:"maxRiskStake"`
	MinWinStake  float64 `json:"minWinStake"`
	MaxWinStake  float64 `json:"maxWinStake"`
	LineID       int64   `json:"lineId"`
	Price        float64 `json:"price"`
	Handicap     float64 `json:"handicap"`
}
