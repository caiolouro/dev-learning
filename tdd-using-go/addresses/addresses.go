package addresses

import "strings"

func AddressType(address string) string {
	types := []string{"Street", "Avenue", "Highway", "Square"}

	words := strings.Split(address, " ")
	lastWord := words[len(words)-1]

	for _, addressType := range types {
		if strings.EqualFold(addressType, lastWord) {
			return addressType
		}
	}

	return "Invalid type"
}
