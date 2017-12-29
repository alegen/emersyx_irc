package main

import(
    "crypto/tls"
    "errors"
    "math"
    "strconv"
    "emersyx.net/emersyx_apis/emircapi"
)

// This struct implements the emircapi.IRCOptionsGenerator interface. It generates configuration functions for the
// emircapi.IRCBot implementation.
type IRCOptionsGenerator struct {
}

func (g IRCOptionsGenerator) Nick(nick string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(nick) == 0 {
            return errors.New("Nick cannot have zero length.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Nick = nick
        return nil
    }
}

func (g IRCOptionsGenerator) Ident(ident string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(ident) == 0 {
            return errors.New("Ident cannot have zero length.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Ident = ident
        return nil
    }
}

func (g IRCOptionsGenerator) Name(name string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(name) == 0 {
            return errors.New("Name cannot have zero length.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Name = name
        return nil
    }
}

func (g IRCOptionsGenerator) Version(version string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.Version = version
        return nil
    }
}

func (g IRCOptionsGenerator) Server(address string, port uint, useSSL bool) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(address) == 0 {
            return errors.New("Address cannot have zero length.")
        }
        if float64(port) > math.Pow(2, 16) - 1 {
            return errors.New("Port value is invalid.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
        cbot.cfg.Server = address + ":" + strconv.Itoa(int(port))
        cbot.cfg.SSL = useSSL;
        if cbot.cfg.SSL {
            cbot.cfg.SSLConfig = &tls.Config{ ServerName: address }
        }
        return nil
    }
}

func (g IRCOptionsGenerator) UseSSL(useSSL bool) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
        cbot.cfg.SSL = useSSL;
        return nil
    }
}

func (g IRCOptionsGenerator) QuitMessage(message string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.QuitMessage = message;
        return nil
    }
}
