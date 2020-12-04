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
	Username string    `json:"username"`
	CustomId string    `json:"custom_id"`
	Tags     *[]string `json:"tags,omitempty"`
}

func (c *Config) Consumer(consumer *Consumer) *consumerBody {
	consumerByte, _ := json.Marshal(consumer)
	return &consumerBody{
		Config: c,
		Body:   consumerByte,
	}
}

func (c *consumerBody) Add() (err error) {
	url := fmt.Sprintf("%s/consumers", c.Config.Url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(c.Body))
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
