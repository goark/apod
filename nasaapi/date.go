package nasaapi

import (
	"strconv"
	"strings"
	"time"

	"github.com/goark/errs"
)

// Date is wrapper class of time.Time.
type Date struct {
	time.Time
}

// NewDate returns Date instance from time.Time.
func NewDate(tm time.Time) Date {
	return Date{tm}
}

func (t Date) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.DateOnly)
}

// MarshalJSON implements the json.Marshaler interface.
func (t Date) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(strconv.Quote(t.String())), nil
}

var timeTemplate = []string{
	time.RFC3339,
	time.DateOnly,
}

// DateFrom returns Date instance from date string
func DateFrom(s string) (Date, error) {
	if len(s) == 0 || strings.EqualFold(s, "null") {
		return NewDate(time.Time{}), nil
	}
	var lastErr error
	for _, tmplt := range timeTemplate {
		if tm, err := time.Parse(tmplt, s); err != nil {
			lastErr = errs.Wrap(err, errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt))
		} else {
			return NewDate(tm), nil
		}
	}
	return NewDate(time.Time{}), lastErr
}

// UnmarshalJSON implements the json.UnmarshalJSON interface.
func (t *Date) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*t, err = DateFrom(s)
	return err
}

/* MIT License
 *
 * Copyright 2023 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
