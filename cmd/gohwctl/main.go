package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/namsral/flag"
	"github.com/sirupsen/logrus"
)

var (
	// version is filled by -ldflags  at compile time
	version        = "no version set"
	displayVersion = flag.Bool("version", false, "Show version and quit")
	logLevel       = flag.String("loglevel", logrus.WarnLevel.String(), "the log level to display (debug,info,error,warning)")

	secretKey = flag.String("key", "", "key used to talk to the server")

	startDate = flag.String("startdate", "", "start date for the change, like <year>-<month>-<day>T<hour>:<min>:<sec>")
	temp      = flag.Float64("temp", 20.0, "Temp to add")
	tempUnit  = flag.String("unit", "celcius", "temperature unit")

	action = flag.String("action", "add", "action like add, del")
	apply  = flag.Bool("y", false, "apply without asking")

	serverURL = flag.String("server", "http://localhost", "URL to connect to the server")
)

type schedule struct {
	startDate time.Time `json:"startdate"`
	temp      float64   `json:"temp"`
	tempUnit  string    `json:"tempunit"`
}

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
		"application": "gohwctl",
	})

	if *displayVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	parsedStartDate, err := time.Parse("2006-01-02T15:04:05", *startDate)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
			"err":   err,
		}).Fatalf("error converting Start Time")
	}
	// shift the time to the local timezone
	localStartDate := parsedStartDate.In(time.Local)

	// init current schedule
	sched := &schedule{
		startDate: localStartDate,
		temp:      *temp,
		tempUnit:  *tempUnit,
	}

	logger.WithFields(logrus.Fields{
		"state": "ok",
	}).Infof("setting temp to %.2f %s starting %v", *temp, *tempUnit, localStartDate)

	if !*apply {
		var code string
		fmt.Printf("setting temp to %.2f %s starting %v\n", *temp, *tempUnit, localStartDate)
		fmt.Println("enter 'y' to confirm")
		fmt.Scan(&code)
		if code != "y" {
			return
		}
	}

	// create a new HTTP client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	logger.WithFields(logrus.Fields{
		"state": "ok",
		"temp":  *temp,
	}).Infof("applied temp at %.2f%s starting %v", *temp, *tempUnit, localStartDate)

	// set the HTTP action
	var actionType string
	switch a := *action; a {
	case "add":
		actionType = "POST"
	case "del":
		actionType = "DELETE"
	default:
		logger.WithFields(logrus.Fields{
			"state":  "error",
			"action": *action,
		}).Fatalln("only 'add' and 'del' actions are supported")
	}

	// create request body
	s, _ := json.Marshal(sched)
	b := bytes.NewBuffer(s)

	// create a request object
	req, _ := http.NewRequest(
		actionType,
		*serverURL+"/api/v1/create",
		b,
	)

	// add a request header
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X_HWAPI_SECRET", *secretKey)

	// send an HTTP using `req` object
	res, err := client.Do(req)

	if err != nil {

		// get `url.Error` struct pointer from `err` interface
		urlErr := err.(*url.Error)

		// check if error occurred due to timeout
		if urlErr.Timeout() {
			logger.WithFields(logrus.Fields{
				"state": "error",
				"err":   err,
			}).Fatalln("Error occurred due to a timeout.")
		}

		// log error and exit
		logger.WithFields(logrus.Fields{
			"state": "error",
			"err":   err,
		}).Fatalln("Error connecting to server.")
	} else {
		fmt.Println("Success: status-code", res.StatusCode)
	}

}
