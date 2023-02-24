package facade

import (
	"context"

	"github.com/goark/apod/service/lookup"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

// newVersionCmd returns cobra.Command instance for show sub-command
func newLookup(ui *rwi.RWI) *cobra.Command {
	lookupCmd := &cobra.Command{
		Use:     "lookup",
		Aliases: []string{"look", "l"},
		Short:   "Look up NASA APOD data",
		Long:    "Look up NASA APOD data.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// global options
			cfg, err := makeAPODConfig()
			if err != nil {
				return debugPrint(ui, err)
			}
			// local options
			rawFlag, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return debugPrint(ui, err)
			}

			// lookup APOD data
			r, err := lookup.New(cfg).Do(context.TODO(), rawFlag)
			if err != nil {
				return debugPrint(ui, err)
			}
			defer r.Close()
			return debugPrint(ui, ui.WriteFrom(r))
		},
	}
	lookupCmd.Flags().BoolP("raw", "", false, "Output raw data from APOD API")

	return lookupCmd
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
