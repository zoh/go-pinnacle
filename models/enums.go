package models

type BetType string

const (
	MoneyLine BetType = "MONEYLINE"
	TeamTotal BetType = "TEAM_TOTAL_POINTS"
	Spread    BetType = "SPREAD"
	Total     BetType = "TOTAL_POINTS"
)

type Side string

const (
	Over  Side = "OVER"
	Under Side = "UNDER"
)

type TeamType string

const (
	Team1 TeamType = "Team1"
	Team2 TeamType = "Team2"
	Draw  TeamType = "Draw"
)

type BetListType string

const (
	Settled        BetListType = "SETTLED"
	Running        BetListType = "RUNNING"
	Cancelled      BetListType = "CANCELLED"
	BetListTypeAll BetListType = "ALL"
)

type EventStatus string

const (
	Unavailable EventStatus = "H"
	LowerMax    EventStatus = "I"
	Open        EventStatus = "O"
)

type OddsFormat string

const (
	Decimal    OddsFormat = "DECIMAL"
	American   OddsFormat = "AMERICAN"
	HongKong   OddsFormat = "HONGKONG"
	Indonesian OddsFormat = "INDONESIAN"
	Malay      OddsFormat = "MALAY"
)

type WinRiskType string

const (
	Win  WinRiskType = "WIN"
	Risk WinRiskType = "RISK"
)

type FillType string

const (
	Normal       FillType = "NORMAL"
	FillAndKill  FillType = "FILLANDKILL"
	FillMaxLimit FillType = "FILLMAXLIMIT"
)

type BetStatusesType string

const (
	Won                      BetStatusesType = "WON"
	Lose                     BetStatusesType = "LOSE"
	BetStatusesTypeCancelled BetStatusesType = "CANCELLED"
	Refunded                 BetStatusesType = "REFUNDED"
	NotAccepted              BetStatusesType = "NOT_ACCEPTED"
	Accepted                 BetStatusesType = "ACCEPTED"
	PendingAcceptance        BetStatusesType = "PENDING_ACCEPTANCE"
)

type SortDirType string

const (
	Asc  SortDirType = "ASC"
	Desc SortDirType = "DESC"
)
