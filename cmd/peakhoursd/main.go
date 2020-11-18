package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/namsral/flag"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var (
	// version is filled by -ldflags  at compile time
	version        = "no version set"
	displayVersion = flag.Bool("version", false, "Show version and quit")
	logLevel       = flag.String("loglevel", logrus.WarnLevel.String(), "the log level to display (debug,info,error,warning)")

	serverBind = flag.String("server", ":8022", "Server:Port for HTTP content")
	configFile = flag.String("confifile", "./peakhours.json", "path of JSON config file")
)

// PeakHour is a start/end date block
type PeakHour struct {
	StartTime time.Time `json:"start,omitempty"`
	EndTime   time.Time `json:"end,omitempty"`
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
		"application": "gohoneywellapid",
	})

	if *displayVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	now := time.Now()
	filterdPeakHours := []PeakHour{}

	// parse config file
	f, err := os.Open(*configFile)
	defer f.Close()

	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
			"err":   err,
		}).Fatalf("error reading Config file")
	}
	dec := json.NewDecoder(f)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
			"err":   err,
		}).Fatalf("error parsing Config file")
	}

	for dec.More() {
		var m PeakHour
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		// only keep peakhours that are still valid now
		if m.EndTime.After(now) {
			filterdPeakHours = append(filterdPeakHours, m)
		}
	}

	logger.WithFields(logrus.Fields{
		"state": "OK",
	}).Infof("Config is OK, starting server...")

	// instanciate http framework
	e := echo.New()
	e.HideBanner = true
	e.Debug = true

	// Middleware
	e.Use(middleware.Logger())
	// Enable metrics middleware
	p := prometheus.NewPrometheus("hwapi", nil)
	p.Use(e)

	// Initialize handler
	h := &Handler{ph: filterdPeakHours}

	// Routes
	e.GET("/peakhours", h.peakHours)

	// Start server
	go func() {
		if err := e.Start(*serverBind); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// Handler is used to pass peakHours into Echo framework handlers
type Handler struct {
	ph []PeakHour
}

// peakHours return only future peak hours blocks
func (h *Handler) peakHours(c echo.Context) error {
	return c.JSON(http.StatusOK, h.ph)
}
