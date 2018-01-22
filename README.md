# emersyx_irc [![Build Status][build-img]][build-url] [![Go Report Card][gorep-img]][gorep-url] [![GoDoc][godoc-img]][godoc-url]

IRC gateway (receptor & resource) for emersyx.

## Build

Source files in `emirc` provide the implementation of the go plugin. The easiest way to get all dependencies is by using
the [dep][4] tool. The commands to build the plugin are:

```
dep ensure
go build -buildmode=plugin -o emirc.so emirc/*
```

The resulting `emirc.so` file can then be used by emersyx core.

## Notes

The `IRCGateway` struct follows the APIs defined in the [emersyx_apis][1] repository, specifically those from the
[emircapi][2] folder.

The `NewIRCGateway` function must be used to create new `IRCGateway` instances. An example of how to use this function
can be found in the `emirc/ircgw_test.go` file.

## Credits

The underlying implementation is provided by [fluffle/goirc][3].

[build-img]: https://travis-ci.org/emersyx/emersyx_irc.svg?branch=master
[build-url]: https://travis-ci.org/emersyx/emersyx_irc
[gorep-img]: https://goreportcard.com/badge/github.com/emersyx/emersyx_irc
[gorep-url]: https://goreportcard.com/report/github.com/emersyx/emersyx_irc
[godoc-img]: https://godoc.org/emersyx.net/emersyx_irc?status.svg
[godoc-url]: https://godoc.org/emersyx.net/emersyx_irc
[1]: https://github.com/emersyx/emersyx_apis
[2]: https://github.com/emersyx/emersyx_apis/tree/master/emircapi
[3]: https://github.com/fluffle/goirc
[4]: https://github.com/golang/dep
