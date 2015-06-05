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
	"reflect"
)

type NMEA struct {
	GPRMC *GPRMCSentence `nmeaDataType:"GPRMC"`
}

func (d *NMEA) updateValue(key string, data reflect.Value) {
	dValue := reflect.ValueOf(d).Elem()
	dType := dValue.Type()

	for i := 0; i < dType.NumField(); i++ {
		fieldType := dType.Field(i)
		fieldValue := dValue.Field(i)

		if fieldType.Tag.Get("nmeaDataType") == key {
			fieldValue.Set(data.Addr())
			return
		}
	}
}

func (d *NMEA) Update(data Sentence) {
	dValue := reflect.ValueOf(data)
	if dValue.Kind() == reflect.Ptr {
		dValue = dValue.Elem()
	}
	d.updateValue(data.GetDataType(), dValue)
}

// vim: foldmethod=marker
