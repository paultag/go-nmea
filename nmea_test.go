/* {{{ Copyright (c) Paul R. Tagliamonte <paultag@gmail.com>, 2015
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE. }}} */

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
	rmc := nmea.GPRMCSentence{}
	err := nmea.Decode(&rmc,
		"$GPRMC,123519,A,4807.038,N,01131.000,E,022.4,084.4,230394,003.1,W*6A",
	)
	isok(t, err)

	assert(t, rmc.DataType == "GPRMC")

	assert(t, rmc.Time == 123519)
	assert(t, rmc.Date == 230394)

	assert(t, rmc.Speed == 022.4)
	assert(t, rmc.Track == 084.4)

	assert(t, rmc.Status == "A")

	assert(t, rmc.Latitude.Parallel == 4807.038)
	assert(t, rmc.Longitude.Meridian == 1131.000)
}

func TestBadChecksum(t *testing.T) {
	rmc := nmea.GPRMCSentence{}
	err := nmea.Decode(&rmc,
		"$GPRMC,123518,A,4807.038,N,01131.000,E,022.4,084.4,230394,003.1,W*6A",
	)
	notok(t, err)
}

func TestLatAndLong(t *testing.T) {
	rmc := nmea.GPRMCSentence{}
	err := nmea.Decode(&rmc,
		"$GPRMC,123519,A,4807.038,N,01131.000,E,022.4,084.4,230394,003.1,W*6A",
	)
	isok(t, err)

	assert(t, rmc.GetLatitude() == 48.07038)
	assert(t, rmc.GetLongitude() == 11.31000)
}

func TestGroundSpeed(t *testing.T) {
	rmc := nmea.GPRMCSentence{}
	err := nmea.Decode(&rmc,
		"$GPRMC,123519,A,4807.038,N,01131.000,E,022.4,084.4,230394,003.1,W*6A",
	)
	isok(t, err)
	assert(t, rmc.Speed == 22.4)
	assert(t, rmc.GetSpeedInKPH() == 41.4848)
}

func TestMagnetsHowDoTheyWork(t *testing.T) {
	rmc := nmea.GPRMCSentence{}
	err := nmea.Decode(&rmc,
		"$GPRMC,123519,A,4807.038,N,01131.000,E,022.4,084.4,230394,003.1,W*6A",
	)
	isok(t, err)
	assert(t, rmc.MagneticVariation.Value == 3.1)
	assert(t, rmc.MagneticVariation.Cardinal == "W")
}

// vim: foldmethod=marker
