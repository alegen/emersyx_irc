# emersyx_irc [![Build Status][build-img]][build-url] [![Go Report Card][gorep-img]][gorep-url] [![GoDoc][godoc-img]][godoc-url]

IRC gateway (resource & peripheral) for emersyx.

## Build

Source files in `emirc` provide the implementation of the go plugin. The plugin can be built by running `make`. The
resulting `emirc.so` file can then be used by the emersyx core and router implementations from the [main emersyx
repository][emersyx-repo].

The `NewPeripheral` function must be used to create new `ircGateway` instances. An example of how to use this function
can be found in the `emirc/ircgw_test.go` file.

## Credits

The underlying implementation is provided by [fluffle/goirc][1].

[build-img]: https://travis-ci.org/emersyx/emersyx_irc.svg?branch=master
[build-url]: https://travis-ci.org/emersyx/emersyx_irc
[gorep-img]: https://goreportcard.com/badge/github.com/emersyx/emersyx_irc
[gorep-url]: https://goreportcard.com/report/github.com/emersyx/emersyx_irc
[godoc-img]: https://godoc.org/emersyx.net/emersyx_irc?status.svg
[godoc-url]: https://godoc.org/emersyx.net/emersyx_irc
[emersyx-repo]: https://github.com/emersyx/emersyx
[1]: https://github.com/fluffle/goirc
