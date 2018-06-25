package gtfs

// A StopTime is the times that a vehicle arrives at and departs from
// individual stops for each trip.
type StopTime struct {
	// The trip_id field contains an ID that identifies a trip. This value is
	// referenced from the trips.txt file.
	TripID string `csv:"trip_id"`

	// The arrival_time specifies the arrival time at a specific stop for a
	// specific trip on a route. The time is measured from “noon minus 12h”
	// (effectively midnight, except for days on which daylight savings time
	// changes occur) at the beginning of the service day. For times occurring
	// after midnight on the service day, enter the time as a value greater than
	// 24:00:00 in HH:MM:SS local time for the day on which the trip schedule
	// begins. If you don’t have separate times for arrival and departure at a
	// stop, enter the same value for arrival_time and departure_time.
	//
	// Scheduled stops where the vehicle strictly adheres to the specified
	// arrival and departure times are timepoints. For example, if a transit
	// vehicle arrives at a stop before the scheduled departure time, it will
	// hold until the departure time. If this stop is not a timepoint, use either
	// an empty string value for the arrival_time field or provide an
	// interpolated time. Further, indicate that interpolated times are provided
	// via the timepoint field with a value of zero. If interpolated times are
	// indicated with timepoint=0, then time points must be indicated with a
	// value of 1 for the timepoint field. Provide arrival times for all stops
	// that are time points.
	//
	// An arrival time must be specified for the first and the last stop in a
	// trip. Times must be eight digits in HH:MM:SS format (H:MM:SS is also
	// accepted, if the hour begins with 0). Do not pad times with spaces. The
	// following columns list stop times for a trip and the proper way to express
	// those times in the arrival_time field:
	//
	// | Time          | arrival_time value  |
	// |---------------+---------------------|
	// | 08:10:00 A.M. | 08:10:00 or 8:10:00 |
	// | 01:05:00 P.M. | 13:05:00            |
	// | 07:40:00 P.M. | 19:40:00            |
	// | 01:55:00 A.M. | 25:55:00            |
	//
	// Note: Trips that span multiple dates will have stop times greater than
	// 24:00:00. For example, if a trip begins at 10:30:00 p.m. and ends at
	// 2:15:00 a.m. on the following day, the stop times would be 22:30:00 and
	// 26:15:00. Entering those stop times as 22:30:00 and 02:15:00 would not
	// produce the desired results.
	ArrivalTime string `csv:"arrival_time"`

	// The departure_time specifies the departure time from a specific stop for a
	// specific trip on a route. The time is measured from “noon minus 12h”
	// (effectively midnight, except for days on which daylight savings time
	// changes occur) at the beginning of the service day. For times occurring
	// after midnight on the service day, enter the time as a value greater than
	// 24:00:00 in HH:MM:SS local time for the day on which the trip schedule
	// begins. If you don’t have separate times for arrival and departure at a
	// stop, enter the same value for arrival_time and departure_time.
	//
	// Scheduled stops where the vehicle strictly adheres to the specified
	// arrival and departure times are timepoints. For example, if a transit
	// vehicle arrives at a stop before the scheduled departure time, it will
	// hold until the departure time. If this stop is not a timepoint, use either
	// an empty string value for the departure_time field or provide an
	// interpolated time (further, indicate that interpolated times are provided
	// via the timepoint field with a value of zero). If interpolated times are
	// indicated with timepoint=0, then time points must be indicated with a
	// value of 1 for the timepoint field. Provide departure times for all stops
	// that are time points.
	//
	// A departure time must be specified for the first and the last stop in a
	// trip even if the vehicle does not allow boarding at the last stop. Times
	// must be eight digits in HH:MM:SS format (H:MM:SS is also accepted, if the
	// hour begins with 0). Do not pad times with spaces. The following columns
	// list stop times for a trip and the proper way to express those times in
	// the departure_time field:
	//
	// | Time          | departure_time value |
	// |---------------+----------------------|
	// | 08:10:00 A.M. | 08:10:00 or 8:10:00  |
	// | 01:05:00 P.M. | 13:05:00             |
	// | 07:40:00 P.M. | 19:40:00             |
	// | 01:55:00 A.M. | 25:55:00             |
	//
	// Note: Trips that span multiple dates will have stop times greater than
	// 24:00:00. For example, if a trip begins at 10:30:00 p.m. and ends at
	// 2:15:00 a.m. on the following day, the stop times would be 22:30:00 and
	// 26:15:00. Entering those stop times as 22:30:00 and 02:15:00 would not
	// produce the desired results.
	DepartureTime string `csv:"departure_time"`

	// The stop_id field contains an ID that uniquely identifies a stop. Multiple
	// routes may use the same stop. The stop_id is referenced from the stops.txt
	// file. If location_type is used in stops.txt, all stops referenced in
	// stop_times.txt must have location_type of 0. Where possible, stop_id
	// values should remain consistent between feed updates. In other words, stop
	// A with stop_id 1 should have stop_id 1 in all subsequent data updates. If
	// a stop is not a time point, enter blank values for arrival_time and
	// departure_time.
	StopID string `csv:"stop_id"`

	// The stop_sequence field identifies the order of the stops for a particular
	// trip. The values for stop_sequence must be non-negative integers, and they
	// must increase along the trip. For example, the first stop on the trip
	// could have a stop_sequence of 1, the second stop on the trip could have a
	// stop_sequence of 23, the third stop could have a stop_sequence of 40, and
	// so on.
	StopSequence int64 `csv:"stop_sequence"`

	// The stop_headsign field contains the text that appears on a sign that
	// identifies the trip’s destination to passengers. Use this field to
	// override the default trip_headsign when the headsign changes between
	// stops. If this headsign is associated with an entire trip, use
	// trip_headsign instead.
	StopHeadsign string `csv:"stop_headsign"`

	// The pickup_type field indicates whether passengers are picked up at a stop
	// as part of the normal schedule or whether a pickup at the stop is not
	// available. This field also allows the transit agency to indicate that
	// passengers must call the agency or notify the driver to arrange a pickup
	// at a particular stop.
	//
	// The default value for this field is ServiceTypeScheduled.
	PickupType *ServiceType `csv:"pickup_type"`

	// The drop_off_type field indicates whether passengers are dropped off at a
	// stop as part of the normal schedule or whether a drop off at the stop is
	// not available. This field also allows the transit agency to indicate that
	// passengers must call the agency or notify the driver to arrange a drop off
	// at a particular stop.
	//
	// The default value for this field is ServiceTypeScheduled.
	DropOffType *ServiceType `csv:"drop_off_type"`

	// When used in the stop_times.txt file, the shape_dist_traveled field
	// positions a stop as a distance from the first shape point. The
	// shape_dist_traveled field represents a real distance traveled along the
	// route in units such as feet or kilometers. For example, if a bus travels a
	// distance of 5.25 kilometers from the start of the shape to the stop, the
	// shape_dist_traveled for the stop ID would be entered as “5.25”. This
	// information allows the trip planner to determine how much of the shape to
	// draw when showing part of a trip on the map. The values used for
	// shape_dist_traveled must increase along with stop_sequence: they cannot be
	// used to show reverse travel along a route. The units used for
	// shape_dist_traveled in the stop_times.txt file must match the units that
	// are used for this field in the shapes.txt file.
	ShapeDistTraveled int64 `csv:"shape_dist_traveled"`

	// The timepoint field can be used to indicate if the specified arrival and
	// departure times for a stop are strictly adhered to by the transit vehicle
	// or if they are instead approximate and/or interpolated times. The field
	// allows a GTFS producer to provide interpolated stop times that potentially
	// incorporate local knowledge, but still indicate if the times are
	// approximate. For stop-time entries with specified arrival and departure
	// times, valid values for this field are:
	//
	//     empty - Times are considered exact.
	//     0 - Times are considered approximate.
	//     1 - Times are considered exact.
	//
	// For stop-time entries without specified arrival and departure times, feed
	// consumers must interpolate arrival and departure times. Feed producers may
	// optionally indicate that such an entry is not a timepoint (value=0) but it
	// is an error to mark a entry as a timepoint (value=1) without specifying
	// arrival and departure times.
	Timepoint *Timepoint `csv:"timepoint"`
}
