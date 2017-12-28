package main

import(
    "os"
    golog "log"
)

var emircLog *golog.Logger
var logInitDone bool = false

func logInit() {
    if logInitDone == false {
        emircLog = golog.New(os.Stdout, "[emirc]", golog.Ldate | golog.Ltime | golog.Lmicroseconds | golog.Lshortfile )
        logInitDone = true
    }
}

func log() *golog.Logger {
    logInit()
    return emircLog
}
