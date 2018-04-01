package main

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"emersyx.net/emersyx_apis/emircapi"
	"flag"
	"fmt"
	"testing"
	"time"
)

var nick = flag.String("nick", "", "IRC gw nick used during testing")
var channel = flag.String("channel", "", "IRC channel to join during testing")
var sendto = flag.String("sendto", "", "IRC user to send message to during testing")

func TestConnection(t *testing.T) {
	opt := NewIRCOptions()

	// create a new IRCGateway
	gw, err := NewIRCGateway(
		opt.Identifier("emirc-test"),
		opt.Nick(*nick),
		opt.Server("chat.freenode.net", 6697, true),
	)

	// check for failure
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	// attempt to connect to the server
	err = gw.Connect()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	// if we reached this point, we will have to quit the server at the end
	defer gw.Quit()

	// for testing, we connect to the channel
	gw.Join(*channel)
	// and send a private message
	gw.Privmsg(*sendto, "hello world!")

	// when running go test with -short option, then do not test received messages
	if testing.Short() {
		// only wait for the connection and everything to happen
		time.Sleep(20)
	} else {
		messages := (gw.(emcomapi.Receptor)).GetEventsOutChannel()
		for i := 0; i < 20; i++ {
			m := <-messages
			// check the source identifier to be correct
			if m.GetSourceIdentifier() != "emirc-test" {
				fmt.Printf("Incorrect source identifier, got %s\n", m.GetSourceIdentifier())
				t.Fail()
				return
			}
			// print all the contents of the Message
			cm := m.(emircapi.Message)
			fmt.Printf("-----\n")
			fmt.Printf("Source      %s\n", cm.Source)
			fmt.Printf("Raw         %s\n", cm.Raw)
			fmt.Printf("Command     %s\n", cm.Command)
			fmt.Printf("Origin      %s\n", cm.Origin)
			fmt.Printf("Parameters:\n")
			for i, p := range cm.Parameters {
				fmt.Printf("%d. %s\n", i, p)
			}
		}
	}
}
