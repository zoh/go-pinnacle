package bets

import (
	"go-pinnacle/models"
	"time"
)

type PlaceStraightBetRequest struct {
	OddsFormat        models.OddsFormat  `url:"oddsFormat,omitempty"`
	UniqueRequestID   string             `url:"uniqueRequestId,omitempty"`
	AcceptBetterLine  bool               `url:"acceptBetterLine,omitempty"`
	Stake             float64            `url:"stake,omitempty"`
	WinRiskStake      models.WinRiskType `url:"winRiskStake,omitempty"`
	LineID            int64              `url:"lineId,omitempty"`
	AltLineID         *int64             `url:"altLineId,omitempty"`
	Pitcher1MustStart bool               `url:"pitcher1MustStart,omitempty"`
	Pitcher2MustStart bool               `url:"pitcher2MustStart,omitempty"`
	FillType          models.FillType    `url:"fillType,omitempty"`
	SportID           int64              `url:"sportId,omitempty"`
	EventID           int64              `url:"eventId,omitempty"`
	PeriodNumber      int64              `url:"periodNumber,omitempty"`
	BetType           models.BetType     `url:"betType,omitempty"`
	Team              models.TeamType    `url:"team,omitempty"`
	Side              *models.Side       `url:"side,omitempty"`
}

type PSBStatus string

// "ACCEPTED" "PENDING_ACCEPTANCE" "PROCESSED_WITH_ERROR"
const (
	PSBStatus_ACCEPTED             PSBStatus = "ACCEPTED"
	PSBStatus_PENDING_ACCEPTANCE   PSBStatus = "PENDING_ACCEPTANCE"
	PSBStatus_PROCESSED_WITH_ERROR PSBStatus = "PROCESSED_WITH_ERROR"
)

type PlaceStraightBetResponse struct {
	Status string `json:"status"`

	//
	//string Nullable
	//Enum:"ALL_BETTING_CLOSED" "ALL_LIVE_BETTING_CLOSED" "ABOVE_EVENT_MAX" "ABOVE_MAX_BET_AMOUNT" "BELOW_MIN_BET_AMOUNT" "BLOCKED_BETTING" "BLOCKED_CLIENT" "INSUFFICIENT_FUNDS" "INVALID_COUNTRY" "INVALID_EVENT" "INVALID_ODDS_FORMAT" "LINE_CHANGED" "LISTED_PITCHERS_SELECTION_ERROR" "OFFLINE_EVENT" "PAST_CUTOFFTIME" "RED_CARDS_CHANGED" "SCORE_CHANGED" "TIME_RESTRICTION" "DUPLICATE_UNIQUE_REQUEST_ID" "INCOMPLETE_CUSTOMER_BETTING_PROFILE" "INVALID_CUSTOMER_PROFILE" "LIMITS_CONFIGURATION_ISSUE" "RESPONSIBLE_BETTING_LOSS_LIMIT_EXCEEDED" "RESPONSIBLE_BETTING_RISK_LIMIT_EXCEEDED" "SYSTEM_ERROR_3" "LICENCE_RESTRICTION_LIVE_BETTING_BLOCKED"
	ErrorCode       *string     `json:"errorCode"`
	UniqueRequestID string      `json:"uniqueRequestId"`
	StraightBet     StraightBet `json:"straightBet"`
}

type StraightBet struct {
	BetID              int64                  `json:"betId"`
	WagerNumber        int64                  `json:"wagerNumber"`
	PlacedAt           time.Time              `json:"placedAt"`
	BetStatus          models.BetStatusesType `json:"betStatus"`
	BetType            models.BetType         `json:"betType"`
	Win                float64                `json:"win"`
	Risk               float64                `json:"risk"`
	WinLoss            *float64               `json:"winLoss"`
	OddsFormat         models.OddsFormat      `json:"oddsFormat"`
	CustomerCommission *float64               `json:"customerCommission"`
	CancellationReason CancellationReason     `json:"cancellationReason"`
	UpdateSequence     int64                  `json:"updateSequence"`
	SportID            int64                  `json:"sportId"`
	LeagueID           int64                  `json:"leagueId"`
	EventID            int64                  `json:"eventId"`
	Handicap           *float64               `json:"handicap"`
	Price              int64                  `json:"price"`
	TeamName           string                 `json:"teamName"`
	Side               *models.Side           `json:"side"`
	Pitcher1           *string                `json:"pitcher1"`
	Pitcher2           *string                `json:"pitcher2"`
	Pitcher1MustStart  *string                `json:"pitcher1MustStart"`
	Pitcher2MustStart  *string                `json:"pitcher2MustStart"`
	Team1              string                 `json:"team1"`
	Team2              string                 `json:"team2"`
	PeriodNumber       int64                  `json:"periodNumber"`
	Team1Score         *float64               `json:"team1Score"`
	Team2Score         *float64               `json:"team2Score"`
	FtTeam1Score       *float64               `json:"ftTeam1Score"`
	FtTeam2Score       *float64               `json:"ftTeam2Score"`
	PTeam1Score        *float64               `json:"pTeam1Score"`
	PTeam2Score        *float64               `json:"pTeam2Score"`
	IsLive             bool                   `json:"isLive"`
}

