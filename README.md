# emersyx_irc

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

The function which must be used to create new `IRCBot` instances is:

```
func NewIRCBot(nick string, server string, port int, useSSL bool) emirc.IRCBot
```

## Credits

The underlying implementation is provided by [fluffle/goirc][2].

[1]: https://github.com/emersyx/emersyx_apis
[2]: https://github.com/emersyx/emersyx_apis/tree/master/emircapi
[2]: https://github.com/fluffle/goirc
