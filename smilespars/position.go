package smilespars

import "strconv"

type position struct {
	line, offset, index int
}

func newPosition() position {
	var p position
	p.line = 1
	p.offset = 1
	p.index = 0
	return p
}

func (p *position) shiftRight() {
	p.offset++
}

func (p *position) shiftNextLine() {
	p.line++
	p.offset = 1
}

func (p position) String() string {
	return "(" + strconv.Itoa(p.line) + ", " + strconv.Itoa(p.offset) + ")"
}