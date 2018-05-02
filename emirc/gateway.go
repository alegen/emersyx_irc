package main

import (
	"emersyx.net/emersyx/api"
	"github.com/BurntSushi/toml"
	goirc "github.com/fluffle/goirc/client"
)

// The ircGateway struct defines the implementation of an IRC receptor and resource. The struct implements the
// api.IRCGateway and api.Receptor interfaces.
type ircGateway struct {
	api.PeripheralBase
	api      *goirc.Conn
	config   *goirc.Config
	messages chan api.Event
}

// NewPeripheral creates a new api.IRCGateway instance and applies to configuration specified in the arguments.
func NewPeripheral(opts api.PeripheralOptions) (api.Peripheral, error) {
	var err error

	// create a new ircGateway and initialize the base
	gw := new(ircGateway)
	gw.InitializeBase(opts)

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

	// connect to the server
	gw.connect()

	return gw, nil
}

// initCallbacks sets the callback functions for the internally used goirc library.
func (gw *ircGateway) initCallbacks() {
	gw.api.HandleFunc(goirc.PRIVMSG, channelCallback(gw))
	gw.api.HandleFunc(goirc.JOIN, channelCallback(gw))
	gw.api.HandleFunc(goirc.QUIT, channelCallback(gw))
	gw.api.HandleFunc(goirc.PART, channelCallback(gw))
}
