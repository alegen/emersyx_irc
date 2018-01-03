package main

import(
    "errors"
    "emersyx.net/emersyx_apis/emcomapi"
)

// This file contains implementations of methods that are mandatory to implement the emircapi.IRCBot interfaces.

// Connect start the connection process of the IRC bot to the server. This is a blocking call. If the bot connects to
// the server without errors, then nil is returned. Otherwise an error with the appropriate message is returned.
func (bot *IRCBot) Connect() error {
    err := bot.api.Connect()
    return err
}

// IsConnected returned true if the bot is connected to the IRC server, otherwise it returns false.
func (bot *IRCBot) IsConnected() bool {
    return bot.api.Connected()
}

// Quit disconnects the IRC bot from the server. If the bot disconnects from the server without errors, then nil is
// returned. Otherwise an error with the appropriate message is returned.
func (bot *IRCBot) Quit() error {
    err := bot.api.Close()
    return err
}

// Join sends the command for the IRC bot to join a channel. The channel is specified in the method argument. If the bot
// joins the channel without errors, then nil is returned. Otherwise an error with the appropriate message is returned.
func (bot *IRCBot) Join(ch string) error {
    if bot.IsConnected() {
        bot.api.Join(ch)
        return nil
    }
    return errors.New("the IRCBot instance is not connected to any server")
}

// Privmsg sends either a message to an IRC channel or a private message to another user, depending on the method
// argument.
func (bot *IRCBot) Privmsg(to, msg string) error {
    if bot.IsConnected() {
        bot.api.Privmsg(to, msg)
        return nil
    }
    return errors.New("the IRCBot instance is not connected to any server")
}

// GetIdentifier returns the identifier of the receptor which generated this emersyx event.
func (bot *IRCBot) GetIdentifier() string {;
    return bot.identifier
}

// GetEventsChannel returns the emcomapi.Event channel through which emersyx events are pushed.
func (bot *IRCBot) GetEventsChannel() chan emcomapi.Event {
    return bot.messages
}
