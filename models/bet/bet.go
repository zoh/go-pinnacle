package bet

import (
	"go-pinnacle/models"
	"time"
)

// req
type BetsRequest struct {
	BetList     models.BetListType     `url:"betlist,omitempty"`
	BetStatuses models.BetStatusesType `url:"betStatuses,omitempty"`

	FromDate         time.Time          `url:"fromDate,omitempty"`
	ToDate           time.Time          `url:"toDate,omitempty"`
	SortDir          models.SortDirType `url:"sortDir,omitempty"`
	PageSize         int                `url:"pageSize,omitempty"`
	FromRecord       int                `url:"fromRecord,omitempty"`
	BetIDs           []int              `url:"betids,omitempty"`
	UniqueRequestIds []string           `url:"uniqueRequestIds,omitempty"`
}

// res
type BetsResponse struct {
	MoreAvailable bool          `json:"moreAvailable"`
	PageSize      int64         `json:"pageSize"`
	FromRecord    int64         `json:"fromRecord"`
	ToRecord      int64         `json:"toRecord"`
	StraightBets  []StraightBet `json:"straightBets"`
	ParlayBets    []ParlayBet   `json:"parlayBets"`
	TeaserBets    []TeaserBet   `json:"teaserBets"`
	SpecialBets   []SpecialBet  `json:"specialBets"`
	ManualBets    []ManualBet   `json:"manualBets"`
}

type ManualBet struct {
	BetID          int64                  `json:"betId"`
	WagerNumber    int64                  `json:"wagerNumber"`
	PlacedAt       time.Time              `json:"placedAt"`
	BetStatus      models.BetStatusesType `json:"betStatus"`
	BetType        models.BetType         `json:"betType"`
	Win            float64                `json:"win"`
	Risk           float64                `json:"risk"`
	WinLoss        float64                `json:"winLoss"`
	UpdateSequence int64                  `json:"updateSequence"`
	Description    string                 `json:"description"`
	ReferenceBetID *int64                 `json:"referenceBetId"`
}

type ParlayBet struct {
	BetID              int64                  `json:"betId"`
	UniqueRequestID    string                 `json:"uniqueRequestId"`
	WagerNumber        int64                  `json:"wagerNumber"`
	PlacedAt           string                 `json:"placedAt"`
	BetStatus          models.BetStatusesType `json:"betStatus"`
	BetType            models.BetType         `json:"betType"`
	Win                float64                `json:"win"`
	Risk               int64                  `json:"risk"`
	WinLoss            *float64               `json:"winLoss"`
	OddsFormat         string                 `json:"oddsFormat"`
	CustomerCommission *float64               `json:"customerCommission"`
	CancellationReason CancellationReason     `json:"cancellationReason"`
	UpdateSequence     int64                  `json:"updateSequence"`
	Legs               []ParlayBetLeg         `json:"legs"`
	Price              float64                `json:"price"`
	FinalPrice         float64                `json:"finalPrice"`
}

type CancellationReason struct {
	Code    string   `json:"code"`
	Details []Detail `json:"details"`
}

type Detail struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ParlayBetLeg struct {
	SportID            int64                  `json:"sportId"`
	LegBetType         string                 `json:"legBetType"`
	LegBetStatus       models.BetStatusesType `json:"legBetStatus"`
	LeagueID           int64                  `json:"leagueId"`
	EventID            int64                  `json:"eventId"`
	EventStartTime     string                 `json:"eventStartTime"`
	Handicap           *float64               `json:"handicap"`
	Price              int64                  `json:"price"`
	TeamName           string                 `json:"teamName"`
	Side               *models.Side           `json:"side"`
	Pitcher1           *string                `json:"pitcher1"`
	Pitcher2           *string                `json:"pitcher2"`
	Pitcher1MustStart  bool                   `json:"pitcher1MustStart"`
	Pitcher2MustStart  bool                   `json:"pitcher2MustStart"`
	Team1              string                 `json:"team1"`
	Team2              string                 `json:"team2"`
	PeriodNumber       int64                  `json:"periodNumber"`
	FtTeam1Score       *float64               `json:"ftTeam1Score"`
	FtTeam2Score       *float64               `json:"ftTeam2Score"`
	PTeam1Score        *float64               `json:"pTeam1Score"`
	PTeam2Score        *float64               `json:"pTeam2Score"`
	CancellationReason CancellationReason     `json:"cancellationReason"`
}

