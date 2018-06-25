package gtfs

import "strconv"

// BikeAccess identifies whether bikes are permitted
type BikeAccess int

const (
	// BikeAccessUnknown indicates that there is no bike information available
	BikeAccessUnknown BikeAccess = 0

	// BikeAccessAllowed indicates that a vehicle can carry at least one bicycle
	BikeAccessAllowed = 1

	// BikeAccessForbidden indicates that bikes are not allowed
	BikeAccessForbidden = 2
)

// MarshalCSV converts the location type to a CSV string
func (b *BikeAccess) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*b)), nil
}

// UnmarshalCSV converts the CSV string to the location type representation
func (b *BikeAccess) UnmarshalCSV(csv string) error {
	d, err := strconv.Atoi(csv)
	n := BikeAccess(d)
	b = &n
	return err
}

// String converts the location type to a string
func (b *BikeAccess) String() string {
	return strconv.Itoa(int(*b))
}
