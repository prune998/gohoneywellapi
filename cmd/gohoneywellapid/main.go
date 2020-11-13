package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/namsral/flag"
	hwapi "github.com/prune998/gohoneywellapi"
	"github.com/prune998/gohoneywellapi/cmd/gohoneywellapid/handler"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var (
	// version is filled by -ldflags  at compile time
	version        = "no version set"
	displayVersion = flag.Bool("version", false, "Show version and quit")
	logLevel       = flag.String("loglevel", logrus.WarnLevel.String(), "the log level to display (debug,info,error,warning)")

	clientKey    = flag.String("key", "", "API (Consumer) Key from the Developper portal")
	clientSecret = flag.String("secret", "", "API (Consumer) Secret from the Developper Portal")
	clientCode   = flag.String("code", "", "code from the authorization API")

	token        = flag.String("token", "", "your Bearer Token from a previous auth")
	refreshToken = flag.String("refreshtoken", "", "Refresh Token from a previous auth")

	// httpPort   = flag.Int("httpPort", 8080, "Port for service HTTP content")
	configFile = flag.String("confifile", "./config.json", "path of JSON config file")

	// conf ConfigData
)

// type ConfigData struct {
// 	AccessToken  string    `json:"access_token"`
// 	TokenType    string    `json:"token_type,omitempty"`
// 	RefreshToken string    `json:"refresh_token,omitempty"`
// 	Expiry       time.Time `json:"expiry,omitempty"`
// }

func main() {
	flag.Parse()

	// set logging
	logrus.SetOutput(os.Stdout)
	myLogLevel, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		myLogLevel = logrus.WarnLevel
	}
	logrus.SetLevel(myLogLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logger := logrus.WithFields(logrus.Fields{
		"application": "gohoneywellapid",
	})

	if *displayVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	// init oauth and honneywell API
	myHwapi := hwapi.New(*clientKey, *clientSecret)

	// init token
	var tok *oauth2.Token

	// open the config file
	if *configFile != "" {
		tok, err = tokenFromFile(*configFile)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"state": "error",
				"file":  *configFile,
			}).Errorf("error opening config file, continuing with arguments")
		}
		logger.WithFields(logrus.Fields{
			"state": "OK",
			"file":  *configFile,
		}).Infof("done reading json file")
	}

	// use data from config file to auth
	if tok.Valid() {

		logger.WithFields(logrus.Fields{
			"state":  "OK",
			"expire": tok.Expiry,
		}).Infof("using token from config file")

		err = myHwapi.AuthFromToken(tok)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"state": "error",
			}).Errorf("error during auth : %v", err)
			os.Exit(1)
		}
	} else {

		logger.WithFields(logrus.Fields{
			"state": "OK",
		}).Infof("using token from command line")

		// do auth
		tok, err = myHwapi.Auth(*clientCode, *token, *refreshToken)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"state": "error",
			}).Errorf("error during auth : %v", err)
			os.Exit(1)
		}
		saveToken(*configFile, tok)
	}

	logger.WithFields(logrus.Fields{
		"state": "OK",
	}).Infof("oAuth2 done, type anything")

	e := echo.New()

	e.HideBanner = true
	e.Debug = true
	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Initialize handler
	h := &handler.Handler{Hw: myHwapi}

	// routes
	apiGroup := e.Group("/api")
	apiGroup.GET("/locations", h.GetLocations).Name = "get-locations"
	apiGroup.GET("/location/:locationid", h.GetLocation).Name = "get-locations"
	apiGroup.GET("/location/:locationid/devices", h.GetDevices).Name = "get-devices"
	apiGroup.GET("/location/:locationid/device/:deviceid", h.GetDevice).Name = "get-device"
	// apiGroup.GET("/location/:locationid/device/:deviceid/schedule", h.GetSchedule).Name = "get-schedule"

	e.Static("/", "frontend/dist")

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)

	for {
		var code string
		fmt.Scan(&code)

		locations, err := myHwapi.GetLocations()
		if err != nil {
			logger.WithFields(logrus.Fields{
				"state": "error",
				"error": err,
			}).Errorf("error getting Locations")
		}

		for _, j := range locations {
			fmt.Println(j.Name)
			fmt.Println(len(j.Devices))

			for _, device := range j.Devices {
				fmt.Printf("temp for %s: in=%.2f | out=%.2f\n", device.Name, device.IndoorTemperature, device.OutdoorTemperature)
				fmt.Printf("humidity for %s: in=%d | out=%d\n", device.Name, device.IndoorHumidity, device.DisplayedOutdoorHumidity)

				s, err := myHwapi.GetSchedule(strconv.Itoa(j.LocationID), device.DeviceID)
				if err != nil {
					logger.WithFields(logrus.Fields{
						"state":      "error",
						"locationID": j.LocationID,
						"deviceID":   device.DeviceID,
						"error":      err,
					}).Errorf("error getting Device Schedule")
					continue
				}
				for _, daySchedule := range s.TimedSchedule.Days {
					fmt.Printf("\t%s:\n", daySchedule.Day)
					for _, period := range daySchedule.Periods {
						if !period.IsCancelled {
							fmt.Printf("\t\t%s -> %.2f\n", period.StartTime, Fahrenheit2Celsius(period.HeatSetPoint))
						}
					}
				}
			}
		}
		logger.WithFields(logrus.Fields{
			"state": "OK",
			"error": err,
		}).Infof("done calling API")
	}
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	_ = json.NewEncoder(f).Encode(token)
}

// Celsius2Fahrenheit
func Celsius2Fahrenheit(c float64) float64 {
	return (c*9/5 + 32)
}

// Fahrenheit2Celsius
func Fahrenheit2Celsius(f float64) float64 {
	return ((f - 32) * 5 / 9)
}