type SpecialBet struct {
	BetID              int64              `json:"betId"`
	UniqueRequestID    string             `json:"uniqueRequestId"`
	WagerNumber        int64              `json:"wagerNumber"`
	PlacedAt           string             `json:"placedAt"`
	BetStatus          string             `json:"betStatus"`
	BetType            string             `json:"betType"`
	Win                int64              `json:"win"`
	Risk               float64            `json:"risk"`
	WinLoss            *float64           `json:"winLoss"`
	OddsFormat         string             `json:"oddsFormat"`
	CustomerCommission *float64           `json:"customerCommission"`
	CancellationReason CancellationReason `json:"cancellationReason"`
	UpdateSequence     int64              `json:"updateSequence"`
	SpecialID          int64              `json:"specialId"`
	SpecialName        string             `json:"specialName"`
	ContestantID       int64              `json:"contestantId"`
	ContestantName     string             `json:"contestantName"`
	Price              int64              `json:"price"`
	Handicap           float64            `json:"handicap"`
	Units              string             `json:"units"`
	SportID            int64              `json:"sportId"`
	LeagueID           int64              `json:"leagueId"`
	EventID            *int64             `json:"eventId"`
	PeriodNumber       *int32             `json:"periodNumber"`
	Team1              *string            `json:"team1"`
	Team2              *string            `json:"team2"`
	EventStartTime     string             `json:"eventStartTime"`
}

type StraightBet struct {
	BetID              int64              `json:"betId"`
	WagerNumber        int64              `json:"wagerNumber"`
	PlacedAt           string             `json:"placedAt"`
	BetStatus          string             `json:"betStatus"`
	BetType            string             `json:"betType"`
	Win                int64              `json:"win"`
	Risk               float64            `json:"risk"`
	WinLoss            *float64           `json:"winLoss"`
	OddsFormat         string             `json:"oddsFormat"`
	CustomerCommission *float64           `json:"customerCommission"`
	CancellationReason CancellationReason `json:"cancellationReason"`
	UpdateSequence     int64              `json:"updateSequence"`
	SportID            int64              `json:"sportId"`
	LeagueID           int64              `json:"leagueId"`
	EventID            int64              `json:"eventId"`
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
	Team1Score         *float64           `json:"team1Score"`
	Team2Score         *float64           `json:"team2Score"`
	FtTeam1Score       *float64           `json:"ftTeam1Score"`
	FtTeam2Score       *float64           `json:"ftTeam2Score"`
	PTeam1Score        *float64           `json:"pTeam1Score"`
	PTeam2Score        *float64           `json:"pTeam2Score"`
	IsLive             bool               `json:"isLive"`
	EventStartTime     string             `json:"eventStartTime"`
}

type TeaserBet struct {
	BetID              int64              `json:"betId"`
	UniqueRequestID    string             `json:"uniqueRequestId"`
	WagerNumber        int64              `json:"wagerNumber"`
	PlacedAt           string             `json:"placedAt"`
	BetStatus          string             `json:"betStatus"`
	BetType            string             `json:"betType"`
	Win                int64              `json:"win"`
	Risk               int64              `json:"risk"`
	WinLoss            int64              `json:"winLoss"`
	OddsFormat         string             `json:"oddsFormat"`
	CustomerCommission int64              `json:"customerCommission"`
	CancellationReason CancellationReason `json:"cancellationReason"`
	UpdateSequence     int64              `json:"updateSequence"`
	TeaserName         string             `json:"teaserName"`
	IsSameEventOnly    bool               `json:"isSameEventOnly"`
	MinPicks           int64              `json:"minPicks"`
	MaxPicks           int64              `json:"maxPicks"`
	Price              int64              `json:"price"`
	FinalPrice         int64              `json:"finalPrice"`
	TeaserID           int64              `json:"teaserId"`
	TeaserGroupID      int64              `json:"teaserGroupId"`
	Legs               []TeaserBetLeg     `json:"legs"`
}

type TeaserBetLeg struct {
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
