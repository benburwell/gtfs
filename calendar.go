package gtfs

// A Calendar is the dates for service IDs using a weekly schedule. Specify
// when service starts and ends, as well as days of the week where service is
// available.
type Calendar struct {
	// The service_id contains an ID that uniquely identifies a set of dates when
	// service is available for one or more routes. Each service_id value can
	// appear at most once in a calendar.txt file. This value is dataset unique.
	// It is referenced by the trips.txt file.
	ServiceID string `csv:"service_id"`

	// The monday field contains a binary value that indicates whether the
	// service is valid for all Mondays.
	//
	//     - A value of 1 indicates that service is available for all Mondays in
	//       the date range. (The date range is specified using the start_date
	//       and end_date fields.)
	//     - A value of 0 indicates that service is not available on Mondays in the
	//       date range.
	//
	// Note: You may list exceptions for particular dates, such as holidays, in
	// the calendar_dates.txt file.
	Monday bool `csv:"monday"`

	// The tuesday field contains a binary value that indicates whether the
	// service is valid for all Tuesdays.
	//
	//     - A value of 1 indicates that service is available for all Tuesdays in
	//       the date range. (The date range is specified using the start_date
	//       and end_date fields.)
	//     - A value of 0 indicates that service is not available on Tuesdays in
	//       the date range.
	//
	// Note: You may list exceptions for particular dates, such as holidays, in
	// the calendar_dates.txt file.
	Tuesday bool `csv:"tuesday"`

	// The wednesday field contains a binary value that indicates whether the
	// service is valid for all Wednesdays.
	//
	//     - A value of 1 indicates that service is available for all Wednesdays
	//       in the date range. (The date range is specified using the start_date
	//       and end_date fields.)
	//     - A value of 0 indicates that service is not available on Wednesdays
	//       in the date range.
	//
	// Note: You may list exceptions for particular dates, such as holidays, in
	// the calendar_dates.txt file.
	Wednesday bool `csv:"wednesday"`

	// The thursday field contains a binary value that indicates whether the
	// service is valid for all Thursdays.
	//
	//     - A value of 1 indicates that service is available for all Thursdays
	//       in the date range. (The date range is specified using the start_date
	//       and end_date fields.)
	//     - A value of 0 indicates that service is not available on Thursdays in
	//       the date range.
	//
	// Note: You may list exceptions for particular dates, such as holidays, in
	// the calendar_dates.txt file.
	Thursday bool `csv:"thursday"`

	// The friday field contains a binary value that indicates whether the
	// service is valid for all Fridays.
	//
	//     - A value of 1 indicates that service is available for all Fridays in
	//       the date range. (The date range is specified using the start_date
	//       and end_date fields.)
	//     - A value of 0 indicates that service is not available on Fridays in
	//       the date range.
	//
	// Note: You may list exceptions for particular dates, such as holidays, in
	// the calendar_dates.txt file.
	Friday bool `csv:"friday"`

	// The saturday field contains a binary value that indicates whether the
	// service is valid for all Saturdays.
	//
	//     - A value of 1 indicates that service is available for all Saturdays
	//       in the date range. (The date range is specified using the start_date
	//       and end_date fields.)
	//     - A value of 0 indicates that service is not available on Saturdays in
	//       the date range.
	//
	// Note: You may list exceptions for particular dates, such as holidays, in
	// the calendar_dates.txt file.
	Saturday bool `csv:"saturday"`

	// The sunday field contains a binary value that indicates whether the
	// service is valid for all Sundays.
	//
	//     - A value of 1 indicates that service is available for all Sundays in
	//       the date range. (The date range is specified using the start_date
	//       and end_date fields.)
	//     - A value of 0 indicates that service is not available on Sundays in
	//       the date range.
	//
	// Note: You may list exceptions for particular dates, such as holidays, in
	// the calendar_dates.txt file.
	Sunday bool `csv:"sunday"`

	// The start_date field contains the start date for the service. The
	// start_date field’s value should be in YYYYMMDD format.
	StartDate string `csv:"start_date"`

	// The end_date field contains the end date for the service. This date is
	// included in the service interval. The end_date field’s value should be in
	// YYYYMMDD format.
	EndDate string `csv:"end_date"`
}
