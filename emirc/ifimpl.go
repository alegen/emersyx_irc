package main

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"errors"
)

// This file contains implementations of methods that are mandatory to implement the emircapi.IRCGateway interfaces.

// Connect start the connection process of the IRC gateway to the server. This is a blocking call. If the gateway
// connects to the server without errors, then nil is returned. Otherwise an error with the appropriate message is
// returned.
func (gw *IRCGateway) Connect() error {
	err := gw.api.Connect()
	return err
}

// IsConnected returned true if the gateway is connected to the IRC server, otherwise it returns false.
func (gw *IRCGateway) IsConnected() bool {
	return gw.api.Connected()
}

// Quit disconnects the IRC gateway from the server. If the gateway disconnects from the server without errors, then nil
// is returned. Otherwise an error with the appropriate message is returned.
func (gw *IRCGateway) Quit() error {
	err := gw.api.Close()
	return err
}

// Join sends the command for the IRC gateway to join a channel. The channel is specified in the method argument. If the
// gateway joins the channel without errors, then nil is returned. Otherwise an error with the appropriate message is
// returned.
func (gw *IRCGateway) Join(ch string) error {
	if gw.IsConnected() {
		gw.api.Join(ch)
		return nil
	}
	return errors.New("the IRCGateway instance is not connected to any server")
}

// Privmsg sends either a message to an IRC channel or a private message to another user, depending on the method
// argument.
func (gw *IRCGateway) Privmsg(to, msg string) error {
	if gw.IsConnected() {
		gw.api.Privmsg(to, msg)
		return nil
	}
	return errors.New("the IRCGateway instance is not connected to any server")
}

// GetIdentifier returns the identifier of this receptor.
func (gw *IRCGateway) GetIdentifier() string {
	return gw.identifier
}

// GetEventsChannel returns the emcomapi.Event channel through which emersyx events are pushed by this receptor.
func (gw *IRCGateway) GetEventsChannel() chan emcomapi.Event {
	return gw.messages
}
