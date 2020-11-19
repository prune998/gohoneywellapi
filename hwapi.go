package hwapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const hwAPIURL string = "https://api.honeywell.com"

type HoneywellAPI struct {
	Client *http.Client
	Config *oauth2.Config
}

// New create a new HoneywellAPI resource
func New(key, secret string) *HoneywellAPI {
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

	return &HoneywellAPI{
		Config: conf,
		Client: &http.Client{},
	}
}

// AuthFromToken create the http client with a provided token
func (hw *HoneywellAPI) AuthFromToken(token *oauth2.Token) error {
	ctx := context.Background()

	client := hw.Config.Client(ctx, token)
	hw.Client = client

	return nil
}

// Auth do the real oauth/token auth
func (hw *HoneywellAPI) Auth(code, accessToken, refreshToken string) (*oauth2.Token, error) {
	ctx := context.Background()

	if accessToken != "" && refreshToken != "" {
		token := &oauth2.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Expiry:       time.Now(),
			// token.TokenType = "",
		}

		client := hw.Config.Client(ctx, token)
		hw.Client = client

		return token, nil
	}

	if code == "" {
		url := hw.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
		fmt.Printf("Visit the URL for the auth dialog then add --code on the commandline: %v\n", url)

		return &oauth2.Token{}, errors.New("please provide the `code` on the command line with --code")
	}

	token, err := hw.Config.Exchange(ctx, code)
	if err != nil {
		return token, err
	}

	fmt.Printf("Please keep your Bearer and Refresh Token for future use: [%v,%v, %v ,%v]\n", token.AccessToken, token.RefreshToken, token.Expiry, token.TokenType)

	client := hw.Config.Client(ctx, token)
	hw.Client = client
	return token, nil

}

func (hw *HoneywellAPI) GetLocations() ([]TSerie, error) {
	// the client will update its token if it's expired
	url := hwAPIURL + "/v2/locations?apikey=" + hw.Config.ClientID

	body, err := hw.getData(url)
	if err != nil {
		return nil, err
	}

	// put data into a Tserie struct
	var m []TSerie
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (hw *HoneywellAPI) GetSchedule(locationID, deviceID string) (*Schedule, error) {
	// the client will update its token if it's expired
	url := hwAPIURL + "/v2/devices/schedule/" + deviceID + "?apikey=" + hw.Config.ClientID + "&type=regular&locationId=" + locationID

	body, err := hw.getData(url)
	if err != nil {
		return nil, err
	}

	// put data into a Schedule struct
	var s Schedule
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// getData effectively GET the remote API
func (hw *HoneywellAPI) getData(url string) ([]byte, error) {
	// the client will update its token if it's expired
	resp, err := hw.Client.Get(url)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// If response code is 200 it was successful
	if resp.StatusCode == 200 {
		return body, nil
	}

	return nil, errors.New("HTTP error")
}
