package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	goupnp "github.com/NebulousLabs/go-upnp"
)

var usage = `**********
Usage: forward port1 port2 port3 etc
Option: -close to close selected ports
**********
`

func main() {
	usage := func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Usage = usage
	close := flag.Bool("close", false, "Close ports")
	flag.Parse()
	if len(os.Args) <= 1 {
		usage()
		return
	}
	f, err := goupnp.Discover()
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range os.Args[1:] {
		if v == "-close" {
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
			continue
		}
		if *close {
			if err := f.Clear(uint16(i)); err != nil {
				log.Println(err)
			}
		} else {
			if err := f.Forward(uint16(i), "go forwarder"); err != nil {
				log.Println(err)
			}
		}
	}
}
