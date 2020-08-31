package smilespars

type tokenType int

const (
	lRoundParen tokenType = iota
	rRoundParen
	lSquareParen
	rSquareParen
	minus
	minusMinus
	plus
	plusPlus
	colon
	equal
	hash
	dollar
	slash
	backslash
	percent
	dot
	digit
	aliphaticOrganic
	aromaticOrganic
	elementSymbol
	aromaticSymbol
	hSymbol
	unknown
	chiral
	space
	tab
	lineFeed
	carriageReturn
	endOfString
	endOfFile
)

type token struct {
	tType tokenType
	attribute string
	start, end position
}

func newToken(tType tokenType, attribute string, start, end position) token {
	var t token
	t.tType = tType
	t.attribute = attribute
	t.start = start
	t.end = end
	return t
}

func (t token) String() string {
	return t.tType.String() + " " + t.start.String() + "-" + t.end.String() + ": " + t.attribute
}

func (tType tokenType) String() string {
	switch tType {
	case lRoundParen:
		return "L_ROUND_PAREN"
	case rRoundParen:
		return "R_ROUND_PAREN"
	case lSquareParen:
		return "L_SQUARE_PAREN"
	case rSquareParen:
		return "R_SQUARE_PAREN"
	case minus:
		return "MINUS"
	case minusMinus:
		return "MINUS_MINUS"
	case plus:
		return "PLUS"
	case plusPlus:
		return "PLUS_PLUS"
	case colon:
		return "COLON"
	case equal:
		return "EQUAL"
	case hash:
		return "HASH"
	case dollar:
		return "DOLLAR"
	case slash:
		return "SLASH"
	case backslash:
		return "BACKSLASH"
	case percent:
		return "PERCENT"
	case dot:
		return "DOT"
	case digit:
		return "DIGIT"
	case aliphaticOrganic:
		return "ALIPHATIC_ORGANIC"
	case aromaticOrganic:
		return "AROMATIC_ORGANIC"
	case elementSymbol:
		return "ELEMENT_SYMBOL"
	case aromaticSymbol:
		return "AROMATIC_SYMBOL"
	case hSymbol:
		return "H_SYMBOL"
	case unknown:
		return "UNKNOWN"
	case chiral:
		return "CHIRAL"
	case space:
		return "SPACE"
	case tab:
		return "TAB"
	case lineFeed:
		return "LINE_FEED------------------------------------"
	case carriageReturn:
		return "CARRIAGE_RETURN"
	case endOfString:
		return "END_OF_STRING"
	case endOfFile:
		return "END_OF_FILE"
	default:
		return "INVALID_TOKEN"
	}
}

func (t token) isTerminator() bool {
	return t.tType == lineFeed || t.tType == carriageReturn || t.tType == endOfString ||
		t.tType == space || t.tType == tab || t.tType == endOfFile
}

func (t token) isBond() bool {
	return t.tType == minus || t.tType == equal || t.tType == hash || t.tType == dollar ||
		t.tType == colon || t.tType == slash || t.tType == backslash
}

func (t token) isElementSymbol() bool {
	return t.tType == hSymbol || t.tType == aliphaticOrganic || t.tType == elementSymbol
}

func (t token) isAromaticSymbol() bool {
	return t.tType == aromaticOrganic || t.tType == aromaticSymbol
}

func (t token) isSymbol() bool {
	return t.isElementSymbol() || t.isAromaticSymbol() || t.tType == unknown
}
