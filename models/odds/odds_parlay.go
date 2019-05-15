package odds

type OddsParlayResponse struct {
	SportID int64          `json:"sportId"`
	Last    int64          `json:"last"`
	Leagues []LeagueParlay `json:"leagues"`
}

type LeagueParlay struct {
	ID     int64         `json:"id"`
	Events []EventParlay `json:"events"`
}

type EventParlay struct {
	ID           int64          `json:"id"`
	AwayScore    float64        `json:"awayScore"`
	HomeScore    float64        `json:"homeScore"`
	AwayRedCards int64          `json:"awayRedCards"`
	HomeRedCards int64          `json:"homeRedCards"`
	Periods      []PeriodParlay `json:"periods"`
}

type PeriodParlay struct {
	LineID       int64           `json:"lineId"`
	Number       int64           `json:"number"`
	Cutoff       string          `json:"cutoff"`
	Status       int64           `json:"status"`
	MaxSpread    float64         `json:"maxSpread"`
	MaxMoneyline float64         `json:"maxMoneyline"`
	MaxTotal     float64         `json:"maxTotal"`
	MaxTeamTotal float64         `json:"maxTeamTotal"`
	Spreads      []SpreadParlay  `json:"spreads"`
	Moneyline    MoneylineParlay `json:"moneyline"`
	Totals       []Away          `json:"totals"`
	TeamTotal    TeamTotalParlay `json:"teamTotal"`
}

type MoneylineParlay struct {
	Home float64 `json:"home"`
	Away float64 `json:"away"`
	Draw float64 `json:"draw"`
}

type SpreadParlay struct {
	AltLineID int64   `json:"altLineId"`
	Hdp       float64 `json:"hdp"`
	Home      float64 `json:"home"`
	Away      float64 `json:"away"`
}

type TeamTotalParlay struct {
	Away AwayParlay `json:"away"`
	Home AwayParlay `json:"home"`
}

type AwayParlay struct {
	AltLineID int64   `json:"altLineId"`
	Points    int64   `json:"points"`
	Over      float64 `json:"over"`
	Under     float64 `json:"under"`
}


//**** Odds teaser
type OddsTeaserResponse struct {
	TeaserID int64          `json:"teaserId"`
	SportID  int64          `json:"sportId"`
	Leagues  []LeagueTeaser `json:"leagues"`
}

type LeagueTeaser struct {
	ID     int64         `json:"id"`
	Events []EventTeaser `json:"events"`
}

type EventTeaser struct {
	ID      int64          `json:"id"`
	Periods []PeriodTeaser `json:"periods"`
}

type PeriodTeaser struct {
	Number int64        `json:"number"`
	LineID int64        `json:"lineId"`
	Spread SpreadTeaser `json:"spread"`
	Total  TotalTeaser  `json:"total"`
}

type SpreadTeaser struct {
	MaxBet  int64   `json:"maxBet"`
	HomeHdp float64 `json:"homeHdp"`
	AwayHdp float64 `json:"awayHdp"`
	AltHdp  bool    `json:"altHdp"`
}

type TotalTeaser struct {
	MaxBet      int64   `json:"maxBet"`
	OverPoints  float64 `json:"overPoints"`
	UnderPoints float64 `json:"underPoints"`
}
