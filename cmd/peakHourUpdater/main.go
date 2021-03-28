package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	startDate    = flag.String("startDate", "", "start date, formated: <year>-<month>-<day>T<hour>:<min>:<sec>")
	endDate      = flag.String("endDate", "", "end date, formated: <year>-<month>-<day>T<hour>:<min>:<sec>")
	calendarFile = flag.String("datafile", "calendar.json", "file with the current calendar")
)

const appName = "peakHourUpdater"

type schedule struct {
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}

type calendar []schedule

func main() {
	flag.Parse()

	// set logging and version
	logrus.SetOutput(os.Stdout)
	myLogLevel, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		myLogLevel = logrus.WarnLevel
	}
	logrus.SetLevel(myLogLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logger := logrus.WithFields(logrus.Fields{
		"application": appName,
		"version":     version,
	})

	if *displayVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	// parse the start date
	start, err := parseDate(startDate)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
		}).Errorf("Unable to parse Start Date %s", *startDate)
		os.Exit(1)
	}

	// parse the end date
	end, err := parseDate(endDate)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
		}).Errorf("Unable to parse Start Date %s", *endDate)
		os.Exit(1)
	}

	// init current schedule
	sched := &schedule{
		StartDate: start,
		EndDate:   end,
	}

	// TODO: git pull

	// load the old file
	cal, err := loadCalendar(*calendarFile)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
			"err":   err,
		}).Errorf("Unable to process calendar file from %s", *calendarFile)
		os.Exit(1)
	}

	// Add the new schedule to the calendar
	// TODO: dedup similar items
	cal = append(cal, *sched)

	// write the modified calendar file
	err = writeCalendar(cal)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"state": "error",
			"err":   err,
		}).Errorf("Unable to process calendar file from %s", *calendarFile)
		os.Exit(1)
	}

	fmt.Println(cal)

	// TODO: create branch
	// TODO: git commit
	// TODO: git push
	// TODO: git PR
	// TODO: merge PR
}

func parseDate(d *string) (time.Time, error) {
	parsedStartDate, err := time.Parse("2006-01-02T15:04:05", *d)
	if err != nil {
		return time.Time{}, err
	}
	// shift the time to the local timezone
	localStartDate := parsedStartDate.In(time.Local)

	return localStartDate, nil
}

func loadCalendar(file string) (calendar, error) {
	// parse config file
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// we initialize our Users array
	var cal calendar

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &cal)
	if err != nil {
		return nil, err
	}

	return cal, nil
}

// writeCalendar write the final calendar into the file
func writeCalendar(cal calendar) error {

	outCal, err := json.MarshalIndent(cal, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(*calendarFile, outCal, 0644)
	if err != nil {
		return err
	}
}
