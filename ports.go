package userial

import (
	"math"
	"os"
	"time"

	"golang.org/x/sys/unix"
)

type Port struct{ f *os.File }

// OpenPort opens a named serial device and return a Port data type and error
func (spops *Serial) OpenPort(path string) (*Port, error) {

	f, err := os.OpenFile(path, unix.O_RDWR|unix.O_NOCTTY|unix.O_NONBLOCK, 0666)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil && f != nil {
			f.Close()
		}
	}()

	fg := uint32(spops.baud) | uint32(spops.data) | uint32(spops.stop) | spops.cbit

	switch spops.prty {
	case 1:
		fg |= unix.PARENB
	case 2:
		fg |= unix.PARENB
		fg |= unix.PARODD
	}

	t := unix.Termios{
		Cflag: fg, Iflag: spops.ibit,
		Ispeed: uint32(spops.baud),
		Ospeed: uint32(spops.baud),
	}

	t.Cc[unix.VMIN], t.Cc[unix.VTIME] = timeout(spops.tout)

	if err := unix.IoctlSetTermios(
		int(f.Fd()), unix.TCSETS, &t); err != nil {
		return nil, err
	}

	if err := unix.SetNonblock(int(f.Fd()), false); err != nil {
		return nil, err
	}

	return &Port{f: f}, err

}

// Read implement read() method for serial port
func (p *Port) Read(b []byte) (int, error) { return p.f.Read(b) }

// Write implement write() method for serial port
func (p *Port) Write(b []byte) (int, error) { return p.f.Write(b) }

// Close implement close() method for serial port
func (p *Port) Close() error { return p.f.Close() }

// Flush implement flush() method for serial port
func (p *Port) Flush() error {
	return unix.IoctlSetInt(int(p.f.Fd()), unix.TCFLSH, unix.TCIOFLUSH)
}

// SendBreak Sends Break Signal
func (p *Port) SendBreak(d time.Duration) error {
	return unix.IoctlSetInt(int(p.f.Fd()), unix.TCSBRKP, int(d))
}

// Available returns how many bytes are unused in the buffer
func (p *Port) Available() (int, error) {
	return unix.IoctlGetInt(int(p.f.Fd()), unix.TIOCINQ)
}

// Buffered returns the number of bytes that have been written into the current buffer
func (p *Port) Buffered() (int, error) {
	return unix.IoctlGetInt(int(p.f.Fd()), unix.TIOCOUTQ)
}

// Expose serial fd to use in some special cases, ioctl etc...
func (p *Port) File() *os.File { return p.f }

func timeout(d time.Duration) (byte, byte) {
	if d > 0x0 {
		t := math.Min(math.Max(float64(d.Nanoseconds()/1e6/100), 0x01), 0xff)
		return 0x0, byte(t)
	}
	return 0x1, 0x0
}
