package main

import (
	"emersyx.net/emersyx/api"
	"emersyx.net/emersyx/log"
	"errors"
	"github.com/BurntSushi/toml"
	goirc "github.com/fluffle/goirc/client"
)

// The ircGateway struct defines the implementation of an IRC receptor and resource. The struct implements the
// api.IRCGateway and api.Receptor interfaces.
type ircGateway struct {
	core       api.Core
	api        *goirc.Conn
	config     *goirc.Config
	log        *log.EmersyxLogger
	identifier string
	messages   chan api.Event
}

// NewPeripheral creates a new api.IRCGateway instance and applies to configuration specified in the arguments.
func NewPeripheral(opts api.PeripheralOptions) (api.Peripheral, error) {
	var err error

	// validate identifier in options
	if len(opts.Identifier) == 0 {
		return nil, errors.New("identifier cannot have 0 length")
	}

	gw := new(ircGateway)

	// create the messages channel
	gw.messages = make(chan api.Event)

	// create a Config object for the underlying library
	gw.config = goirc.NewConfig("placeholder")

	// override several default values from the underlying library
	gw.config.Me.Ident = "emersyx"
	gw.config.Me.Name = "emersyx"
	gw.config.Version = "emersyx"
	gw.config.SSL = false
	gw.config.QuitMessage = "bye"

	// standard function for generating new nicks
	gw.config.NewNick = func(n string) string { return n + "^" }

	// generate a bare logger, to be updated via options
	gw.log, err = log.NewEmersyxLogger(nil, "", log.ELNone)
	if err != nil {
		return nil, errors.New("could not create a bare logger")
	}

	// apply the options received as argument
	gw.identifier = opts.Identifier
	gw.core = opts.Core
	gw.log.SetOutput(opts.LogWriter)
	gw.log.SetLevel(opts.LogLevel)
	gw.log.SetComponentID(gw.identifier)

	// apply the extended options from the config file
	config := new(ircGatewayConfig)
	if _, err = toml.DecodeFile(opts.ConfigPath, config); err != nil {
		return nil, err
	}
	if err = config.validate(); err != nil {
		return nil, err
	}
	config.apply(gw)

	// create the underlying Conn object
	gw.api = goirc.Client(gw.config)

	// initialize callbacks
	gw.initCallbacks()

	return gw, nil
}

// initCallbacks sets the callback functions for the internally used goirc library.
func (gw *ircGateway) initCallbacks() {
	gw.api.HandleFunc(goirc.PRIVMSG, channelCallback(gw))
	gw.api.HandleFunc(goirc.JOIN, channelCallback(gw))
	gw.api.HandleFunc(goirc.QUIT, channelCallback(gw))
	gw.api.HandleFunc(goirc.PART, channelCallback(gw))
}
