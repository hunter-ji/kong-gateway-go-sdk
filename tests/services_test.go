package tests

import (
	"os"
	"testing"

	kongGatewayGoSdk "github.com/Kuari/kong-gateway-go-sdk"
)

func TestServicesAdd(t *testing.T) {
	client := kongGatewayGoSdk.Connect(os.Getenv("KONG_URL"))
	err := client.Services(&kongGatewayGoSdk.Services{Name: "mockServices", Url: "http://mock.com"}).Add()
	if err != nil {
		t.Error(err)
	}
}
