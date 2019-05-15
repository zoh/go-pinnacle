package endpoints

import (
	"fmt"
	"go-pinnacle/models"
	"go-pinnacle/models/lines"
	"go-pinnacle/request"
	"net/url"
	"strconv"
)

type Lines struct {
	request.Request
}

func NewLines(r request.Request) *Lines {
	return &Lines{Request: r}
}

func (l *Lines) GetFixtures(r *models.FixturesRequest) (*models.FixturesResponse, error) {
	var res = new(models.FixturesResponse)
	err := l.Send(res, "GET", "v1/fixtures", r, nil)
	return res, err
}

func (l *Lines) GetFixturesSpecial(r *models.FixturesSpecialRequest) (*models.FixturesSpecialResponse, error) {
	var res = new(models.FixturesSpecialResponse)
	err := l.Send(res, "GET", "v1/fixtures/special", r, nil)
	return res, err
}

func (l *Lines) GetFixturesSettled(r *models.FixturesSettledRequest) (*models.FixturesSettledResponse, error) {
	var res = new(models.FixturesSettledResponse)
	err := l.Send(res, "GET", "v1/fixtures/settled", r, nil)
	return res, err
}

func (l *Lines) GetFixturesSpecialSettled(r *models.FixturesSpecialSettledRequest) (*models.FixturesSpecialSettledResponse, error) {
	var res = new(models.FixturesSpecialSettledResponse)
	err := l.Send(res, "GET", "v1/fixtures/special/settled", r, nil)
	return res, err
}

// Get Straight Line
func (l *Lines) GetLine(
	leagueId int32,
	handicap float32, // This is needed for SPREAD, TOTAL_POINTS and TEAM_TOTAL_POINTS bet types
	oddsFormat models.OddsFormat,
	sportId int32,
	eventId int64,
	periodNumber int32, // This represents the period of the match. For example, for soccer we have 0 (Game), 1 (1st Half) & 2 (2nd Half)
	betType models.BetType,

	opt *lines.LineRequest,
) (*lines.LineResponse, error) {
	r := url.Values{
		"leagueId":     []string{strconv.Itoa(int(leagueId))},
		"handicap":     []string{fmt.Sprintf("%f", handicap)},
		"oddsFormat":   []string{string(oddsFormat)},
		"sportId":      []string{strconv.Itoa(int(sportId))},
		"eventId":      []string{strconv.Itoa(int(eventId))},
		"periodNumber": []string{strconv.Itoa(int(periodNumber))},
		"betType":      []string{string(betType)},
	}
	if opt != nil {
		r.Add("Team", string(opt.Team))
		r.Add("Side", string(opt.Side))
	}

	var res = new(lines.LineResponse)
	err := l.Send(res, "GET", "v1/line", r, nil)
	return res, err
}

// Get Parlay line
func (l *Lines) GetLineParlay(r *lines.LineParlayRequest) (*lines.LineParlayResponse, error) {
	var res = new(lines.LineParlayResponse)
	err := l.Send(res, "POST", "v1/line/parlay", nil, r)
	return res, err
}

func (l *Lines) GetLineTeaser(r *lines.LineTeaserRequest) (*lines.LineTeaserResponse, error) {
	var res = new(lines.LineTeaserResponse)
	err := l.Send(res, "POST", "v1/line/teaser", nil, r)
	return res, err
}

// Get Special Line
func (l *Lines) GetSpecialLine(
	oddsFormat models.OddsFormat,
	specialId int64,
	contestantId int64,

/**
  handicap of the contestant. As contestant's handicap is a mutable property,
  it may happened that line/special returns status:SUCCESS,
  but with the different handicap from the one that client had at the moment of calling the line/special.
  One can specify handicap parameter in the request and if the contestant's handicap changed,
  it would return status:NOT_EXISTS. This way line/special is more aligned to how /line works.
*/
	handicap *int32,
) (*lines.LineSpecialResponse, error) {
	r := url.Values{
		"handicap":   []string{fmt.Sprintf("%d", handicap)},
		"oddsFormat": []string{string(oddsFormat)},
		"betType":    []string{strconv.Itoa(int(contestantId))},
	}
	var res = new(lines.LineSpecialResponse)
	err := l.Send(res, "GET", "v1/line", r, nil)
	return res, err
}
