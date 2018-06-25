package gtfs

import "strconv"

// ServiceType identifies whether bikes are permitted
type ServiceType int

const (
	// ServiceTypeScheduled indicates regularly scheduled service
	ServiceTypeScheduled ServiceType = 0

	// ServiceTypeNone indicates no service is available
	ServiceTypeNone = 1

	// ServiceTypePhone indicates one must phone agency to arrange service
	ServiceTypePhone = 2

	// ServiceTypeDriver indicates one must coordinate with driver to arrange
	// pickup
	ServiceTypeDriver = 3
)

// MarshalCSV converts the service  type to a CSV string
func (s *ServiceType) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*s)), nil
}

// UnmarshalCSV converts the CSV string to the service type representation
func (s *ServiceType) UnmarshalCSV(csv string) error {
	d, err := strconv.Atoi(csv)
	n := ServiceType(d)
	s = &n
	return err
}

// String converts the location type to a string
func (s *ServiceType) String() string {
	return strconv.Itoa(int(*s))
}
