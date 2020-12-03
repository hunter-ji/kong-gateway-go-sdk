# kong-gateway-go-sdk



## Intro

Use the api of kong gateway encapsulated by golang.The function is under development.



## Usage

```go
package main

import (
	"os"

	gatewayGoSdk "github.com/Kuari/kong-gateway-go-sdk"
)

func main() {

	// Connect the gateway
	client := gatewayGoSdk.Connect(os.Getenv("KONG_URL"))

	// Add a Service
	err := client.Services(&gatewayGoSdk.Services{Name: "example_service", Url: "http://mockbin.org"}).Add()
	if err != nil {
		panic("Failed to add services.")
	}

	// Add a Route
	err = client.Route(&gatewayGoSdk.Route{
		Name: "mocking", Paths: []string{"/mock"}, Service: map[string]string{"name": "mockServices"}}).Add()
	if err != nil {
		panic("Failed to add route.")
	}

	// Add a Consumer
	err = client.Consumer(&gatewayGoSdk.Consumer{Username: "tom", CustomId: "tom_id"}).Add()
	if err != nil {
		panic("Failed to add consumer.")
	}

}
```



## Features

The currently completed functions are as follows:

| Module                  | Method                |
| ----------------------- | --------------------- |
| Services                | `Add`                 |
| Consumer                | `Add`                 |
| PluginKeyAuthentication | `CreateKey`、`DelKey` |

