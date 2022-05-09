package serial

import (
	"time"

	"golang.org/x/sys/unix"
)

type Serial struct {
	baud baudrates
	data databitcs
	prty paritybit
	stop stopbitcs
	cbit uint32
	ibit uint32
	tout time.Duration
}

// NewSerial return a Serial data type with default settings
func NewSerial() *Serial {
	return &Serial{
		baud: B0009600, data: CS8, prty: ParityNon, stop: StopBitA,
		cbit: unix.CREAD | unix.CLOCAL, ibit: unix.IGNPAR,
	}
}

func (p *Serial) GetBuadRate() baudrates    { return p.baud }
func (p *Serial) GetDataBit() databitcs     { return p.data }
func (p *Serial) GetParity() paritybit      { return p.prty }
func (p *Serial) GetStopBit() stopbitcs     { return p.stop }
func (p *Serial) GetTimeout() time.Duration { return p.tout }

func (p *Serial) SetBuadRate(v baudrates) *Serial    { p.baud = v; return p }
func (p *Serial) SetDataBit(v databitcs) *Serial     { p.data = v; return p }
func (p *Serial) SetParity(v paritybit) *Serial      { p.prty = v; return p }
func (p *Serial) SetStopBit(v stopbitcs) *Serial     { p.stop = v; return p }
func (p *Serial) SetTimeout(d time.Duration) *Serial { p.tout = d; return p }

func (p *Serial) SetFlowControl(val hwcontrol) *Serial {
	switch val {
	case Hardware:
		p.ibit &^= unix.IXON | unix.IXOFF | unix.IXANY
		p.cbit |= unix.CRTSCTS
	case Softeare:
		p.ibit |= unix.IXON | unix.IXOFF | unix.IXANY
		p.cbit &^= unix.CRTSCTS
	case None:
		p.ibit &^= unix.IXON | unix.IXOFF | unix.IXANY
		p.cbit &^= unix.CRTSCTS
	}
	return p
}

func (p *Serial) GetFlowControl() hwcontrol {

	if p.cbit&unix.CRTSCTS == unix.CRTSCTS {
		return Hardware
	}

	if p.ibit&(unix.IXON|unix.IXOFF|
		unix.IXANY) == (unix.IXON | unix.IXOFF | unix.IXANY) {
		return Softeare
	}

	return None
}
