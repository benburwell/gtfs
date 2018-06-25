package gtfs

import "strconv"

// A LocationType distinguishes a stop, station, or station entrance.
type LocationType int

const (
	// LocationTypeStop is a location where passengers board or disembark from a
	// transit vehicle.
	LocationTypeStop LocationType = 0

	// LocationTypeStation is a physical structure or area that contains one or
	// more stop.
	LocationTypeStation = 1

	// LocationTypeStationEntranceExit is a location where passengers can enter
	// or exit a station from the street.
	LocationTypeStationEntranceExit = 2
)

// MarshalCSV converts the location type to a CSV string
func (l *LocationType) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*l)), nil
}

// UnmarshalCSV converts the CSV string to the location type representation
func (l *LocationType) UnmarshalCSV(csv string) (err error) {
	d, err := strconv.Atoi(csv)
	n := LocationType(d)
	l = &n
	return err
}

func (l *LocationType) String() string {
	return strconv.Itoa(int(*l))
}
