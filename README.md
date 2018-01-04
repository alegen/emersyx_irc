# emersyx_irc [![Build Status][build-img]][build-url] [![Go Report Card][gorep-img]][gorep-url]
IRC receptor and resource for emersyx.

## Build

Source files in `emirc` provide the implementation of the go plugin. They have to be built using the command:

```
go build -buildmode=plugin -o emirc.so emirc/*
```

The resulting `emirc.so` file can then be used by emersyx core.

## Notes

The `IRCBot` struct follows the APIs defined in the [emersyx_apis][1] repository, specifically those from the
[emircapi][2] folder.

The `NewIRCBot` function must be used to create new `IRCBot` instances. An example of how to use this function can be
found in the `emirc/ircbot_test.go` file.

## Credits

The underlying implementation is provided by [fluffle/goirc][3].

[build-img]: https://travis-ci.org/emersyx/emersyx_irc.svg?branch=master
[build-url]: https://travis-ci.org/emersyx/emersyx_irc
[gorep-img]: https://goreportcard.com/badge/github.com/emersyx/emersyx_irc
[gorep-url]: https://goreportcard.com/report/github.com/emersyx/emersyx_irc
[1]: https://github.com/emersyx/emersyx_apis
[2]: https://github.com/emersyx/emersyx_apis/tree/master/emircapi
[3]: https://github.com/fluffle/goirc
