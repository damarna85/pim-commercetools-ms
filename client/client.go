package client

import (
	"context"
	"os"
	"strings"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"golang.org/x/oauth2/clientcredentials"
)

var cl *commercetools.Client

// GetClient gets existing client
func GetClient() *commercetools.Client {
	if cl == nil {
		auth := &clientcredentials.Config{
			ClientID:     os.Getenv("CTP_CLIENT_ID"),
			ClientSecret: os.Getenv("CTP_CLIENT_SECRET"),
			Scopes:       strings.Split(os.Getenv("CTP_SCOPES"), ","),
			TokenURL:     os.Getenv("CTP_AUTH_URL"),
		}
		cl = commercetools.New(&commercetools.Config{
			ProjectKey: os.Getenv("CTP_PROJECT_KEY"),
			HTTPClient: auth.Client(context.TODO()),
			URL:        os.Getenv("CTP_API_URL"),
		})
	}
	return cl
}
