package addresses

import "testing"

type scenario struct {
	address           string
	wantedAddressType string
}

func TestAddressType(t *testing.T) {
	scenarios := []scenario{
		{"Ocasion Avenue", "Avenue"},
		{"My Street", "Street"},
		{"Infinite Highway", "Highway"},
		{"Infinite HiGHway", "Highway"},
		{"Carlos Square", "Square"},
		{"Avenue Inverted", "Invalid type"},
		{"", "Invalid type"},
	}

	for _, scenario := range scenarios {
		gotAddressType := AddressType(scenario.address)

		if gotAddressType != scenario.wantedAddressType {
			t.Errorf("Wanted type \"%s\" got \"%s\"", scenario.wantedAddressType, gotAddressType)
		}
	}
}
