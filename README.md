# [apod] - CLI Tool for Astronomy Picture of the Day with NASA API

[![check vulns](https://github.com/goark/apod/workflows/vulns/badge.svg)](https://github.com/goark/apod/actions)
[![lint status](https://github.com/goark/apod/workflows/lint/badge.svg)](https://github.com/goark/apod/actions)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/goark/apod/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/apod.svg)](https://github.com/goark/apod/releases/latest)

This package is required Go 1.20 or later.

## Build and Install

```
$ go install github.com/goark/apod@latest
```

### Binaries

See [latest release](https://github.com/goark/apod/releases/latest).

## Usage

```
$ apod -h
OpenPGP (RFC 4880) packet visualizer by golang.

Usage:
  apod [flags]
  apod [command]

Available Commands:
  download    Download NASA APOD data
  help        Help about any command
  lookup      Look up NASA APOD data
  version     Print the version number

Flags:
      --api-key string      NASA API key
      --config string       Config file (default /home/spiegel/.config/apod/config.yaml)
      --count int           count randomly chosen images
      --date string         date of the APOD image to retrieve (YYYY-MM-DD)
      --debug               for debug
      --end-date string     end of a date range (YYYY-MM-DD)
  -h, --help                help for apod
      --start-date string   start of a date range (YYYY-MM-DD)
      --thumbs              return the URL of video thumbnail

Use "apod [command] --help" for more information about a command.
```

### Config file

```yaml:config.yaml
api-key: your_api_key_string
```

### Lookup APOD data

```
$ apod lookup -h
Look up NASA APOD data.

Usage:
  apod lookup [flags]

Aliases:
  lookup, look, l

Flags:
  -h, --help   help for lookup
      --raw    Output raw data from APOD API

Global Flags:
      --api-key string      NASA API key
      --config string       Config file (default /home/spiegel/.config/apod/config.yaml)
      --count int           count randomly chosen images
      --date string         date of the APOD image to retrieve (YYYY-MM-DD)
      --debug               for debug
      --end-date string     end of a date range (YYYY-MM-DD)
      --start-date string   start of a date range (YYYY-MM-DD)
      --thumbs              return the URL of video thumbnail

$ apod lookup | jq .
[
  {
    "copyright": "Serge\nBrunier, Jean-François Bax, David Vernet",
    "date": "2023-02-24",
    "explanation": "Planetary nebula Jones-Emberson 1 is the death shroud of a dying Sun-like star. It lies some 1,600 light-years from Earth toward the sharp-eyed constellation Lynx. About 4 light-years across, the expanding remnant of the dying star's atmosphere was shrugged off into interstellar space, as the star's central supply of hydrogen and then helium for fusion was finally depleted after billions of years. Visible near the center of the planetary nebula is what remains of the stellar core, a blue-hot white dwarf star.  Also known as PK 164 +31.1, the nebula is faint and very difficult to glimpse at a telescope's eyepiece. But this deep broadband image combining 22 hours of exposure time does show it off in exceptional detail. Stars within our own Milky Way galaxy as well as background galaxies across the universe are scattered through the clear field of view. Ephemeral on the cosmic stage, Jones-Emberson 1 will fade away over the next few thousand years. Its hot, central white dwarf star will take billions of years to cool.",
    "hdurl": "https://apod.nasa.gov/apod/image/2302/jonesemberson1.jpg",
    "media_type": "image",
    "service_version": "v1",
    "title": "Jones-Emberson 1",
    "url": "https://apod.nasa.gov/apod/image/2302/jonesemberson1_1024.jpg"
  }
]
```

### Download APOD data

```
$ apod download -h
Download NASA APOD data.

Usage:
  apod download [flags]

Aliases:
  download, dl, d

Flags:
  -d, --base-dir string   Base directory for daownload (default "./apod")
  -h, --help              help for download
      --include-nopd      Download no public domain images or videos
      --overwrite         Overwrite Download files

Global Flags:
      --api-key string      NASA API key
      --config string       Config file (default /home/spiegel/.config/apod/config.yaml)
      --count int           count randomly chosen images
      --date string         date of the APOD image to retrieve (YYYY-MM-DD)
      --debug               for debug
      --end-date string     end of a date range (YYYY-MM-DD)
      --start-date string   start of a date range (YYYY-MM-DD)
      --thumbs              return the URL of video thumbnail

$ apod download --include-nopd

$ LANG=C ls -l ~/ws/work/apod
total 4
drwxrwxr-x 2 spiegel spiegel 4096 Feb 24 19:58 2023-02-24

$ LANG=C ls -l ~/ws/work/apod/2023-02-24
total 3376
-rw-rw-r-- 1 spiegel spiegel 3094863 Feb 24 19:58 jonesemberson1.jpg
-rw-rw-r-- 1 spiegel spiegel  352679 Feb 24 19:58 jonesemberson1_1024.jpg
-rw-rw-r-- 1 spiegel spiegel    1365 Feb 24 19:58 metadata.json
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

## Reference

- [NASA Open APIs](https://api.nasa.gov/)

[apod]: https://github.com/goark/apod "goark/apod: CLI Tool for Astronomy Picture of the Day with NASA API"
