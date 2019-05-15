package endpoints

import (
	"go-pinnacle/models/bet"
	"go-pinnacle/models/bets"
	"go-pinnacle/request"
)

type Bets struct {
	request.Request
}

func NewBets(r request.Request) *Bets {
	return &Bets{Request: r}
}

type BettingStatusResponse struct {
	Status string
}

func (l *Bets) GetBettingStatus() (*BettingStatusResponse, error) {
	var res = new(BettingStatusResponse)
	err := l.Send(res, "GET", "v1/bets/betting-status", nil, nil)
	return res, err
}

func (l *Bets) PlaceStraightBet(r *bets.PlaceStraightBetRequest) (*bets.PlaceStraightBetResponse, error) {
	var res = new(bets.PlaceStraightBetResponse)
	err := l.Send(res, "POST", "v2/bets/straight", nil, r)
	return res, err
}

func (l *Bets) PlaceParlayBet(r *bets.ParlayBetRequest) (*bets.PlaceParlayBetResponse, error) {
	var res = new(bets.PlaceParlayBetResponse)
	err := l.Send(res, "POST", "v1/bets/parlay", nil, r)
	return res, err
}

func (l *Bets) PlaceTeaserBet(r *bets.TeaserBetRequest) (*bets.TeaserBetResponse, error) {
	var res = new(bets.TeaserBetResponse)
	err := l.Send(res, "POST", "v1/bets/teaser", nil, r)
	return res, err
}

// todo: implement special bet.

func (l *Bets) GetBets(r *bet.BetsRequest) (*bet.BetsResponse, error) {
	var res = new(bet.BetsResponse)
	err := l.Send(res, "POST", "v3/bets", r, nil)
	return res, err
}
