package facade

import (
	"fmt"
	"runtime"

	"github.com/goark/errs"

	"github.com/goark/apod/ecode"
	"github.com/goark/apod/nasaapi"
	"github.com/goark/apod/nasaapi/apod"
	"github.com/goark/gocli/config"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	//Name is applicatin name
	Name = "apod"
	//Version is version for applicatin
	Version = "dev-version"
)

var (
	debugFlag         bool   //debug flag
	cfgFile           string //config file
	configFile        = "config"
	defaultConfigPath = config.Path(Name, configFile+".yaml")
)

// newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "OpenPGP packet visualizer",
		Long:  "OpenPGP (RFC 4880) packet visualizer by golang.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, errs.Wrap(ecode.ErrNoCommand))
		},
	}
	// global options (binding)
	rootCmd.PersistentFlags().StringP("api-key", "", "", "NASA API key")
	rootCmd.PersistentFlags().StringP("date", "", "", "date of the APOD image to retrieve (YYYY-MM-DD)")
	rootCmd.PersistentFlags().StringP("start-date", "", "", "start of a date range (YYYY-MM-DD)")
	rootCmd.PersistentFlags().StringP("end-date", "", "", "end of a date range (YYYY-MM-DD)")
	rootCmd.PersistentFlags().IntP("count", "", 0, "count randomly chosen images")
	rootCmd.PersistentFlags().BoolP("thumbs", "", false, "return the URL of video thumbnail")

	//Bind config file
	_ = viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
	_ = viper.BindPFlag("date", rootCmd.PersistentFlags().Lookup("date"))
	_ = viper.BindPFlag("start-date", rootCmd.PersistentFlags().Lookup("start-date"))
	_ = viper.BindPFlag("end-date", rootCmd.PersistentFlags().Lookup("end-date"))
	_ = viper.BindPFlag("count", rootCmd.PersistentFlags().Lookup("count"))
	_ = viper.BindPFlag("thumbs", rootCmd.PersistentFlags().Lookup("thumbs"))
	cobra.OnInitialize(initConfig)

	// global options (other)
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("Config file (default %v)", defaultConfigPath))

	rootCmd.SilenceUsage = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetArgs(args)
	rootCmd.SetIn(ui.Reader())       //Stdin
	rootCmd.SetOut(ui.ErrorWriter()) //Stdout -> Stderr
	rootCmd.SetErr(ui.ErrorWriter()) //Stderr
	rootCmd.AddCommand(
		newVersionCmd(ui),
		newLookup(ui),
	)

	return rootCmd
}

func makeAPODConfig() (*apod.Request, error) {
	date, err := nasaapi.DateFrom(viper.GetString("date"))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	startDate, err := nasaapi.DateFrom(viper.GetString("start-date"))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	endDate, err := nasaapi.DateFrom(viper.GetString("end-date"))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return apod.New(
		apod.WithDate(date),
		apod.WithStartDate(startDate),
		apod.WithEndDate(endDate),
		apod.WithCount(viper.GetInt("count")),
		apod.WithThumbs(viper.GetBool("thumbs")),
		apod.WithAPIKey(viper.GetString("api-key")),
	), nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find config directory.
		confDir := config.Dir(Name)
		if len(confDir) == 0 {
			confDir = "." //current directory
		}
		// Search config in home directory with name ".books-data.yaml" (without extension).
		viper.AddConfigPath(confDir)
		viper.SetConfigName(configFile)
	}
	viper.AutomaticEnv()     // read in environment variables that match
	_ = viper.ReadInConfig() // If a config file is found, read it in.
}

// Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
