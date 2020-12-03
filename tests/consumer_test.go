package tests

import (
	"os"
	"testing"

	kongGatewayGoSdk "github.com/Kuari/kong-gateway-go-sdk"
)

func TestConsumerAdd(t *testing.T) {
	client := kongGatewayGoSdk.Connect(os.Getenv("KONG_URL"))
	err := client.Consumer(&kongGatewayGoSdk.Consumer{
		Username: "jack",
		CustomId: "jack_id",
		Tags:     &[]string{"dev"},
	}).Add()
	if err != nil {
		t.Error(err)
	}
}
