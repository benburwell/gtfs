package gtfs

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gocarina/gocsv"
)

// A Feed is a complete GTFS database
type Feed struct {
	Agencies       []Agency
	Stops          []Stop
	Routes         []Route
	Trips          []Trip
	StopTimes      []StopTime
	Calendars      []Calendar
	CalendarDates  []CalendarDate
	FareAttributes []FareAttribute
	FareRules      []FareRule
	Shapes         []Shape
	Frequencies    []Frequency
	Transfers      []Transfer
	FeedInfos      []FeedInfo
}

// NewFeed reads a ZIP archived GTFS feed and returns a Feed
func NewFeed(r io.Reader) (*Feed, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("could not read GTFS archive: %v", err)
	}
	zipReader, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, fmt.Errorf("could not read zip archive: %v", err)
	}
	feed := &Feed{}
	for _, zipFile := range zipReader.File {
		rc, err := zipFile.Open()
		if err != nil {
			return nil, fmt.Errorf("could not open zip file: %v", err)
		}
		defer rc.Close()
		switch zipFile.Name {
		case "agency.txt":
			err = gocsv.Unmarshal(rc, &feed.Agencies)
		case "stops.txt":
			err = gocsv.Unmarshal(rc, &feed.Stops)
		case "routes.txt":
			err = gocsv.Unmarshal(rc, &feed.Routes)
		case "trips.txt":
			err = gocsv.Unmarshal(rc, &feed.Trips)
		case "stop_times.txt":
			err = gocsv.Unmarshal(rc, &feed.StopTimes)
		case "calendar.txt":
			err = gocsv.Unmarshal(rc, &feed.Calendars)
		case "calendar_dates.txt":
			err = gocsv.Unmarshal(rc, &feed.CalendarDates)
		case "fare_attributes.txt":
			err = gocsv.Unmarshal(rc, &feed.FareAttributes)
		case "shapes.txt":
			err = gocsv.Unmarshal(rc, &feed.Shapes)
		case "frequencies.txt":
			err = gocsv.Unmarshal(rc, &feed.Frequencies)
		case "transfers.txt":
			err = gocsv.Unmarshal(rc, &feed.Transfers)
		case "feed_info.txt":
			err = gocsv.Unmarshal(rc, &feed.FeedInfos)
		}
		if err != nil {
			return nil, err
		}
	}
	return feed, nil
}
