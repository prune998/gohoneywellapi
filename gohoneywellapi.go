package gohoneywellapi

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const hwAPIURL string = "https://api.honeywell.com"

type honeywellapi struct {
	client *http.Client
	config *oauth2.Config
}

// NewHW connect to Honeywell Developper API
// if the token is provided, just use it
func NewHW(key, secret, code, accessToken, refreshToken string) (*honeywellapi, error) {
	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Scopes:       []string{"none"},
		RedirectURL:  "none",
		Endpoint: oauth2.Endpoint{
			AuthURL:   hwAPIURL + "/oauth2/authorize",
			TokenURL:  hwAPIURL + "/oauth2/token",
			AuthStyle: 2,
		},
	}

	if accessToken != "" {
		token := &oauth2.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Expiry:       time.Now(),
		}

		client := conf.Client(ctx, token)

		return &honeywellapi{
			client: client,
			config: conf,
		}, nil
	}

	if code == "" {
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		fmt.Printf("Visit the URL for the auth dialog then add --code on the commandline: %v\n", url)

		return nil, errors.New("please provide the `code` on the command line with --code")
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Please keep your Bearer and Refresh Token for future use: [%v,%v]\n", token.AccessToken, token.RefreshToken)

	client := conf.Client(ctx, token)

	return &honeywellapi{
		client: client,
		config: conf,
	}, nil

}

func (hw *honeywellapi) GetLocation() {
	// the client will update its token if it's expired
	resp, err := hw.client.Get(hwAPIURL + "/v2/locations?apikey=" + hw.config.ClientID)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// If response code is 200 it was successful
	if resp.StatusCode == 200 {
		fmt.Println("The request was successful. Response below:")
		fmt.Println(string(body))
	} else {
		fmt.Println("Could not perform request to the endpoind. Response below:")
		fmt.Println(string(body))
	}
}
