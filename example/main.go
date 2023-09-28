package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	client "github.com/SKF/go-integrations-client"
	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-rest-utility/client/retry"
)

func main() {
	tokenProvider := auth.NewCachedTokenProvider(&auth.CredentialsTokenProvider{
		Username:  "username",
		Password:  "password",
		Endpoint:  "https://sso-api.users.enlight.skf.com/sign-in/initiate",
		TokenType: "identityToken",
		Client:    http.DefaultClient,
		Retry: &retry.ExponentialJitterBackoff{
			Base:        500 * time.Millisecond, //nolint:gomnd
			Cap:         30 * time.Second,       //nolint:gomnd
			MaxAttempts: 30,                     //nolint:gomnd
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := client.New(rest.WithTokenProvider(tokenProvider))

	integrations, err := c.GetIntegrations(ctx)
	if err != nil {
		panic(err)
	}

	for _, integration := range integrations {
		fmt.Printf("%+v\n", integration)
	}
}
