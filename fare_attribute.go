package gtfs

// A FareAttribute holds fare information for a transit organization’s routes.
type FareAttribute struct {
	// The fare_id field contains an ID that uniquely identifies a fare class.
	// The fare_id is dataset unique.
	FareID string `csv:"fare_id"`

	// The price field contains the fare price, in the unit specified by
	// currency_type.
	Price string `csv:"price"`

	// The currency_type field defines the currency used to pay the fare. Please
	// use the ISO 4217 alphabetical currency codes which can be found here.
	CurrencyType string `csv:"currency_type"`

	// The payment_method field indicates when the fare must be paid. Valid
	// values for this field are:
	//
	//     0 - Fare is paid on board.
	//     1 - Fare must be paid before boarding.
	PaymentMethod *PaymentMethod `csv:"payment_method"`

	// The transfers field specifies the number of transfers permitted on this
	// fare. Valid values for this field are:
	//
	//     0 - No transfers permitted on this fare.
	//     1 - Passenger may transfer once.
	//     2 - Passenger may transfer twice.
	//     (empty) - If this field is empty, unlimited transfers are permit.
	Transfers *int64 `csv:"transfers"`

	// Required for feeds with multiple agencies defined in the agency.txt file.
	// Each fare attribute must specify an agency_id value to indicate which
	// agency the fare applies to.
	AgencyID string `csv:"agency_id"`

	// The transfer_duration field specifies the length of time in seconds before
	// a transfer expires. When used with a transfers value of 0, the
	// transfer_duration field indicates how long a ticket is valid for a fare
	// where no transfers are allowed. Unless you intend to use this field to
	// indicate ticket validity, transfer_duration should be omitted or empty
	// when transfers is set to 0.
	TransferDuration *int64 `csv:"transfer_duration"`
}
