package kong_gateway_go_sdk

import (
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

func (p *PluginKeyAuthBody) CreateKey() (err error) {
	url := fmt.Sprintf("%s/consumers/%s/key-auth", p.Config.Url, p.Consumer)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
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

func (p *PluginKeyAuthBody) DelKey() (err error) {
	url := fmt.Sprintf("%s/consumers/%s/key-auth/%s", p.Config.Url, p.Consumer, p.Key)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status != "204 No Content" {
		body, bodyErr := ioutil.ReadAll(resp.Body)
		if bodyErr != nil {
			err = bodyErr
			return
		}
		err = errors.New(string(body))
	}

	return
}
