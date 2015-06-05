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
	"strconv"
	"strings"
)

func decodeToValue(incoming reflect.Value, data string) error {
	switch incoming.Type().Kind() {
	case reflect.String:
		incoming.SetString(data)
	case reflect.Int:
		if data == "" {
			incoming.SetInt(0)
			return nil
		}
		value, err := strconv.Atoi(data)
		if err != nil {
			return err
		}
		incoming.SetInt(int64(value))
	case reflect.Float64:
		if data == "" {
			incoming.SetFloat(0.0)
			return nil
		}
		value, err := strconv.ParseFloat(data, 64)
		if err != nil {
			return err
		}
		incoming.SetFloat(value)
	}
	return nil
}

func decodeToPointer(incoming reflect.Value, data []string) error {
	itype := incoming.Type()
	if itype.Kind() == reflect.Ptr {
		/* deref */
		return decodeToPointer(incoming.Elem(), data)
	}

	for i := 0; i < incoming.NumField(); i++ {
		field := incoming.Field(i)
		fieldType := itype.Field(i)

		if field.Type().Kind() == reflect.Struct {
			err := decodeToPointer(field, data)
			if err != nil {
				return err
			}
		}

		if it := fieldType.Tag.Get("nmea"); it != "" {
			index, err := strconv.Atoi(it)
			if err != nil {
				return fmt.Errorf("nmea tag failed to int-ize: %s", err)
			}
			err = decodeToValue(field, data[index])
			if err != nil {
				return fmt.Errorf("failed to set %s: %s", fieldType.Name, err)
			}
		}

	}

	return nil
}

func Decode(incoming Sentence, data string) error {
	incomingData := strings.SplitN(data, "*", 2)
	if len(incomingData) != 2 {
		return fmt.Errorf("Was expecting a checksum")
	}

	checksum, err := strconv.ParseUint(incomingData[1], 16, 8)
	if err != nil {
		return err
	}

	var computedChecksum = uint8(0)

	for i := 1; i < len(data); i++ {
		if data[i] == '*' {
			break
		}
		computedChecksum = computedChecksum ^ data[i]
	}

	if computedChecksum != uint8(checksum) {
		return fmt.Errorf("Bad checksum: %d vs %d", checksum, computedChecksum)
	}

	nmeaData := strings.Split(incomingData[0][1:], ",")
	return decodeToPointer(reflect.ValueOf(incoming), nmeaData)
}

// vim: foldmethod=marker
