package smilespars

type message struct {
	text string
	pos position
}

func (m message) String() string {
	return m.pos.String() + ": " + m.text
}