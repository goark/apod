package lookup

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/goark/apod/nasaapi"
	"github.com/goark/apod/nasaapi/apod"
	"github.com/goark/errs"
)

// Lookup is configuration for lookup command.
type Lookup struct {
	*apod.Context
}

// New returns new Lookup instance.
func New(apiKey, dateStr, startDateStr, endDateStr string, count int, thumbs bool) (*Lookup, error) {
	date, err := nasaapi.DateFrom(dateStr)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("date", dateStr))
	}
	startDate, err := nasaapi.DateFrom(dateStr)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("start_date", startDateStr))
	}
	endDate, err := nasaapi.DateFrom(endDateStr)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("end_date", endDateStr))
	}
	return &Lookup{apod.New(
		apod.WithAPIKey(apiKey),
		apod.WithDate(date),
		apod.WithStartDate(startDate),
		apod.WithEndDate(endDate),
		apod.WithCount(count),
		apod.WithThumbs(thumbs),
	)}, nil
}

// Do method is looking up APOD data from NASA API.
func (l *Lookup) Do(ctx context.Context, rawFlag bool) (io.ReadCloser, error) {
	if rawFlag {
		return l.GetRawData(ctx)
	}
	resp, err := l.Get(ctx)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return &readClose{bytes.NewReader(b)}, nil
}

type readClose struct {
	io.Reader
}

func (rc *readClose) Close() error {
	if c, ok := rc.Reader.(io.Closer); ok {
		return c.Close()
	}
	return nil
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
