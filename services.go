package kong_gateway_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type servicesBody struct {
	*Config
	Body []byte
}

type Services struct {
	Name              string         `json:"name"`
	Retries           *int           `json:"retries,omitempty"`
	Protocol          *[]string      `json:"protocol,omitempty"`
	Host              string         `json:"host,omitempty"`
	Port              *int           `json:"port,omitempty"`
	path              string         `json:"path,omitempty"`
	ConnectTimeout    *int           `json:"connect_timeout,omitempty"`
	WriteTimeout      *int           `json:"write_timeout,omitempty"`
	ReadTimeout       *int           `json:"read_timeout,omitempty"`
	Tags              *[]string      `json:"tags,omitempty"`
	ClientCertificate interface{}    `json:"client_certificate,omitempty"`
	TlsVerify         interface{}    `json:"tls_verify,omitempty"`
	TlsVerifyDepth    interface{}    `json:"tls_verify_depth,omitempty"`
	CaCertificates    *[]interface{} `json:"ca_certificates,omitempty"`
	Url               string         `json:"url,omitempty"`
}

func (c *Config) Services(services *Services) *servicesBody {

	servicesByte, _ := json.Marshal(services)
	return &servicesBody{
		Config: c,
		Body:   servicesByte,
	}
}

func (r *servicesBody) Add() (err error) {

	resp, err := http.Post(r.Config.Url+"/services", "application/json", bytes.NewBuffer(r.Body))
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
