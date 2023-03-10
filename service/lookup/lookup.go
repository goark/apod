package lookup

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/goark/apod/ecode"
	"github.com/goark/apod/nasaapi/apod"
	"github.com/goark/errs"
)

// Lookup is configuration for lookup command.
type Lookup struct {
	*apod.Request
	rawFlag bool
}

// New returns new Lookup instance.
func New(cfg *apod.Request, rawFlag bool) *Lookup {
	return &Lookup{Request: cfg, rawFlag: rawFlag}
}

// Do method is looking up APOD data from NASA API.
func (l *Lookup) Do(ctx context.Context) (io.ReadCloser, error) {
	if l == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	if l.rawFlag {
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
