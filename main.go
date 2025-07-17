package main

import (
	"fmt"
	"log"
)

func main() {
	pickupDB := &PickupDB{}
	err := GetInfoClientDB("fake-tenant", "fake-token", pickupDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", pickupDB)
}
