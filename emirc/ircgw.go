package main

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"emersyx.net/emersyx_apis/emircapi"
	"errors"
	irc "github.com/fluffle/goirc/client"
)

// The IRCGateway struct defines the implementation of an IRC receptor and resource. The struct implements the
// emircapi.IRCGateway and emcomapi.Receptor interfaces.
type IRCGateway struct {
	api        *irc.Conn
	cfg        *irc.Config
	identifier string
	messages   chan emcomapi.Event
}

// NewIRCGateway creates a new emircapi.IRCGateway instance and applies to configuration specified in the arguments.
func NewIRCGateway(options ...func(emircapi.IRCGateway) error) (emircapi.IRCGateway, error) {
	gw := new(IRCGateway)

	// create the messages channel
	gw.messages = make(chan emcomapi.Event)

	// create a Config object for the underlying library
	gw.cfg = irc.NewConfig("placeholder")

	// override several default values from the underlying library
	gw.cfg.Me.Ident = "emersyx"
	gw.cfg.Me.Name = "emersyx"
	gw.cfg.Version = "emersyx"
	gw.cfg.SSL = false
	gw.cfg.QuitMessage = "bye"

	// standard function for generating new nicks
	gw.cfg.NewNick = func(n string) string { return n + "^" }

	// apply the configuration options received as arguments
	// this configuration method is inspired from
	// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
	for _, option := range options {
		err := option(gw)
		if err != nil {
			return nil, err
		}
	}

	// check if the mandatory identifier value has been set
	if len(gw.identifier) == 0 {
		return nil, errors.New("identifier option has not been set")
	}

	// create the underlying Conn object
	gw.api = irc.Client(gw.cfg)

	// initialize callbacks
	gw.initCallbacks()

	return gw, nil
}

func (gw *IRCGateway) initCallbacks() {
	gw.api.HandleFunc(irc.PRIVMSG, channelCallback(gw))
	gw.api.HandleFunc(irc.JOIN, channelCallback(gw))
	gw.api.HandleFunc(irc.QUIT, channelCallback(gw))
	gw.api.HandleFunc(irc.PART, channelCallback(gw))
}
