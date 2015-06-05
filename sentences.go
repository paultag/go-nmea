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

package nmea

type Sentence interface {
	GetDataType() string
}

type SentenceCore struct {
	Sentence
	DataType string `nmea:"0"`
}

func (s *SentenceCore) GetDataType() string {
	return s.DataType
}

// GPRMC {{{

type GPRMCSentence struct {
	SentenceCore

	Time   int    `nmea:"1"`
	Status string `nmea:"2"`

	Latitude struct {
		Parallel   float64 `nmea:"3"`
		Hemisphere string  `nmea:"4"`
	}

	Longitude struct {
		Meridian   float64 `nmea:"5"`
		Hemisphere string  `nmea:"6"`
	}

	Speed float64 `nmea:"7"`
	Track float64 `nmea:"8"`
	Date  int     `nmea:"9"`

	MagneticVariation struct {
		Value    float64 `nmea:"10"`
		Cardinal string  `nmea:"11"`
	}
}

/* 	            m/s      km/h      mph      knot      ft/s
 * 		   +----------+--------+----------+-------+---------
 * 	1 knot | 0.514444 | 1.852  | 1.150779 |   1   | 1.687810 */

func (rmc *GPRMCSentence) GetSpeedInKPH() float64 {
	return rmc.Speed * 1.852
}

var nmeaToDegreeMultiplier = 0.01

func (rmc *GPRMCSentence) GetLatitude() float64 {
	switch rmc.Latitude.Hemisphere {
	case "N":
		return (rmc.Latitude.Parallel * nmeaToDegreeMultiplier)
	case "S":
		return -(rmc.Latitude.Parallel * nmeaToDegreeMultiplier)
	default:
		return 0
	}
}

func (rmc *GPRMCSentence) GetLongitude() float64 {
	switch rmc.Longitude.Hemisphere {
	case "E":
		return (rmc.Longitude.Meridian * nmeaToDegreeMultiplier)
	case "W":
		return -(rmc.Longitude.Meridian * nmeaToDegreeMultiplier)
	default:
		return 0
	}
}

// }}}

// vim: foldmethod=marker
