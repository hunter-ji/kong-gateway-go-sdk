package kong_gateway_go_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PluginKeyAuthBody struct {
	*Config
	Consumer string
	Key      string
}

type PluginKeyAuth struct {
	Consumer string
	Key      string
}

func (c *Config) PluginKeyAuth(auth *PluginKeyAuth) *PluginKeyAuthBody {
	return &PluginKeyAuthBody{
		Config:   c,
		Consumer: auth.Consumer,
		Key:      auth.Key,
	}
}

func (p *PluginKeyAuthBody) CreateKey() (key string, err error) {

	type Response struct {
		Key string `json:"key"`
	}

	url := fmt.Sprintf("%s/consumers/%s/key-auth", p.Config.Url, p.Consumer)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.Status != "201 Created" {
		err = errors.New(string(body))
		return
	}

	var res Response
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}

	key = res.Key
	return
}

func (p *PluginKeyAuthBody) DelKey() (err error) {
	url := fmt.Sprintf("%s/consumers/%s/key-auth/%s", p.Config.Url, p.Consumer, p.Key)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.Status != "204 No Content" {
		body, bodyErr := ioutil.ReadAll(resp.Body)
		if bodyErr != nil {
			err = bodyErr
			return
		}
		err = errors.New(string(body))
		return
	}

	return
}
