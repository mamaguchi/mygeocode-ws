package main

import (
	"os"
	"context"
	"fmt"
	"log"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func main() {
	GMAP_API_KEY := os.Getenv("GMAP_API_KEY")
	if GMAP_API_KEY == "" {
		log.Fatalf("Google maps api key not found!")
	}
	client, err := maps.NewClient(maps.WithAPIKey(GMAP_API_KEY))
	check(err)

	ADDRESS := "aksdj, dke,lakeeje"
	r := &maps.GeocodingRequest{
		Address:  ADDRESS,
		Language: "en",
		Region:   "my",
	}

	resp, err := client.Geocode(context.Background(), r)
	// 'err' will not be nil if Geocode status != "OK" && != "ZERO_RESULTS".
	check(err) 

	// Now we check for "ZERO_RESULTS" Geocode status.
	if len(resp) == 0 {
		log.Printf("GeocodingResult: No results\n")
		os.Exit(0)
	}

	// Here we print the Geocode result.
	pretty.Println(resp)
	fmt.Println(resp[0].FormattedAddress)
	fmt.Printf("Longitude: %v\n", resp[0].Geometry.Location.Lng)
	fmt.Printf("Latitude: %v\n", resp[0].Geometry.Location.Lat)
}