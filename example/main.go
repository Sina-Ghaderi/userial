package main

import (
	"io"
	"log"
	"os"

	"github.com/sina-ghaderi/userial"
)

// we need something to read,  so we gonna mess with terminal..
const cisco_router_cmd = `


enable
show ip int br
conf t
int gig0/0
ip addr 192.168.122.200 255.255.255.0 
no shut

do show ip int br

`

func main() {

	// default conf
	p, err := userial.NewSerial().OpenPort("/dev/ttyACM0")

	// serial port with custom config
	//
	// p_dev := "/dev/ttyACM0"
	// s := userial.NewSerial().SetBuadRate(userial.B0019200)
	// p, err = s.SetParity(userial.ParityEvn).OpenPort(p_dev)

	if err != nil {
		log.Fatal(err)
	}

	defer p.Close()

	if _, err := p.Write([]byte(cisco_router_cmd)); err != nil {
		panic(err)
	}

	if _, err := io.Copy(os.Stdout, p); err != nil {
		panic(err)
	}

}
