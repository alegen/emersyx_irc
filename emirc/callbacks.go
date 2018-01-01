package main

import(
    "emersyx.net/emersyx_apis/emircapi"
    "log"
    irc "github.com/fluffle/goirc/client"
)

// This function converts a Line object received from the underlying IRC library into an emircapi.Message object
// containing the same information.
func newMessage(id string, line *irc.Line) emircapi.Message {
    var m emircapi.Message

    m.Source = id
    m.Raw = line.Raw
    m.Command = line.Cmd
    m.Origin = line.Nick
    m.Parameters = make([]string, len(line.Args))
    copy(m.Parameters, line.Args)

    return m
}

// This function creates a callback for the underlying IRC library. The callback receives the Line object, converts it
// into a emircapi.Message object and sends it via the IRCBot event channel.
func channelCallback(bot *IRCBot) func(*irc.Conn, *irc.Line) {
    return func(conn *irc.Conn, line *irc.Line) {
        go func() {
            bot.messages <- newMessage(bot.identifier, line)
        }()
    }
}

// This function creates a callback for the underlying IRC library. The callback receives the Line object, converts it
// into a emircapi.Message object and logs the contents.
func loggingCallback(bot *IRCBot, logger *log.Logger) func(*irc.Conn, *irc.Line) {
    return func(conn *irc.Conn, line *irc.Line) {
        m := newMessage(bot.identifier, line)
        logger.Printf("New message:\n")
        logger.Printf("Source      %s\n", m.Source)
        logger.Printf("Raw         %s\n", m.Raw)
        logger.Printf("Command     %s\n", m.Command)
        logger.Printf("Origin      %s\n", m.Origin)
        logger.Printf("Parameters:\n")
        for i, p := range m.Parameters {
            logger.Printf("%d. %s\n", i, p)
        }
        logger.Printf("-----\n")
    }
}
