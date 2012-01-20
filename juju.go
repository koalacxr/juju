package main

import "launchpad.net/juju/go/control"
import "launchpad.net/juju/go/log"
import "os"

func main() {
    jc := control.JujuMainCommand()
    if err := jc.Parse(os.Args); err != nil {
        jc.Usage()
        os.Exit(2)
    }
	log.Debug = jc.Verbose()
	if err := log.SetFile(jc.Logfile()); err != nil {
		log.Printf(err)
        os.Exit(1)
	}
	if err := jc.Run(); err != nil {
		log.Printf(err)
		os.Exit(1)
	}
}
