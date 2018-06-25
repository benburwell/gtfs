package gtfs

import "strconv"

// TransferType specifies the connection type
type TransferType int

// MarshalCSV marshals to CSV
func (t *TransferType) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*t)), nil
}

// UnmarshalCSV unmarshals from CSV
func (t *TransferType) UnmarshalCSV(csv string) error {
	d, err := strconv.Atoi(csv)
	n := TransferType(d)
	t = &n
	return err
}

// String returns a human-readable string
func (t *TransferType) String() string {
	return strconv.Itoa(int(*t))
}

const (
	// TransferTypeRecommended indicates a recommended transfer point between
	// routes.
	TransferTypeRecommended TransferType = 0

	// TransferTypeTimed indicates a timed transfer point between two routes. The
	// departing vehicle is expected to wait for the arriving one, with
	// sufficient time for a passenger to transfer between routes.
	TransferTypeTimed = 1

	// TransferTypeMinimum indicates this transfer requires a minimum amount of
	// time between arrival and departure to ensure a connection. The time
	// required to transfer is specified by min_transfer_time.
	TransferTypeMinimum = 2

	// TransferTypeNone indicates transfers are not possible between routes at
	// this location.
	TransferTypeNone = 3
)
