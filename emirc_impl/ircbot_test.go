package main

import(
    "fmt"
    "testing"
    "time"
)

func TestConnection(t *testing.T) {
    // when running go test -short
    if testing.Short() {
        fmt.Println("TestConnection not executed with -short argument")
    } else {
        bot := emirc.NewIRCBot("emersyx", "chat.freenode.net", 6667, false)
        err := bot.Connect()
        if err != nil {
            fmt.Println(err)
        } else {
            bot.Join("#emersyx")
            bot.Privmsg("alegen", "hey dad")
            time.Sleep(20 * time.Second)
        }

        bot.Quit()
    }
}
