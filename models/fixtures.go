package models

import "time"

type FixturesResponse struct {
	SportID int64    `json:"sportId"`
	Last    int64    `json:"last"`
	League  []League `json:"league"`
}

type League struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Events []Event `json:"events"`
}

type Event struct {
	ID                int64     `json:"id"`
	ParentID          int64     `json:"parentId"`
	Starts            time.Time `json:"starts"`
	Home              string    `json:"home"`
	Away              string    `json:"away"`
	RotNum            string    `json:"rotNum"`
	LiveStatus        int64     `json:"liveStatus"`
	HomePitcher       string    `json:"homePitcher"`
	AwayPitcher       string    `json:"awayPitcher"`
	Status            string    `json:"status"`
	ParlayRestriction int64     `json:"parlayRestriction"`
	AltTeaser         bool      `json:"altTeaser"`
	ResultingUnit     string    `json:"resultingUnit"`
}

// request
type FixturesRequest struct {
	SportID   int32 `json:"sportId"`
	LeagueIDs []int `json:"leagueIds"`
	IsLive    bool  `json:"isLive"`
	Since     int64 `json:"since"`
	EventIDs  []int `json:"eventIds"`
}

// ************************
// *** special fixtures ******
type FixturesSpecialResponse struct {
	SportID int64           `json:"sportId"`
	Last    int64           `json:"last"`
	Leagues []LeagueSpecial `json:"leagues"`
}

type LeagueSpecial struct {
	ID       int64     `json:"id"`
	Specials []Special `json:"specials"`
}

type Special struct {
	ID          int64        `json:"id"`
	BetType     string       `json:"betType"`
	Name        string       `json:"name"`
	Date        string       `json:"date"`
	Cutoff      string       `json:"cutoff"`
	Category    string       `json:"category"`
	Units       string       `json:"units"`
	Status      string       `json:"status"`
	Contestants []Contestant `json:"contestants"`
	LiveStatus  int64        `json:"liveStatus"`
}

type Contestant struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	RotNum int64  `json:"rotNum"`
}

type FixturesSpecialRequest struct {
	SportID   int32  `json:"sportId"`
	LeagueIDs []int  `json:"leagueIds"`
	Since     int64  `json:"since"`
	EventIDs  []int  `json:"eventIds"`
	Category  string `json:"category"`
	EventId   int64  `json:"eventId"`
	SpecialId int64  `json:"specialId"`
}

// ****************************
// *** fixtures settled ****
type FixturesSettledRequest struct {
	SportID   int32 `json:"sportId"`
	LeagueIDs []int `json:"leagueIds"`
	Since     int64 `json:"since"`
}

type FixturesSettledResponse struct {
	SportID int64    `json:"sportId"`
	Last    int64    `json:"last"`
	Leagues []League `json:"leagues"`
}

type LeagueSettled struct {
	ID     int64   `json:"id"`
	Events []Event `json:"events"`
}

type EventSettled struct {
	ID      int64    `json:"id"`
	Periods []Period `json:"periods"`
}

type Period struct {
	Number             int64               `json:"number"`
	Status             int64               `json:"status"`
	SettlementID       int64               `json:"settlementId"`
	SettledAt          string              `json:"settledAt"`
	Team1Score         int64               `json:"team1Score"`
	Team2Score         int64               `json:"team2Score"`
	Team1ScoreSets     int64               `json:"team1ScoreSets"`
	Team2ScoreSets     int64               `json:"team2ScoreSets"`
	CancellationReason *CancellationReason `json:"cancellationReason,omitempty"`
}

type CancellationReason struct {
	Code    string  `json:"code"`
	Details Details `json:"details"`
}

type Details struct {
	CorrectTeam1Score *string `json:"correctTeam1Score,omitempty"`
	CorrectTeam2Score *string `json:"correctTeam2Score,omitempty"`
}

// **********************
// ***** Settled Special Fixtures
type FixturesSpecialSettledRequest struct {
	SportID   int32 `json:"sportId"`
	LeagueIDs []int `json:"leagueIds"`
	Since     int64 `json:"since"`
}

type FixturesSpecialSettledResponse struct {
	SportID int64                  `json:"sportId"`
	Last    int64                  `json:"last"`
	Leagues []LeagueSpecialSettled `json:"leagues"`
}

type LeagueSpecialSettled struct {
	ID       int64            `json:"id"`
	Specials []SpecialSettled `json:"specials"`
}

type SpecialSettled struct {
	ID           int64  `json:"id"`
	Status       int64  `json:"status"`
	SettlementID int64  `json:"settlementId"`
	SettledAt    string `json:"settledAt"`
}
