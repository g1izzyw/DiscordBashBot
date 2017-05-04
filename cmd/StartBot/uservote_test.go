package uservote

import {
	"testing"

}

func TestYesGreateThanNo(t *testing.T) {  
	Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger


    total := LoggerConstructor(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}