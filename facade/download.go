package facade

import (
	"context"

	"github.com/goark/apod/service/download"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newVersionCmd returns cobra.Command instance for show sub-command
func newDownload(ui *rwi.RWI) *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:     "download",
		Aliases: []string{"dl", "d"},
		Short:   "Download NASA APOD data",
		Long:    "Download NASA APOD data.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// global options
			dir := viper.GetString("base-dir")
			copyrightFlag := viper.GetBool("include-nopd")
			overwriteFlag := viper.GetBool("overwrite")
			cfg, err := makeAPODConfig()
			if err != nil {
				return debugPrint(ui, err)
			}

			// download APOD data
			if err := download.New(cfg, dir, copyrightFlag, overwriteFlag).Do(context.TODO()); err != nil {
				return debugPrint(ui, err)
			}
			return nil
		},
	}
	downloadCmd.Flags().StringP("base-dir", "d", "./apod", "Base directory for daownload")
	downloadCmd.Flags().BoolP("include-nopd", "", false, "Download no public domain images or videos")
	downloadCmd.Flags().BoolP("overwrite", "", false, "Overwrite Download files")

	//Bind config file
	_ = viper.BindPFlag("base-dir", downloadCmd.Flags().Lookup("base-dir"))
	_ = viper.BindPFlag("include-nopd", downloadCmd.Flags().Lookup("include-nopd"))
	_ = viper.BindPFlag("overwrite", downloadCmd.Flags().Lookup("overwrite"))

	return downloadCmd
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
