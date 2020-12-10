package kong_gateway_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type consumerBody struct {
	*Config
	Body []byte
}

type Consumer struct {
	Username   string    `json:"username,omitempty"`
	CustomId   string    `json:"custom_id,omitempty"`
	ConsumerId string    `json:"consumer_id,omitempty"`
	Tags       *[]string `json:"tags,omitempty"`
}

func (c *Config) Consumer(consumer *Consumer) *consumerBody {
	consumerByte, _ := json.Marshal(consumer)
	return &consumerBody{
		Config: c,
		Body:   consumerByte,
	}
}

func (c *consumerBody) Add() (id string, err error) {

	type Response struct {
		Id string `json:"id"`
	}

	url := fmt.Sprintf("%s/consumers", c.Config.Url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(c.Body))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 201 {
		err = errors.New(string(body))
		return
	}

	var res Response
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}

	id = res.Id
	return
}

func (c *consumerBody) Retrieve() (exist bool, err error) {

	var consumerInfo Consumer
	err = json.Unmarshal(c.Body, &consumerInfo)
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/consumers/%s", c.Config.Url, consumerInfo.ConsumerId)
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		exist = false
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(string(body))
		return
	}

	exist = true
	return
}