type CancellationReason struct {
	Code    string   `json:"code"`
	Details []Detail `json:"details"`
}

type Detail struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// req & res for bet parlay
type ParlayBetRequest struct {
	UniqueRequestID  string            `url:"uniqueRequestId,omitempty"`
	AcceptBetterLine bool              `url:"acceptBetterLine,omitempty"`
	RiskAmount       float64           `url:"riskAmount,omitempty"`
	OddsFormat       models.OddsFormat `url:"oddsFormat,omitempty"`
	Legs             []LegParlay       `url:"legs,omitempty"`

	//Array of string
	//Items Enum:"Parlay" "TwoLegRoundRobin" "ThreeLegRoundRobin" "FourLegRoundRobin" "FiveLegRoundRobin" "SixLegRoundRobin" "SevenLegRoundRobin" "EightLegRoundRobin"
	RoundRobinOptions []string `url:"roundRobinOptions,omitempty"`
}

type LegParlay struct {
	UniqueLegID       string         `url:"uniqueLegId,omitempty"`
	LineID            int64          `url:"lineId,omitempty"`
	AltLineID         *int64         `url:"altLineId,omitempty"`
	Pitcher1MustStart bool           `url:"pitcher1MustStart,omitempty"`
	Pitcher2MustStart bool           `url:"pitcher2MustStart,omitempty"`
	SportID           int64          `url:"sportId,omitempty"`
	EventID           int64          `url:"eventId,omitempty"`
	PeriodNumber      int32          `url:"periodNumber,omitempty"` // 0 - Game, 1 - 1st Half and 2 - 2nd Half
	LegBetType        models.BetType `url:"legBetType,omitempty"`
	Team              string         `url:"team,omitempty"`
	Side              *models.Side   `url:"side,omitempty"`
}

type PlaceParlayBetResponse struct {
	Status                   string                    `json:"status"`
	ErrorCode                *string                   `json:"errorCode"`
	BetID                    int64                     `json:"betId"`
	UniqueRequestID          string                    `json:"uniqueRequestId"`
	RoundRobinOptionWithOdds []RoundRobinOptionWithOdd `json:"roundRobinOptionWithOdds"`
	MaxRiskStake             int64                     `json:"maxRiskStake"`
	MinRiskStake             int64                     `json:"minRiskStake"`
	ValidLegs                []ValidLeg                `json:"validLegs"`
	InvalidLegs              []ValidLeg                `json:"invalidLegs"`
	ParlayBet                ParlayBet                 `json:"parlayBet"`
}

type ValidLeg struct {
	Status         string   `json:"status"`
	ErrorCode      *string  `json:"errorCode"`
	LegID          string   `json:"legId"`
	LineID         int64    `json:"lineId"`
	AltLineID      *int64   `json:"altLineId"`
	Price          int64    `json:"price"`
	CorrelatedLegs []string `json:"correlatedLegs"`
}

type ParlayBet struct {
	BetID              int64               `json:"betId"`
	UniqueRequestID    string              `json:"uniqueRequestId"`
	WagerNumber        int64               `json:"wagerNumber"`
	PlacedAt           string              `json:"placedAt"`
	BetStatus          string              `json:"betStatus"`
	BetType            string              `json:"betType"`
	Win                float64             `json:"win"`
	Risk               int64               `json:"risk"`
	WinLoss            *float64            `json:"winLoss"`
	OddsFormat         string              `json:"oddsFormat"`
	CustomerCommission *float64            `json:"customerCommission"`
	CancellationReason CancellationReason2 `json:"cancellationReason"`
	UpdateSequence     int64               `json:"updateSequence"`
	Legs               []Leg               `json:"legs"`
	Price              int64               `json:"price"`
	FinalPrice         int64               `json:"finalPrice"`
}

type CancellationReason2 struct {
	Code    string    `json:"code"`
	Details []Detail2 `json:"details"`
}

