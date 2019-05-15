package request

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"go-pinnacle/models"
	"gopkg.in/resty.v1"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

type Request interface {
	Send(target interface{}, requestMethod string, method string, query interface{}, data interface{}) error
}

type request struct {
	request *resty.Client
	url     string
}

func (r *request) Send(target interface{}, requestMethod string, method string, params interface{}, data interface{}) error {
	req := r.request.R()
	var err error

	if params != nil {
		var v = make(url.Values)

		if reflect.TypeOf(params) == reflect.TypeOf(url.Values{}) {
			v = params.(url.Values)
		} else {
			v, err = query.Values(params)
			if err != nil {
				return err
			}
		}

		req.SetMultiValueQueryParams(v)
	}

	if data != nil {
		req.SetBody(data)
	}

	resp, err := req.Execute(requestMethod, r.url+method)
	if err != nil {
		return err
	}
	body := resp.Body()

	if resp.StatusCode() != 200 {
		var errorData = new(models.PinnacleError)

		err = json.Unmarshal(body, errorData)
		if err != nil {
			println("ERROR")
			return err
		}
		return errorData
	}

	return json.Unmarshal(body, target)
}

func NewRequest(username, pass string, hostUrl string, httpClient *http.Client) Request {
	var r *resty.Client

	if httpClient != nil {
		r = resty.NewWithClient(httpClient)
	} else {
		r = resty.New()
	}

	r.
		SetBasicAuth(username, pass).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		})

	if os.Getenv("DEBUG") != "" {
		r.SetDebug(true)
	}

	return &request{request: r, url: hostUrl}
}
