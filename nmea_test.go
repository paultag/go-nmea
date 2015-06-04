package nmea_test

import (
	"log"
	"testing"

	"pault.ag/go/nmea"
)

func isok(t *testing.T, err error) {
	if err != nil {
		log.Printf("Error! Error is not nil! %s\n", err)
		t.FailNow()
	}
}

func notok(t *testing.T, err error) {
	if err == nil {
		log.Printf("Error! Error is nil!\n")
		t.FailNow()
	}
}

func assert(t *testing.T, expr bool) {
	if !expr {
		log.Printf("Assertion failed!")
		t.FailNow()
	}
}

func TestGPRMCSentence(t *testing.T) {
	rmc, err := nmea.NewGPRMCSentence(
		"$GPRMC,123519,A,4807.038,N,01131.000,E,022.4,084.4,230394,003.1,W*6A",
	)
	isok(t, err)
	assert(t, rmc != nil)
	assert(t, rmc.Time == 123519)
	assert(t, rmc.Date == 230394)

	assert(t, rmc.Speed == 022.4)
	assert(t, rmc.Track == 084.4)

	assert(t, rmc.Status == 'A')

	assert(t, rmc.Latitude == 4807.038)
	assert(t, rmc.Longitude == 01131.000)

	assert(t, rmc.Checksum == "6A")
}
