package nmea

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Sentence struct {
	DataType string
}

type GPRMCSentence struct {
	Sentence

	Time              int
	Status            rune
	Latitude          float64
	Longitude         float64
	Speed             float64
	Track             float64
	Date              int
	MagneticVariation string
	Checksum          string
}

func NewGPRMCSentence(data string) (*GPRMCSentence, error) {
	ret := GPRMCSentence{}
	ret.DataType = "GPRMC"

	elements := strings.Split(data, ",")
	if len(elements) != 12 {
		return nil, errors.New("Malformed GPRMC Sentence")
	}

	/* Now, we need to clean up the checksum */
	lastEl := elements[11]
	els := strings.SplitN(lastEl, "*", 2)
	if len(els) != 2 {
		return nil, errors.New("Can't find the checksum!")
	}

	elements[11] = els[0]
	ret.Checksum = els[1]

	ret.Status = rune(elements[2][0])

	for target, index := range map[*int]int{
		&ret.Time: 1,
		&ret.Date: 9,
	} {
		when, err := strconv.Atoi(elements[index])
		if err != nil {
			return nil, fmt.Errorf("Malformed time: %s", err)
		}
		*target = when
	}

	for target, index := range map[*float64]int{
		&ret.Latitude:  3,
		&ret.Longitude: 5,
		&ret.Speed:     7,
		&ret.Track:     8,
	} {
		when, err := strconv.ParseFloat(elements[index], 64)
		if err != nil {
			return nil, fmt.Errorf("Malformed time: %s", err)
		}
		*target = when
	}

	/* N positive, S, neg */
	if elements[4] == "S" {
		ret.Latitude = -ret.Latitude
	}

	/* E positive, W, neg */
	if elements[6] == "W" {
		ret.Latitude = -ret.Latitude
	}

	return &ret, nil
}
