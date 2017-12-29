package main

import(
    "flag"
    "fmt"
    "testing"
    "time"
)

var nick *string = flag.String("nick", "", "IRC bot nick used during testing")
var channel *string = flag.String("channel", "", "IRC channel to join during testing")
var sendto *string = flag.String("sendto", "", "IRC user to send message to during testing")

func TestConnection(t *testing.T) {
    // when running go test -short
    if testing.Short() {
        fmt.Println("TestConnection not executed with -short argument")
    } else {
        opt := NewIRCOptions()

        bot, err := NewIRCBot(
            opt.Nick(*nick),
            opt.Server("chat.freenode.net", 6697, true),
        )
        if err != nil {
            fmt.Println(err)
        }

        err = bot.Connect()
        if err != nil {
            fmt.Println(err)
        } else {
            bot.Join(*channel)
            bot.Privmsg(*sendto, "hello world!")
            time.Sleep(20 * time.Second)
        }

        bot.Quit()
    }
}
