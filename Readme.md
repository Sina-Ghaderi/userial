
# userial

linux serial package for handling serial ports with go, this package is just for serial ports and may not be suitable for all kind of terminals.. see [termios(3)](https://linux.die.net/man/3/termios)

- this package should work with generic serial ports and standard baudrate
- userial currently only works with linux (im not planing to use this with any other os'es)


### how to use userial

userial lives on [github](https://github.com/sina-ghaderi/userial) and [snix](https://git.snix.ir/userial) git repositories, in order to use this package you should run one of the following commands


get userial from snix repository `go mod download snix.ir/userial` or github `go mod download github.com/sina-ghaderi/userial`

### examples
simple read and write with cisco router serial port
```go
package main

import (
	"io"
	"log"
	"os"

	"snix.ir/userial"
)

// we need something to read,  so we gonna mess with terminal..
const cisco_router_cmd = `

enable
conf t
int gig0/0
ip addr 192.168.122.200 255.255.255.0 
no shut
do show ip int br

`

func main() {
	p, err := serial.NewSerial().OpenPort("/dev/ttyACM0") 
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

```
note that `/dev/ttyACM0` indicate serial port device which opened with default configuration.
above program read data from serial port and write it on stdout..

```

Router>
Router>
Router>enable
Router#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
Router(config)#int gig0/0
Router(config-if)#ip addr 192.168.122.200 255.255.255.0 
Router(config-if)#no shut
Router(config-if)#
Router(config-if)#do show ip int br
Interface                  IP-Address      OK? Method Status                Protocol
Embedded-Service-Engine0/0 unassigned      YES unset  administratively down down    
GigabitEthernet0/0         192.168.122.200 YES manual down                  down   
GigabitEthernet0/1         unassigned      YES unset  administratively down down    
Router(config-if)#
```

### licence
Copyright 2021 SNIX LLC sina@snix.ir
This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License version 2 as published by the Free Software Foundation.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
