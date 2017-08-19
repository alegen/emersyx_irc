# emersyx_irc

IRC receptor and resource for emersyx.

## Usage

### Plugin implementation

Source files in `emirc_impl` provide the implementation of the go plugin. They have to be built using the command:

```
go build -buildmode=plugin -o emersyx_irc.so emirc_impl/*
```

### Plugin interface

Source files in `emirc` provide the interface to go plugin. This package has to be imported into projects which use the
plugin. The function which must be used to create new `IRCBot` instances (which implement the `emirc.IRCBot` interface)
is:

```
func NewIRCBot(nick string, server string, port int, useSSL bool) emirc.IRCBot
```

## Credits

The underlying implementation is provided by [fluffle/goirc][1].

[1]: https://github.com/fluffle/goirc
