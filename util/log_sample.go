package util

import (
	"github.com/comail/colog"
	"log"
)

func logger() {

	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	log.Printf("trace: this is a trace log.")
	log.Printf("debug: this is a debug log.")
	log.Printf("info: this is an info log.")
	log.Printf("warn: this is a warning log.")
	log.Printf("error: this is an error log.")
	log.Printf("alert: this is an alert log.")

	log.Printf("this is a default level log.")

}
