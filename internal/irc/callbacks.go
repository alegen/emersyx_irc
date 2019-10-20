package main

import (
	"emersyx.net/emersyx/pkg/api/irc"
	goirc "github.com/fluffle/goirc/client"
)

// newMessage converts a Line object received from the underlying IRC library into an irc.Message object containing the
// same information.
func newMessage(id string, line *goirc.Line) irc.Message {
	var m irc.Message

	m.Source = id
	m.Raw = line.Raw
	m.Command = line.Cmd
	m.Origin = line.Nick
	m.Parameters = make([]string, len(line.Args))
	copy(m.Parameters, line.Args)

	return m
}

// channelCallback creates a callback for the underlying IRC library. The callback receives the Line object, converts it
// into a irc.Message object and sends it via the gateway event channel.
func channelCallback(gw *gateway) func(*goirc.Conn, *goirc.Line) {
	return func(conn *goirc.Conn, line *goirc.Line) {
		go func() {
			gw.messages <- newMessage(gw.Identifier, line)
		}()
	}
}

// loggingCallback creates a callback for the underlying IRC library. The callback receives the Line object, converts it
// into a irc.Message object and logs the contents.
func loggingCallback(gw *gateway) func(*goirc.Conn, *goirc.Line) {
	return func(conn *goirc.Conn, line *goirc.Line) {
		m := newMessage(gw.Identifier, line)
		gw.Log.Debugf("New message:\n")
		gw.Log.Debugf("Source      %s\n", m.Source)
		gw.Log.Debugf("Raw         %s\n", m.Raw)
		gw.Log.Debugf("Command     %s\n", m.Command)
		gw.Log.Debugf("Origin      %s\n", m.Origin)
		gw.Log.Debugf("Parameters:\n")
		for i, p := range m.Parameters {
			gw.Log.Debugf("%d. %s\n", i, p)
		}
		gw.Log.Debugf("-----\n")
	}
}
