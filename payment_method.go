package gtfs

import "strconv"

// PaymentMethod indicates when the fare must be paid
type PaymentMethod int

// MarshalCSV marshals to CSV
func (m *PaymentMethod) MarshalCSV() (string, error) {
	return strconv.Itoa(int(*m)), nil
}

// UnmarshalCSV unmarshals from CSV
func (m *PaymentMethod) UnmarshalCSV(csv string) error {
	d, err := strconv.Atoi(csv)
	n := PaymentMethod(d)
	m = &n
	return err
}

// String returns a human-readable string
func (m *PaymentMethod) String() string {
	return strconv.Itoa(int(*m))
}

const (
	// PaymentMethodOnBoard indicates fare is paid on board
	PaymentMethodOnBoard PaymentMethod = 0

	// PaymentMethodBeforeBoard indicates fare must be paid before boarding
	PaymentMethodBeforeBoard = 1
)
