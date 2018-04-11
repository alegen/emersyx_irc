package main

import (
	"emersyx.net/emersyx/api"
	irc "github.com/fluffle/goirc/client"
)

// newMessage converts a Line object received from the underlying IRC library into an api.IRCMessage object containing the
// same information.
func newMessage(id string, line *irc.Line) api.IRCMessage {
	var m api.IRCMessage

	m.Source = id
	m.Raw = line.Raw
	m.Command = line.Cmd
	m.Origin = line.Nick
	m.Parameters = make([]string, len(line.Args))
	copy(m.Parameters, line.Args)

	return m
}

// channelCallback creates a callback for the underlying IRC library. The callback receives the Line object, converts it
// into a api.Message object and sends it via the ircGateway event channel.
func channelCallback(gw *ircGateway) func(*irc.Conn, *irc.Line) {
	return func(conn *irc.Conn, line *irc.Line) {
		go func() {
			gw.messages <- newMessage(gw.identifier, line)
		}()
	}
}

// loggingCallback creates a callback for the underlying IRC library. The callback receives the Line object, converts it
// into a api.Message object and logs the contents.
func loggingCallback(gw *ircGateway) func(*irc.Conn, *irc.Line) {
	return func(conn *irc.Conn, line *irc.Line) {
		m := newMessage(gw.identifier, line)
		gw.log.Debugf("New message:\n")
		gw.log.Debugf("Source      %s\n", m.Source)
		gw.log.Debugf("Raw         %s\n", m.Raw)
		gw.log.Debugf("Command     %s\n", m.Command)
		gw.log.Debugf("Origin      %s\n", m.Origin)
		gw.log.Debugf("Parameters:\n")
		for i, p := range m.Parameters {
			gw.log.Debugf("%d. %s\n", i, p)
		}
		gw.log.Debugf("-----\n")
	}
}
