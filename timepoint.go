package gtfs

import "strconv"

// Timepoint identifies whether bikes are permitted
type Timepoint int

const (
	// TimepointApproximate indicates that an arrival/departure time may be an
	// approximation or interpolation
	TimepointApproximate Timepoint = 0

	// TimepointExact indicates that an arrival/departure time is considered
	// exact.
	TimepointExact = 1
)

// MarshalCSV converts the timepoint to a CSV string
func (t *Timepoint) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*t)), nil
}

// UnmarshalCSV converts the CSV string to the service type representation
func (t *Timepoint) UnmarshalCSV(csv string) error {
	d, err := strconv.Atoi(csv)
	n := Timepoint(d)
	t = &n
	return err
}

// String converts the location type to a string
func (t *Timepoint) String() string {
	return strconv.Itoa(int(*t))
}
