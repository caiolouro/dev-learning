package main

import (
	"fmt"
	"tdd-using-go/addresses"
)

func main() {
	addressType := addresses.AddressType("Ocasion Avenue")
	fmt.Println(addressType)
}
