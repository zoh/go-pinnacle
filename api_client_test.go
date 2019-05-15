package go_pinnacle_test

import (
	"fmt"
	"go-pinnacle"
	"go-pinnacle/models"
	"go-pinnacle/models/lines"
	"go-pinnacle/models/odds"
	"go-pinnacle/request"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "go-pinnacle/testing_utils"
)

var (
	username string
	pass     string
	r        *go_pinnacle.ApiClient
)

func TestMain(t *testing.M) {
	username = os.Getenv("PINNACLE_USERNAME")
	pass = os.Getenv("PINNACLE_PASS")

	r = go_pinnacle.NewApiClient(username, pass)

	t.Run()
}

func TestGetClientBalance(t *testing.T) {
	res, _ := r.Account.GetClientBalance()
	log.Printf("%+v\n", res)
}

func iTestGetFixtures(t *testing.T) {
	res, err := r.Lines.GetFixtures(&models.FixturesRequest{
		SportID: 33,
	})

	fmt.Println(res, err)
}

func TestGetSports(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		Equals(t, req.URL.String(), "/v2/sports")
		// Send response to be tested

		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		rw.Write([]byte(`{
"sports": [
{
"id": 1,
"name": "Badminton",
"hasOfferings": false,
"leagueSpecialsCount": 0,
"eventSpecialsCount": 0,
"eventCount": 0
},
{
"id": 2,
"name": "Bandy",
"hasOfferings": false,
"leagueSpecialsCount": 0,
"eventSpecialsCount": 0,
"eventCount": 0
}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	r.Lines.Request = request.NewRequest("", "", server.URL+"/", server.Client())

	res, err := r.Lines.GetSports()
	Equals(t, err, nil)

	Equals(t, len(res.Sports), 2)
	Equals(t, res.Sports[1].ID, int64(2))
	Equals(t, res.Sports[1].Name, "Bandy")
}

func TestOdds(t *testing.T) {
	username = os.Getenv("PINNACLE_USERNAME")
	pass = os.Getenv("PINNACLE_PASS")

	r = go_pinnacle.NewApiClient(username, pass)

	res, err := r.Lines.GetOdds(&odds.OddsRequest{IsLive: true, SportID: 33})
	fmt.Println(res, err)
}

const oddsSpecialResp = `{
    "sportId": 33,
    "last": 165750832,
    "leagues": [
        {
            "id": 198470,
            "specials": [
                {
                    "id": 971955086,
                    "maxBet": 250,
                    "contestantLines": [
                        {
                            "id": 971955113,
                            "lineId": 165750591,
                            "price": 5472,
                            "handicap": null
                        },
                        {
                            "id": 971955105,
                            "lineId": 165750584,
                            "price": 3014,
                            "handicap": null
                        },
                        {
                            "id": 971955096,
                            "lineId": 165750577,
                            "price": 1539,
                            "handicap": null
                        },
                        {
                            "id": 971955101,
                            "lineId": 165750581,
                            "price": 2467,
                            "handicap": null
                        },
                        {
                            "id": 971955092,
                            "lineId": 165750573,
                            "price": 1039,
                            "handicap": null
                        },
                        {
                            "id": 971955111,
                            "lineId": 165750589,
                            "price": 2807,
                            "handicap": null
                        },
                        {
                            "id": 971955093,
                            "lineId": 165750574,
                            "price": 1320,
                            "handicap": null
                        },
                        {
                            "id": 971955107,
                            "lineId": 165750586,
                            "price": 2467,
                            "handicap": null
                        },
                        {
                            "id": 971955104,
                            "lineId": 165750583,
                            "price": 3014,
                            "handicap": null
                        },
                        {
                            "id": 971955095,
                            "lineId": 165750576,
                            "price": 1158,
                            "handicap": null
                        },
                        {
                            "id": 971955089,
                            "lineId": 165750571,
                            "price": 1229,
                            "handicap": null
                        },
                        {
                            "id": 971955100,
                            "lineId": 165750580,
                            "price": 2195,
                            "handicap": null
                        },
                        {
                            "id": 971955114,
                            "lineId": 165750592,
                            "price": 5472,
                            "handicap": null
                        },
                        {
                            "id": 971955108,
                            "lineId": 165750587,
                            "price": 4928,
                            "handicap": null
                        },
                        {
                            "id": 971955099,
                            "lineId": 165750579,
                            "price": 1977,
                            "handicap": null
                        },
                        {
                            "id": 971955110,
                            "lineId": 165750588,
                            "price": 3616,
                            "handicap": null
                        },
                        {
                            "id": 971955087,
                            "lineId": 165750569,
                            "price": 427,
                            "handicap": null
                        },
                        {
                            "id": 971955106,
                            "lineId": 165750585,
                            "price": 3559,
                            "handicap": null
                        },
                        {
                            "id": 971955102,
                            "lineId": 165750582,
                            "price": 2742,
                            "handicap": null
                        },
                        {
                            "id": 971955116,
                            "lineId": 165750593,
                            "price": 5472,
                            "handicap": null
                        },
                        {
                            "id": 971955112,
                            "lineId": 165750590,
                            "price": 4928,
                            "handicap": null
                        },
                        {
                            "id": 971955094,
                            "lineId": 165750575,
                            "price": 612,
                            "handicap": null
                        },
                        {
                            "id": 971955088,
                            "lineId": 165750570,
                            "price": 812,
                            "handicap": null
                        },
                        {
                            "id": 971955090,
                            "lineId": 165750572,
                            "price": 944,
                            "handicap": null
                        },
                        {
                            "id": 971955098,
                            "lineId": 165750578,
                            "price": 1977,
                            "handicap": null
                        }
                    ]
                }
            ]
        },
        {
            "id": 198469,
            "specials": [
                {
                    "id": 971949113,
                    "maxBet": 250,
                    "contestantLines": [
                        {
                            "id": 971949137,
                            "lineId": 165750829,
                            "price": 10300,
                            "handicap": null
                        },
                        {
                            "id": 971949123,
                            "lineId": 165750816,
                            "price": 6318,
                            "handicap": null
                        },
                        {
                            "id": 971949127,
                            "lineId": 165750820,
                            "price": 11422,
                            "handicap": null
                        },
                        {
                            "id": 971949130,
                            "lineId": 165750822,
                            "price": 14257,
                            "handicap": null
                        },
                        {
                            "id": 971949134,
                            "lineId": 165750826,
                            "price": 14257,
                            "handicap": null
                        },
                        {
                            "id": 971949116,
                            "lineId": 165750809,
                            "price": 177,
                            "handicap": null
                        },
                        {
                            "id": 971949120,
                            "lineId": 165750813,
                            "price": 2770,
                            "handicap": null
                        },
                        {
                            "id": 971949124,
                            "lineId": 165750817,
                            "price": 2493,
                            "handicap": null
                        },
                        {
                            "id": 971949118,
                            "lineId": 165750811,
                            "price": 2418,
                            "handicap": null
                        },
                        {
                            "id": 971949132,
                            "lineId": 165750824,
                            "price": 14257,
                            "handicap": null
                        },
                        {
                            "id": 971949136,
                            "lineId": 165750828,
                            "price": 15199,
                            "handicap": null
                        },
                        {
                            "id": 971949122,
                            "lineId": 165750815,
                            "price": 5753,
                            "handicap": null
                        },
                        {
                            "id": 971949126,
                            "lineId": 165750819,
                            "price": 9530,
                            "handicap": null
                        },
                        {
                            "id": 971949138,
                            "lineId": 165750830,
                            "price": 4699,
                            "handicap": null
                        },
                        {
                            "id": 971949119,
                            "lineId": 165750812,
                            "price": 1433,
                            "handicap": null
                        },
                        {
                            "id": 971949115,
                            "lineId": 165750808,
                            "price": 124,
                            "handicap": null
                        },
                        {
                            "id": 971949139,
                            "lineId": 165750831,
                            "price": 15199,
                            "handicap": null
                        },
                        {
                            "id": 971949131,
                            "lineId": 165750823,
                            "price": 9530,
                            "handicap": null
                        },
                        {
                            "id": 971949129,
                            "lineId": 165750821,
                            "price": 11422,
                            "handicap": null
                        },
                        {
                            "id": 971949117,
                            "lineId": 165750810,
                            "price": 382,
                            "handicap": null
                        },
                        {
                            "id": 971949121,
                            "lineId": 165750814,
                            "price": 1773,
                            "handicap": null
                        },
                        {
                            "id": 971949125,
                            "lineId": 165750818,
                            "price": 8583,
                            "handicap": null
                        },
                        {
                            "id": 971949133,
                            "lineId": 165750825,
                            "price": 6341,
                            "handicap": null
                        }
                    ]
                }
            ]
        }
    ]
}`

func TestGetOddsSpecial(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		Equals(t, req.URL.String(), "/v1/odds/special?leagueIds=198470&leagueIds=198474&sportId=33")

		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		rw.Write([]byte(oddsSpecialResp))
	}))
	defer server.Close()

	r := go_pinnacle.NewApiClient(username, pass)
	r.Lines.Request = request.NewRequest("", "", server.URL+"/", server.Client())

	res, err := r.Lines.GetOddsSpecial(&odds.OddsSpecialRequest{
		SportID:   33,
		LeagueIDs: []int{198470, 198474},
	})
	Ok(t, err)
	Equals(t, res.SportID, int64(33))
}

func TestGetLine(t *testing.T) {
	r := go_pinnacle.NewApiClient(username, pass)
	res, err := r.Lines.GetLine(
		1,
		0.5,
		models.Decimal,
		33,
		1,
		0,
		models.MoneyLine,
		&lines.LineRequest{
			Side: models.Under,
			Team: models.Team1,
		},
	)
	Ok(t, err)
	Equals(t, res.Status, "NOT_EXISTS")
}
