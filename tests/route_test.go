package tests

import (
	"os"
	"testing"

	kongGatewayGoSdk "github.com/Kuari/kong-gateway-go-sdk"
)

func TestRouteAdd(t *testing.T) {
	client := kongGatewayGoSdk.Connect(os.Getenv("KONG_URL"))
	err := client.Route(&kongGatewayGoSdk.Route{
		Name:      "mocking",
		Protocols: &[]string{"http"},
		Paths:     []string{"/mock"},
		Service:   map[string]string{"name": "mockServices"}},
	).Add()
	if err != nil {
		t.Error(err)
	}
}
