package gtfs

import "strconv"

// WheelchairAccess identifies whether wheelchair boardings are possible
// from the specified stop, station, or station entrance.
type WheelchairAccess int

const (
	// WheelchairAccessUnknown indicates that there is no accessibility
	// information for the stop
	WheelchairAccessUnknown WheelchairAccess = 0

	// WheelchairAccessPossible indicates that at least some vehicles at this
	// stop can be boarded by a rider in a wheelchair
	WheelchairAccessPossible = 1

	// WheelchairAccessNotPossible indicates wheelchair boarding is not
	// possible at this stop
	WheelchairAccessNotPossible = 2
)

// MarshalCSV converts the location type to a CSV string
func (w *WheelchairAccess) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*w)), nil
}

// UnmarshalCSV converts the CSV string to the location type representation
func (w *WheelchairAccess) UnmarshalCSV(csv string) error {
	d, err := strconv.Atoi(csv)
	n := WheelchairAccess(d)
	w = &n
	return err
}

// String converts the location type to a string
func (w *WheelchairAccess) String() string {
	return strconv.Itoa(int(*w))
}
