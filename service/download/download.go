package download

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/goark/apod/ecode"
	"github.com/goark/apod/nasaapi/apod"
	"github.com/goark/errs"
	"github.com/goark/fetch"
)

const maxDataSize = 1024 * 1024 * 1024 //1GB

// Download is configuration for download command.
type Download struct {
	*apod.Request
	baseDir       string
	copyrightFlag bool
	overwriteFlag bool
}

// New returns new Lookup instance.
func New(cfg *apod.Request, baseDir string, copyrightFlag, overwriteFlag bool) *Download {
	if len(baseDir) == 0 {
		baseDir = "."
	}
	return &Download{
		Request:       cfg,
		baseDir:       baseDir,
		copyrightFlag: copyrightFlag,
		overwriteFlag: overwriteFlag,
	}
}

// Do method is downloading APOD data from NASA API.
func (dl *Download) Do(ctx context.Context) error {
	if dl == nil {
		return errs.Wrap(ecode.ErrNullPointer)
	}
	// get APOD data from NASA API
	resps, err := dl.Get(ctx)
	if err != nil {
		return errs.Wrap(err)
	}

	// make directory
	if _, err := os.Stat(dl.baseDir); err != nil { // dirextory is not found
		if err := os.MkdirAll(dl.baseDir, os.ModePerm); err != nil {
			return errs.Wrap(err)
		}
	}

	for _, resp := range resps {
		// make directory
		dir := filepath.Join(dl.baseDir, resp.Date.String())
		if _, err := os.Stat(dir); err != nil { // dirextory is not found
			if err := os.Mkdir(dir, os.ModePerm); err != nil {
				return errs.Wrap(err, errs.WithContext("dir", dir))
			}
		} else if !dl.overwriteFlag {
			continue
		} else {
			if err := os.RemoveAll(dir); err != nil {
				return errs.Wrap(err, errs.WithContext("dir", dir))
			}
			if err := os.Mkdir(dir, 0755); err != nil {
				return errs.Wrap(err, errs.WithContext("dir", dir))
			}
		}

		// output metadata.json file
		if err := saveMetadate(resp, filepath.Join(dir, "metadata.json")); err != nil {
			return errs.Wrap(err)
		}
		// download image/video files
		if len(resp.Copyright) > 0 && !dl.copyrightFlag {
			continue
		}
		if len(resp.HdUrl) > 0 {
			if err := downloadImage(ctx, resp.HdUrl, dir); err != nil {
				return errs.Wrap(err, errs.WithContext("hdUrl", resp.HdUrl))
			}
		}
		if len(resp.Url) > 0 {
			if err := downloadImage(ctx, resp.Url, dir); err != nil {
				return errs.Wrap(err, errs.WithContext("url", resp.Url))
			}
		}
		if len(resp.ThumbnailUrl) > 0 {
			if err := downloadImage(ctx, resp.ThumbnailUrl, dir); err != nil {
				return errs.Wrap(err, errs.WithContext("thumbnailUrl", resp.ThumbnailUrl))
			}
		}

	}

	return nil
}

func saveMetadate(resp *apod.Response, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("path", path))
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	return errs.Wrap(enc.Encode(resp))
}

func downloadImage(ctx context.Context, urlStr string, dir string) error {
	u, err := fetch.URL(urlStr)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("url", urlStr))
	}
	_, fname := path.Split(u.Path)
	resp, err := fetch.New().Get(u, fetch.WithContext(ctx))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("url", urlStr))
	}
	defer resp.Close()

	path := filepath.Join(dir, fname)
	file, err := os.Create(path)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("path", path))
	}
	defer file.Close()
	if _, err := io.CopyN(file, resp.Body(), maxDataSize); err != nil {
		if !errors.Is(err, io.EOF) {
			return errs.Wrap(err, errs.WithContext("path", path))
		}
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
