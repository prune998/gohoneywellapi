package main

import (
	"fmt"
	"os"

	"github.com/namsral/flag"
	"github.com/prune998/gohoneywellapi"
	"github.com/sirupsen/logrus"
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

	logger.WithFields(logrus.Fields{
		"state": "OK",
	}).Infof("app started, authenticating")

	// do auth
	hwapi, err := gohoneywellapi.NewHW(*clientKey, *clientSecret, *clientCode, *token, *refreshToken)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
		}).Errorf("error during auth : %v", err)
		os.Exit(1)
	}

	for {
		var code string
		fmt.Scan(&code)
		hwapi.GetLocation()
		logger.WithFields(logrus.Fields{
			"state": "OK",
		}).Infof("done calling oauth API")
	}
}
