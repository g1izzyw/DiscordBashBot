package logger

import (
    "io"
    "io/ioutil"
    "log"
    "os"
)

type logger struct {	
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
}

func LoggerConstructor(handle io.Writer) *logger {
	l := new(logger)

    l.Info = log.New(handle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    l.Warning = log.New(handle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    l.Error = log.New(handle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    return l
}

// Trace.Println("I have something standard to say")
// Info.Println("Special Information")
// Warning.Println("There is something you need to know about")
// Error.Println("Something has failed")
