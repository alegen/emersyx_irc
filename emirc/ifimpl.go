package main

import(
    "errors"
    "emersyx.net/emersyx_apis/emircapi"
)

// This file contains implementations of methods that are mandatory for the emirc.IRCBot struct to implement the
// emircapi.IRCBot and emcomapi.Receptor interfaces.

// This method returns an IRCOptionsGenerator which implement the methods to configure an IRCBot instance.
func (bot *IRCBot) GetOptionsGenerator() emircapi.IRCOptionsGenerator {
    var optgen IRCOptionsGenerator
    return optgen
}

// This method starts the connection process to the selected IRC server. This is a blocking method. When the method
// returns, the IRCBot instance is connected to the IRC server if the return value is nil. Otherwise, an error occured.
func (bot *IRCBot) Connect() error {
    err := bot.api.Connect()
    return err
}

// This method returns a boolean which is true if the bot is connected to the server, and false otherwise.
func (bot *IRCBot) IsConnected() bool {
    return bot.api.Connected()
}

// This method disconnects the bot from the IRC server. This is a blocking method. When the method returns, the IRCBot
// instance is not connected to the IRC server anymore.
func (bot *IRCBot) Quit() error {
    err := bot.api.Close()
    return err
}

// This method sends a JOIN command to the IRC server. The argument specifies the channel to be joined. If the IRCBot
// instance is not connected to any IRC server, then an error is returned.
func (bot *IRCBot) Join(ch string) error {
    if bot.IsConnected() {
        bot.api.Join(ch)
        return nil
    } else {
        return errors.New("The IRCBot instance is not connected to any server.");
    }
}

// This method sends a PRIVMSG command to the IRC server. The first argument specifies the destination (i.e. either a
// user or a channel) and the second argument is the actual message. If the IRCBot instance is not connected to any IRC
// server, then an error is returned.
func (bot *IRCBot) Privmsg(to, msg string) error {
    if bot.IsConnected() {
        bot.api.Privmsg(to, msg)
        return nil
    } else {
        return errors.New("The IRCBot instance is not connected to any server.");
    }
}

// This method returns the Messages field of the IRCBot struct.
func (bot *IRCBot) GetEventsChannel() chan emircapi.Message {
    return bot.Messages;
}
