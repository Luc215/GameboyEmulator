package registers

type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f uint8
	h uint8
	l uint8
}
type FlagRegisters struct {
	zero        bool // bit 7
	subtraction bool // bit 6
	halfCarry   bool // bit 5
	carry       bool // bit 4
}

func (r *Registers) getBC() uint16 {
	return uint16(r.b)<<8 | uint16(r.c)
}
func (r *Registers) setBC(value uint16) {
	r.b = uint8((value & 0xFF00) >> 8) // gets top
	r.c = uint8((value & 0xFF) >> 8)   // gets bottom
}
func (r *Registers) getAF() uint16 {
	return uint16(r.a)<<8 | uint16(r.f)
}
func (r *Registers) setAF(value uint16) {
	r.a = uint8((value & 0xFF00) >> 8)
	r.f = uint8((value & 0xFF) >> 8)
}
func (r *Registers) getDE() uint16 {
	return uint16(r.d)<<8 | uint16(r.e)
}
func (r *Registers) setDE(value uint16) {
	r.d = uint8((value & 0xFF00) >> 8)
	r.e = uint8((value & 0xFF) >> 8)
}
func (r *Registers) getHL() uint16 {
	return uint16(r.h)<<8 | uint16(r.l)
}
func (r *Registers) setHL(value uint16) {
	r.h = uint8((value & 0xFF00) >> 8)
	r.l = uint8((value & 0xFF) >> 8)
}

func (flag *FlagRegisters) setFlagVal() uint8 {
	var zero int
	var subtraction int
	var halfCarry int
	var carry int

	if flag.zero {
		zero = 1
	} else {
		zero = 0
	}
	if flag.subtraction {
		subtraction = 1
	} else {
		subtraction = 0
	}
	if flag.halfCarry {
		halfCarry = 1
	} else {
		halfCarry = 0
	}
	if flag.carry {
		carry = 1
	} else {
		carry = 0
	}
	return uint8(zero<<7 | subtraction<<6 | halfCarry<<5 | carry<<4)

}

func main() {
}
