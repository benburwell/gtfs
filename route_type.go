package gtfs

import "strconv"

// RouteType is the type of transportation for a route
type RouteType int

// MarshalCSV marshals to CSV
func (r *RouteType) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*r)), nil
}

// UnmarshalCSV unmarshals from CSV
func (r *RouteType) UnmarshalCSV(csv string) error {
	d, err := strconv.Atoi(csv)
	n := RouteType(d)
	r = &n
	return err
}

// String returns a human-readable string
func (r *RouteType) String() string {
	return strconv.Itoa(int(*r))
}

const (
	// RouteTypeTram indicates any light rail or street level system within a
	// metropolitan area.
	RouteTypeTram RouteType = 0

	// RouteTypeMetro indicates any underground rail system within a
	// metropolitan area.
	RouteTypeMetro = 1

	// RouteTypeRail is used for intercity or long-distance travel.
	RouteTypeRail = 2

	// RouteTypeBus is used for short- and long-distance bus routes.
	RouteTypeBus = 3

	// RouteTypeFerry is used for short- and long-distance boat service.
	RouteTypeFerry = 4

	// RouteTypeCableCar is used for street-level cable cars where the cable runs
	// beneath the car.
	RouteTypeCableCar = 5

	// RouteTypeGondola is typically used for aerial cable cars where the car is
	// suspended from the cable.
	RouteTypeGondola = 6

	// RouteTypeFunicular is any rail system designed for steep inclines.
	RouteTypeFunicular = 7
)
