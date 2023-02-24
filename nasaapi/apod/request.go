package apod

import (
	"context"
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/goark/apod/ecode"
	"github.com/goark/apod/nasaapi"
	"github.com/goark/errs"
)

const APIPath = "/planetary/apod"

// Request is for context of APOD API.
type Request struct {
	Date      nasaapi.Date `json:"date,omitempty"`       // The date of the APOD image to retrieve
	StartDate nasaapi.Date `json:"start_date,omitempty"` // The start of a date range, when requesting date for a range of dates. Cannot be used with date.
	EndDate   nasaapi.Date `json:"end_date,omitempty"`   // The end of the date range, when used with start_date.
	Count     int          `json:"count,omitempty"`      // If this is specified then count randomly chosen images will be returned. Cannot be used with date or start_date and end_date.
	Thumbs    bool         `json:"thumbs,omitempty"`     // Return the URL of video thumbnail. If an APOD is not a video, this parameter is ignored.
	APIKey    string       `json:"api_key"`              // api.nasa.gov key for expanded usage
}

type Opts func(*Request)

// New returns new Request instance for APOD API.
func New(opts ...Opts) *Request {
	ctx := &Request{}
	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}

// WithDate returns function for setting Request.Date.
func WithDate(date nasaapi.Date) Opts {
	return func(ctx *Request) {
		if ctx != nil {
			ctx.Date = date
		}
	}
}

// WithStartDate returns function for setting Request.StartDate.
func WithStartDate(startDate nasaapi.Date) Opts {
	return func(ctx *Request) {
		if ctx != nil {
			ctx.StartDate = startDate
		}
	}
}

// WithEndDate returns function for setting Request.EndDate.
func WithEndDate(endDate nasaapi.Date) Opts {
	return func(ctx *Request) {
		if ctx != nil {
			ctx.EndDate = endDate
		}
	}
}

// WithCount returns function for setting Request.Count.
func WithCount(count int) Opts {
	return func(ctx *Request) {
		if ctx != nil {
			ctx.Count = count
		}
	}
}

// WithThumbs returns function for setting Request.Thumbs.
func WithThumbs(thumbs bool) Opts {
	return func(ctx *Request) {
		if ctx != nil {
			ctx.Thumbs = thumbs
		}
	}
}

// WithAPIKey returns function for setting Request.APIKey.
func WithAPIKey(apiKey string) Opts {
	return func(ctx *Request) {
		if ctx != nil {
			ctx.APIKey = apiKey
		}
	}
}

// Encode returns JSON string.
func (apod *Request) Encode() (string, error) {
	if apod == nil {
		return "", errs.Wrap(ecode.ErrNullPointer)
	}
	b, err := json.Marshal(apod)
	if err != nil {
		return "", errs.Wrap(err)
	}
	return string(b), err
}

// Stringger method.
func (apod *Request) String() string {
	s, err := apod.Encode()
	if err != nil {
		return ""
	}
	return s
}

// Get method gets APOD data from NASA API, and returns []*Response instance.
func (apod *Request) Get(ctx context.Context) ([]*Response, error) {
	if apod == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	resp, err := apod.GetRawData(ctx)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	return decode(resp, apod.isSingle())
}

// GetRawData method gets APOD data from NASA API, and returns raw response string.
func (apod *Request) GetRawData(ctx context.Context) (io.ReadCloser, error) {
	if apod == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	q, err := apod.makeQuery()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return nasaapi.Request(ctx, APIPath, q)
}

func (apod *Request) isSingle() bool {
	if !apod.Date.IsZero() {
		return true
	}
	if apod.StartDate.IsZero() && apod.EndDate.IsZero() && apod.Count == 0 {
		return true
	}
	return false
}

func (apod *Request) makeQuery() (url.Values, error) {
	v := url.Values{}
	if !apod.Date.IsZero() {
		if !apod.StartDate.IsZero() || !apod.EndDate.IsZero() || apod.Count > 0 {
			return nil, errs.Wrap(ecode.ErrCombination, errs.WithContext("config", apod))
		}
		v.Set("date", apod.Date.Format(time.DateOnly))
	}
	if !apod.StartDate.IsZero() {
		if !apod.Date.IsZero() || apod.Count > 0 {
			return nil, errs.Wrap(ecode.ErrCombination, errs.WithContext("config", apod))
		}
		v.Set("start_date", apod.StartDate.Format(time.DateOnly))
	}
	if !apod.EndDate.IsZero() {
		if apod.StartDate.IsZero() || !apod.Date.IsZero() || apod.Count > 0 {
			return nil, errs.Wrap(ecode.ErrCombination, errs.WithContext("config", apod))
		}
		v.Set("end_date", apod.EndDate.Format(time.DateOnly))
	}
	if apod.Count > 0 {
		v.Set("count", strconv.Itoa(apod.Count))
	}
	if apod.Thumbs {
		v.Set("thumbs", "true")
	}
	if len(apod.APIKey) > 0 {
		v.Set("api_key", apod.APIKey)
	} else {
		v.Set("api_key", nasaapi.DefaultAPIKey)
	}
	return v, nil
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
