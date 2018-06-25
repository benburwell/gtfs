package gtfs

// A Transfer is a rule for making  making connections at transfer points
// between routes.
//
// Trip planners normally calculate transfer points based on the relative
// proximity of stops in each route. For potentially ambiguous stop pairs, or
// transfers where you want to specify a particular choice, use transfers.txt
// to define additional rules for making connections between routes.
type Transfer struct {
	// The from_stop_id field contains a stop ID that identifies a stop or
	// station where a connection between routes begins. Stop IDs are referenced
	// from the stops.txt file. If the stop ID refers to a station that contains
	// multiple stops, this transfer rule applies to all stops in that station.
	FromStopID string `csv:"from_stop_id"`

	// The to_stop_id field contains a stop ID that identifies a stop or station
	// where a connection between routes ends. Stop IDs are referenced from the
	// stops.txt file. If the stop ID refers to a station that contains multiple
	// stops, this transfer rule applies to all stops in that station.
	ToStopID string `csv:"to_stop_id"`

	// The transfer_type field specifies the type of connection for the specified
	// (from_stop_id, to_stop_id) pair.
	TransferType *TransferType `csv:"transfer_type"`

	// When a connection between routes requires an amount of time between
	// arrival and departure (transfer_type=2), the min_transfer_time field
	// defines the amount of time that must be available in an itinerary to
	// permit a transfer between routes at these stops. The min_transfer_time
	// must be sufficient to permit a typical rider to move between the two
	// stops, including buffer time to allow for schedule variance on each route.
	//
	// The min_transfer_time value must be entered in seconds, and must be a
	// non-negative integer.
	MinTransferTime *int64 `csv:"min_transfer_time"`
}
