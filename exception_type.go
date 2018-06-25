package gtfs

import "strconv"

// ExceptionType indicates whether a schedule exception is an addition or a removal
type ExceptionType int

// MarshalCSV marshals to CSV
func (e *ExceptionType) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*e)), nil
}

// UnmarshalCSV unmarshals from CSV
func (e *ExceptionType) UnmarshalCSV(csv string) error {
	d, err := strconv.Atoi(csv)
	n := ExceptionType(d)
	e = &n
	return err
}

// String returns a human-readable string
func (e *ExceptionType) String() string {
	return strconv.Itoa(int(*e))
}

const (
	// ExceptionTypeAddition indicates that service has been added
	ExceptionTypeAddition ExceptionType = 1

	// ExceptionTypeRemoval indicates that service has been removed
	ExceptionTypeRemoval = 2
)
