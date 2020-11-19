package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/imroc/req"
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

	serverBind = flag.String("server", ":8080", "Server:Port for HTTP content")
	configFile = flag.String("configfile", "./config.json", "path of JSON config file")
	tokenFile  = flag.String("tokenfile", "./token.json", "path of JSON file containing the saved token")

	peakHoursURL = flag.String("peakhoururl", "http://localhost:8022/peakhours", "URL to get the Peak Hours JSON content")
)

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

	// load the config file
	config, err := parseConfig(*configFile)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
			"err":   err,
		}).Errorf("can't load config file %s", *configFile)
		os.Exit(1)
	}

	// Query the PeakHour service
	peaks, err := getPeakHours(*peakHoursURL)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
			"err":   err,
		}).Errorf("can't grab Peak Hours from server URL %s", *peakHoursURL)
		os.Exit(1)
	}

	logger.WithFields(logrus.Fields{
		"state": "OK",
		"peaks": peaks,
		"conf":  config,
	}).Infof("done reading all files")

	// init oauth and honeywellAPI
	myHwapi := hwapi.New(*clientKey, *clientSecret)

	// init token
	var tok *oauth2.Token

	// grab the last valid token from the backup file
	if *tokenFile != "" {
		tok, err = tokenFromFile(*tokenFile)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"state": "error",
				"file":  *tokenFile,
			}).Errorf("error opening token file, continuing with arguments")
		}
		logger.WithFields(logrus.Fields{
			"state": "OK",
			"file":  *tokenFile,
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
		saveToken(*tokenFile, tok)
	}

	logger.WithFields(logrus.Fields{
		"state": "OK",
	}).Infof("oAuth2 done, type anything")

	e := echo.New()

	e.HideBanner = true
	e.Debug = true
	// Enable metrics middleware
	p := prometheus.NewPrometheus("hwapi", nil)
	p.Use(e)

	// Initialize handler
	h := &handler.Handler{Hw: myHwapi}

	// routes
	apiGroup := e.Group("/api")
	apiGroup.GET("/locations", h.GetLocations).Name = "get-locations"
	apiGroup.GET("/location/:locationid", h.GetLocation).Name = "get-location"
	apiGroup.GET("/location/:locationid/devices", h.GetDevices).Name = "get-devices"
	apiGroup.GET("/location/:locationid/device/:deviceid", h.GetDevice).Name = "get-device"
	apiGroup.GET("/location/:locationid/device/:deviceid/schedule", h.GetSchedule).Name = "get-schedule"

	e.Static("/", "frontend/dist")

	// Start server
	go func() {
		if err := e.Start(*serverBind); err != nil {
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
				fmt.Printf("humidity for %s: in=%.2f | out=%.2f\n", device.Name, device.IndoorHumidity, device.DisplayedOutdoorHumidity)

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
							fmt.Printf("\t\t%s\t%s -> %.2f\n", period.PeriodName, period.StartTime, Fahrenheit2Celsius(period.HeatSetPoint))
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

// getPeakHours do a REST call to get the peak hours json
func getPeakHours(url string) ([]PeakHourPeriod, error) {

	var s []PeakHourPeriod

	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	err = r.ToJSON(&s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

type PeakHourPeriod struct {
	StartTime time.Time `json:"start"`
	EndTime   time.Time `json:"end"`
}
