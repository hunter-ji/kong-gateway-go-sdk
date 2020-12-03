package tests

import (
	"os"
	"testing"

	kongGatewayGoSdk "github.com/Kuari/kong-gateway-go-sdk"
)

func TestPluginAuthKeyCreate(t *testing.T) {
	client := kongGatewayGoSdk.Connect(os.Getenv("KONG_URL"))
	err := client.PluginKeyAuth(&kongGatewayGoSdk.PluginKeyAuth{Consumer: "jack"}).CreateKey()
	if err != nil {
		t.Error(err)
	}
}

func TestPluginAuthKeyDel(t *testing.T) {
	client := kongGatewayGoSdk.Connect(os.Getenv("KONG_URL"))
	err := client.PluginKeyAuth(&kongGatewayGoSdk.PluginKeyAuth{
		Consumer: "jack", Key: "T7DpeSP6AWQPX4Q4chOHjGYZOxkbhyHv"}).DelKey()
	if err != nil {
		t.Error(err)
	}
}
