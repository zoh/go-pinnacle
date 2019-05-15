package endpoints

import (
	"go-pinnacle/models"
	"go-pinnacle/models/lines"
	"net/url"
	"strconv"
)

func (l *Lines) GetSports() (*models.SportsResponse, error) {
	var res = new(models.SportsResponse)
	err := l.Send(res, "GET", "v2/sports", nil, nil)
	return res, err
}

func (l *Lines) GetLeagues(sportId int32) (*lines.LeaguesResponse, error) {
	var res = new(lines.LeaguesResponse)
	err := l.Send(
		res,
		"GET",
		"v2/leagues",
		url.Values{"sportId": []string{strconv.Itoa(int(sportId))}},
		nil,
	)
	return res, err
}

type PeriodsResponse struct {
	Periods []Period `json:"periods"`
}

type Period struct {
	Number                     int64  `json:"number"`
	Description                string `json:"description"`
	ShortDescription           string `json:"shortDescription"`
	SpreadDescription          string `json:"spreadDescription"`
	MoneylineDescription       string `json:"moneylineDescription"`
	TotalDescription           string `json:"totalDescription"`
	Team1TotalDescription      string `json:"team1TotalDescription"`
	Team2TotalDescription      string `json:"team2TotalDescription"`
	SpreadShortDescription     string `json:"spreadShortDescription"`
	MoneylineShortDescription  string `json:"moneylineShortDescription"`
	TotalShortDescription      string `json:"totalShortDescription"`
	Team1TotalShortDescription string `json:"team1TotalShortDescription"`
	Team2TotalShortDescription string `json:"team2TotalShortDescription"`
}

func (l *Lines) GetPeriods(sportId int) (*PeriodsResponse, error) {
	var res = new(PeriodsResponse)
	err := l.Send(
		res,
		"GET",
		"v2/periods",
		url.Values{"sportId": []string{strconv.Itoa(int(sportId))}},
		nil,
	)
	return res, err
}

type InrunningResponse struct {
	Sports []Sport `json:"sports"`
}

type Sport struct {
	ID      int64    `json:"id"`
	Leagues []League `json:"leagues"`
}

type League struct {
	ID     int64   `json:"id"`
	Events []Event `json:"events"`
}

type Event struct {
	ID      int64 `json:"id"`
	State   int64 `json:"state"`
	Elapsed int64 `json:"elapsed"`
}

// Get In-Running
func (l *Lines) GetInRunning() (*InrunningResponse, error) {
	var res = new(InrunningResponse)
	err := l.Send(
		res,
		"GET",
		"v1/inrunning",
		nil,
		nil,
	)
	return res, err
}

type TeaserGroupsResponse struct {
	TeaserGroups []TeaserGroup `json:"teaserGroups"`
}

type TeaserGroup struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Teasers []Teaser `json:"teasers"`
}

type Teaser struct {
	ID            int64          `json:"id"`
	Description   string         `json:"description"`
	SportID       int64          `json:"sportId"`
	MinLegs       int64          `json:"minLegs"`
	MaxLegs       int64          `json:"maxLegs"`
	SameEventOnly bool           `json:"sameEventOnly"`
	Payouts       []Payout       `json:"payouts"`
	Leagues       []LeagueTeaser `json:"leagues"`
}

type LeagueTeaser struct {
	ID     int64  `json:"id"`
	Spread Spread `json:"spread"`
	Total  Spread `json:"total"`
}

type Spread struct {
	Points float64 `json:"points"`
}

type Payout struct {
	NumberOfLegs int64   `json:"numberOfLegs"`
	Price        float64 `json:"price"`
}

func (l *Lines) GetTeaserGroups(oddsFormat models.OddsFormat) (*TeaserGroupsResponse, error) {
	var res = new(TeaserGroupsResponse)
	err := l.Send(res, "GET", "v1/teaser/groups",
		url.Values{"oddsFormat": []string{string(oddsFormat)}}, nil, )
	return res, err
}

type CancellationReasonsResponse struct {
	CancellationReasons []CancellationReason `json:"cancellationReasons"`
}

type CancellationReason struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Get Cancellation Reasons
func (l *Lines) GetCancellationReasons() (*CancellationReasonsResponse, error) {
	var res = new(CancellationReasonsResponse)
	err := l.Send(res, "GET", "v1/cancellationreasons", nil, nil, )
	return res, err
}

type CurrenciesResponse struct {
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

func (l *Lines) GetCurrencies() (*CurrenciesResponse, error) {
	var res = new(CurrenciesResponse)
	err := l.Send(res, "GET", "v2/currencies", nil, nil, )
	return res, err
}
