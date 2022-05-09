package serial

import (
	"golang.org/x/sys/unix"
)

type (
	baudrates rune
	paritybit byte
	databitcs byte
	stopbitcs byte
	hwcontrol byte
)

const (
	Hardware hwcontrol = iota
	Softeare
	None
)

const (
	B0000050 baudrates = 0x01 + iota
	B0000075
	B0000110
	B0000134
	B0000150
	B0000200
	B0000300
	B0000600
	B0001200
	B0001800
	B0002400
	B0004800
	B0009600
	B0019200
	B0038400

	B0057600 baudrates = 0xFF2 + iota
	B0115200
	B0230400
	B0460800
	B0500000
	B0576000
	B0921600
	B1000000
	B1152000
	B1500000
	B2000000
	B2500000
	B3000000
	B3500000
	B4000000
)

const (
	ParityNon paritybit = iota
	ParityEvn
	ParityOdd
)

const (
	StopBitA stopbitcs = 0x00        // stop bit 1
	StopBitB stopbitcs = unix.CSTOPB // stop bit 2
)

const (
	CS5 databitcs = unix.CS5 // databit 5
	CS6 databitcs = unix.CS6 // databit 6
	CS7 databitcs = unix.CS7 // databit 7
	CS8 databitcs = unix.CS8 // databit 8 (default)
)
