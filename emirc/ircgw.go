package main

import (
	"emersyx.net/emersyx/api"
	"emersyx.net/emersyx/log"
	"errors"
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
func NewPeripheral(options ...func(api.Peripheral) error) (api.Peripheral, error) {
	var err error

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

	// apply the configuration options received as arguments
	for _, option := range options {
		err := option(gw)
		if err != nil {
			return nil, err
		}
	}

	if len(gw.identifier) == 0 {
		return nil, errors.New("identifier option has not been set")
	}

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
