package hwapi

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

type Honeywellapi struct {
	Client *http.Client
	Config *oauth2.Config
}

func New(key, secret string) *Honeywellapi {
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

	return &Honeywellapi{
		Config: conf,
		Client: &http.Client{},
	}
}

// Auth do the real oauth/token auth
func (hw *Honeywellapi) Auth(code, accessToken, refreshToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if accessToken != "" {
		token := &oauth2.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Expiry:       time.Now(),
		}

		client := hw.Config.Client(ctx, token)

		hw.Client = client

		return nil
	}

	if code == "" {
		url := hw.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
		fmt.Printf("Visit the URL for the auth dialog then add --code on the commandline: %v\n", url)

		return errors.New("please provide the `code` on the command line with --code")
	}

	token, err := hw.Config.Exchange(ctx, code)
	if err != nil {
		return err
	}

	fmt.Printf("Please keep your Bearer and Refresh Token for future use: [%v,%v]\n", token.AccessToken, token.RefreshToken)

	client := hw.Config.Client(ctx, token)
	hw.Client = client
	return nil

}

func (hw *Honeywellapi) GetLocation() {
	// the client will update its token if it's expired
	resp, err := hw.Client.Get(hwAPIURL + "/v2/locations?apikey=" + hw.Config.ClientID)
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