type Detail2 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Leg struct {
	SportID            int64              `json:"sportId"`
	LegBetType         string             `json:"legBetType"`
	LegBetStatus       string             `json:"legBetStatus"`
	LeagueID           int64              `json:"leagueId"`
	EventID            int64              `json:"eventId"`
	EventStartTime     string             `json:"eventStartTime"`
	Handicap           *float64           `json:"handicap"`
	Price              int64              `json:"price"`
	TeamName           string             `json:"teamName"`
	Side               *models.Side       `json:"side"`
	Pitcher1           *string            `json:"pitcher1"`
	Pitcher2           *string            `json:"pitcher2"`
	Pitcher1MustStart  bool               `json:"pitcher1MustStart"`
	Pitcher2MustStart  bool               `json:"pitcher2MustStart"`
	Team1              string             `json:"team1"`
	Team2              string             `json:"team2"`
	PeriodNumber       int64              `json:"periodNumber"`
	FtTeam1Score       *float64           `json:"ftTeam1Score"`
	FtTeam2Score       *float64           `json:"ftTeam2Score"`
	PTeam1Score        *float64           `json:"pTeam1Score"`
	PTeam2Score        *float64           `json:"pTeam2Score"`
	CancellationReason CancellationReason `json:"cancellationReason"`
}

type RoundRobinOptionWithOdd struct {
	RoundRobinOption     string  `json:"roundRobinOption"`
	Odds                 int64   `json:"odds"`
	UnroundedDecimalOdds float64 `json:"unroundedDecimalOdds"`
}

////////////////
type TeaserBetRequest struct {
	UniqueRequestID string      `json:"uniqueRequestId"`
	TeaserID        int64       `json:"teaserId"`
	OddsFormat      string      `json:"oddsFormat"`
	WinRiskStake    string      `json:"winRiskStake"`
	Stake           int64       `json:"stake"`
	Legs            []LegTeaser `json:"legs"`
}

type LegTeaser struct {
	LegID   string         `json:"legId"`
	BetType models.BetType `json:"betType"`
	LineID  int64          `json:"lineId"`
	EventID int64          `json:"eventId"`
	Team    string         `json:"team"`
	Side    models.Side    `json:"side"`
}

type TeaserBetResponse struct {
	Status          string      `json:"status"`
	ErrorCode       string      `json:"errorCode"`
	BetID           int64       `json:"betId"`
	UniqueRequestID string      `json:"uniqueRequestId"`
	Price           int64       `json:"price"`
	RiskAmount      int64       `json:"riskAmount"`
	WinAmount       int64       `json:"winAmount"`
	WinRiskStake    string      `json:"winRiskStake"`
	InvalidLegs     []ValidLeg2 `json:"invalidLegs"`
	ValidLegs       []ValidLeg2 `json:"validLegs"`
	TeaserBet       TeaserBet   `json:"teaserBet"`
}

type ValidLeg2 struct {
	Status    string `json:"status"`
	ErrorCode string `json:"errorCode"`
	LegID     string `json:"legId"`
	LineID    int64  `json:"lineId"`
	Points    int64  `json:"points"`
}

type TeaserBet struct {
	BetID              int64                  `json:"betId"`
	UniqueRequestID    string                 `json:"uniqueRequestId"`
	WagerNumber        int64                  `json:"wagerNumber"`
	PlacedAt           string                 `json:"placedAt"`
	BetStatus          models.BetStatusesType `json:"betStatus"`
	BetType            models.BetType         `json:"betType"`
	Win                int64                  `json:"win"`
	Risk               int64                  `json:"risk"`
	WinLoss            int64                  `json:"winLoss"`
	OddsFormat         models.OddsFormat      `json:"oddsFormat"`
	CustomerCommission int64                  `json:"customerCommission"`
	CancellationReason CancellationReason3    `json:"cancellationReason"`
	UpdateSequence     int64                  `json:"updateSequence"`
	TeaserName         string                 `json:"teaserName"`
	IsSameEventOnly    bool                   `json:"isSameEventOnly"`
	MinPicks           int64                  `json:"minPicks"`
	MaxPicks           int64                  `json:"maxPicks"`
	Price              int64                  `json:"price"`
	FinalPrice         int64                  `json:"finalPrice"`
	TeaserID           int64                  `json:"teaserId"`
	TeaserGroupID      int64                  `json:"teaserGroupId"`
	Legs               []Leg3                 `json:"legs"`
}

type CancellationReason3 struct {
	Code    string    `json:"code"`
	Details []Detail3 `json:"details"`
}

type Detail3 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Leg3 struct {
	SportID        int64  `json:"sportId"`
	LegBetType     string `json:"legBetType"`
	LegBetStatus   string `json:"legBetStatus"`
	LeagueID       int64  `json:"leagueId"`
	EventID        int64  `json:"eventId"`
	EventStartTime string `json:"eventStartTime"`
	Handicap       int64  `json:"handicap"`
	TeamName       string `json:"teamName"`
	Side           string `json:"side"`
	Team1          string `json:"team1"`
	Team2          string `json:"team2"`
	PeriodNumber   int64  `json:"periodNumber"`
}
