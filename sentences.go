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

	Time   float64 `nmea:"1"`
	Status string  `nmea:"2"`

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

// GPGSA {{{

type GPGSASentence struct {
	SentenceCore

	Fix struct {
		Selection string `nmea:"1"`
		Status    int    `nmea:"2"`
	}

	Satellites struct {
		PRN1  int `nmea:"3"`
		PRN2  int `nmea:"4"`
		PRN3  int `nmea:"5"`
		PRN4  int `nmea:"6"`
		PRN5  int `nmea:"7"`
		PRN6  int `nmea:"8"`
		PRN7  int `nmea:"9"`
		PRN8  int `nmea:"10"`
		PRN9  int `nmea:"11"`
		PRN10 int `nmea:"12"`
		PRN11 int `nmea:"13"`
		PRN12 int `nmea:"14"`
	}

	Dilution struct {
		Precision  float64 `nmea:"15"`
		Horizontal float64 `nmea:"16"`
		Vertical   float64 `nmea:"17"`
	}
}

func (g *GPGSASentence) GetSatellites() (ret []int) {
	for _, el := range []int{
		g.Satellites.PRN1,
		g.Satellites.PRN2,
		g.Satellites.PRN3,
		g.Satellites.PRN4,
		g.Satellites.PRN5,
		g.Satellites.PRN6,
		g.Satellites.PRN7,
		g.Satellites.PRN8,
		g.Satellites.PRN9,
		g.Satellites.PRN10,
		g.Satellites.PRN11,
		g.Satellites.PRN12,
	} {
		if el != 0 {
			ret = append(ret, el)
			continue
		}
		break
	}
	return
}

// }}}

// GPGGA {{{

type GPGGASentence struct {
	SentenceCore

	Time float64 `nmea:"1"`

	Latitude struct {
		Parallel   float64 `nmea:"2"`
		Hemisphere string  `nmea:"3"`
	}

	Longitude struct {
		Meridian   float64 `nmea:"4"`
		Hemisphere string  `nmea:"5"`
	}

	Quality            int     `nmea:"6"`
	SateliteCount      int     `nmea:"7"`
	HorizontalDilution float64 `nmea:"8"`

	Altitude struct {
		Value float64 `nmea:"9"`
		Unit  string  `nmea:"10"`
	}

	GeoidHeight struct {
		Value float64 `nmea:"11"`
		Unit  string  `nmea:"12"`
	}

	DGPS struct {
		SinceLastUpdate float64 `nmea:"13"`
		ID              int     `nmea:"14"`
	}
}

// }}}

// vim: foldmethod=marker
