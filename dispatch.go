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

import (
	"fmt"
	"reflect"
	"strings"
)

type NMEA struct {
	parsers map[string](func(string) (Sentence, error))
	RMC     *GPRMCSentence `nmeaDataType:"GPRMC"`
	GSA     *GPGSASentence `nmeaDataType:"GPGSA"`
	GGA     *GPGGASentence `nmeaDataType:"GPGGA"`
}

func NewNMEA() NMEA {
	n := NMEA{parsers: map[string](func(string) (Sentence, error)){}}

	n.Register("GPRMC", func(data string) (Sentence, error) {
		rmc := GPRMCSentence{}
		err := Decode(&rmc, data)
		return &rmc, err
	})

	return n
}

func (d *NMEA) updateValue(key string, data reflect.Value) error {
	dValue := reflect.ValueOf(d).Elem()
	dType := dValue.Type()

	for i := 0; i < dType.NumField(); i++ {
		fieldType := dType.Field(i)
		fieldValue := dValue.Field(i)

		if fieldType.Tag.Get("nmeaDataType") == key {
			fieldValue.Set(data.Addr())
			return nil
		}
	}
	return fmt.Errorf("No such tag found: %s", key)
}

func (d *NMEA) Register(key string, handler func(string) (Sentence, error)) {
	d.parsers[key] = handler
}

func (d *NMEA) Parse(data string) error {
	els := strings.SplitN(data, ",", 2)
	if handler, ok := d.parsers[els[0][1:]]; ok {
		el, err := handler(data)
		if err != nil {
			return err
		}
		d.Update(el)
	}
	return nil
}

func (d *NMEA) Update(data Sentence) error {
	dValue := reflect.ValueOf(data)
	if dValue.Kind() == reflect.Ptr {
		dValue = dValue.Elem()
	}
	return d.updateValue(data.GetDataType(), dValue)
}

// vim: foldmethod=marker
