package endpoints

import "go-pinnacle/models/odds"

func (l *Lines) GetOdds(r *odds.OddsRequest) (*odds.OddsResponse, error) {
	var res = new(odds.OddsResponse)
	err := l.Send(res, "GET", "v1/odds", r, nil)
	return res, err
}

func (l *Lines) GetOddsParlay(r *odds.OddsRequest) (*odds.OddsParlayResponse, error) {
	var res = new(odds.OddsParlayResponse)
	err := l.Send(res, "GET", "v1/odds/parlay", r, nil)
	return res, err
}

func (l *Lines) GetOddsTeaser(teaserId int64) (*odds.OddsTeaserResponse, error) {
	var res = new(odds.OddsTeaserResponse)
	err := l.Send(res, "GET", "v1/odds/parlay", map[string]int64{"teaserId": teaserId}, nil)
	return res, err
}

func (l *Lines) GetOddsSpecial(r *odds.OddsSpecialRequest) (*odds.OddsSpecialResponse, error) {
	var res = new(odds.OddsSpecialResponse)
	err := l.Send(res, "GET", "v1/odds/special", r, nil)
	return res, err
}
