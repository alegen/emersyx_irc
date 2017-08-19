package emirc

type IRCBot interface {
    GetEventsChannel() chan interface{}
    Connect() error
    Quit()
    Join(string)
    Privmsg(string, string)
}
