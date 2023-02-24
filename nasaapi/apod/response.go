package apod

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/goark/apod/nasaapi"
	"github.com/goark/errs"
)

// Response is response data from NASA APOD API.
type Response struct {
	Copyright      string       `json:"copyright,omitempty"`
	Date           nasaapi.Date `json:"date,omitempty"`
	Explanation    string       `json:"explanation,omitempty"`
	HdUrl          string       `json:"hdurl,omitempty"`
	MediaType      string       `json:"media_type,omitempty"`
	ServiceVersion string       `json:"service_version,omitempty"`
	Title          string       `json:"title,omitempty"`
	Url            string       `json:"url,omitempty"`
	ThumbnailUrl   string       `json:"thumbnail_url,omitempty"`
}

func decode(r io.Reader, isSingle bool) ([]Response, error) {
	var resps []Response
	dec := json.NewDecoder(r)
	if isSingle {
		for {
			var resp Response
			if err := dec.Decode(&resp); err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return nil, errs.Wrap(err)
			}
			resps = append(resps, resp)
		}
	} else {
		for {
			var resp []Response
			if err := dec.Decode(&resp); err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return nil, errs.Wrap(err)
			}
			resps = append(resps, resp...)
		}
	}
	return resps, nil
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
