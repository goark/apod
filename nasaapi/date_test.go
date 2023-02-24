package nasaapi

import (
	"encoding/json"
	"testing"
)

func TestDate(t *testing.T) {
	testCases := []struct {
		s     string
		isErr bool
	}{
		{s: "2023-02-22", isErr: false},
		{s: "2023", isErr: true},
		{s: "", isErr: false},
	}

	for _, tc := range testCases {
		s, err := DateFrom(tc.s)
		if (err != nil) != tc.isErr {
			t.Errorf("Is \"%v\" error ? %v, want %v", tc.s, err != nil, tc.isErr)
		}
		if err == nil {
			if s.String() != tc.s {
				t.Errorf("DateFrom(\"%v\") is \"%v\" , want \"%v\"", tc.s, s, tc.s)
			}
		}
	}
}

func TestDateJSON(t *testing.T) {
	testCases := []struct {
		s string
	}{
		{s: `{"date":"2023-02-22"}`},
		{s: `{"date":""}`},
	}

	for _, tc := range testCases {
		var data struct {
			Dt Date `json:"date"`
		}
		if err := json.Unmarshal([]byte(tc.s), &data); err != nil {
			t.Errorf("Unmarshal(\"%v\") is %v, want nil", tc.s, err)
		} else if b, err := json.Marshal(data); err != nil {
			t.Errorf("Marshal(\"%v\") is %v, want nil", tc.s, err)
		} else if string(b) != tc.s {
			t.Errorf("Unmarshal/Marshal is \"%v\", want \"%v\"", string(b), tc.s)
		}
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
