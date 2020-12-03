package kong_gateway_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RouteBody struct {
	*Config
	Services map[string]string
	Body     []byte
}

type Route struct {
	Name                    string            `json:"name,omitempty"`
	Protocols               *[]string         `json:"protocols,omitempty"`
	Methods                 *[]string         `json:"methods,omitempty"`
	Hosts                   *[]string         `json:"hosts,omitempty"`
	Paths                   []string          `json:"paths,omitempty"`
	Headers                 *[]string         `json:"headers,omitempty"`
	HttpsRedirectStatusCode string            `json:"https_redirect_status_code,omitempty"`
	RegexPriority           *int              `json:"regex_priority,omitempty"`
	StripPath               *bool             `json:"strip_path,omitempty"`
	PathHandling            string            `json:"path_handling,omitempty"`
	PreserveHost            *bool             `json:"preserve_host,omitempty"`
	RequestBuffering        *bool             `json:"request_buffering,omitempty"`
	ResponseBuffering       *bool             `json:"response_buffering,omitempty"`
	Snis                    *[]string         `json:"snis,omitempty"`
	Sources                 *[]string         `json:"sources,omitempty"`
	Destinations            string            `json:"destinations,omitempty"`
	Tags                    *[]string         `json:"tags,omitempty"`
	Service                 map[string]string `json:"service"`
}

func (c *Config) Route(route *Route) *RouteBody {
	routeByte, _ := json.Marshal(route)
	return &RouteBody{
		Config:   c,
		Services: route.Service,
		Body:     routeByte,
	}
}

func (r *RouteBody) Add() (err error) {

	var services string
	if r.Services["id"] != "" {
		services = r.Services["id"]
	} else if r.Services["name"] != "" {
		services = r.Services["name"]
	} else {
		err = errors.New("services参数缺失")
		return
	}

	url := fmt.Sprintf("%s/services/%s/routes", r.Config.Url, services)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(r.Body))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.Status != "201 Created" {
		body, bodyErr := ioutil.ReadAll(resp.Body)
		if bodyErr != nil {
			err = bodyErr
			return
		}
		err = errors.New(string(body))
	}

	return
}
