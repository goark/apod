package nasaapi

import (
	"context"
	"io"
	"net/url"

	"github.com/goark/errs"
	"github.com/goark/fetch"
)

const DefaultAPIKey = "DEMO_KEY" // Default NASA API key (for demo)

const (
	scheme = "https"
	host   = "api.nasa.gov"
)

// Request function requests to NASA API, and returns response data.
func Request(ctx context.Context, path string, q url.Values) (io.ReadCloser, error) {
	resp, err := fetch.New().Get(getURL(path, q), fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return resp.Body(), nil
}

func getURL(path string, q url.Values) *url.URL {
	return &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     path,
		RawQuery: q.Encode(),
	}
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
