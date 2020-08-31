package smilespars

import (
	"log"
	"regexp"
)

type lexer struct {
	text string
	messages []message
	currentPosition position
	roundParensStack []byte

}

func (l lexer) isEnd() bool {
	return l.currentPosition.index >= len(l.text)
}

func (l lexer) currentSymbol() byte {
	return l.text[l.currentPosition.index]
}

func newLexer(text string) *lexer {
	l := new(lexer)
	l.text = text
	l.currentPosition = newPosition()
	return l
}

func (l *lexer) run() []token {
	var tokens []token
	for (!l.isEnd()) {
		t, msg := l.nextToken(message{})
		if msg != (message{}) {
			l.messages = append(l.messages, msg)
		}
		if t != (token{}) {
			tokens = append(tokens, t)
		}
	}

	if len(l.roundParensStack) != 0 {
		log.Fatalf("not all round parenthesises were closed at the end of input")
	}

	if len(l.messages) > 0 {
		log.Println("lexer produced following messages during work:")
		for i := 0; i < len(l.messages); i++ {
			log.Println(l.messages[i])
		}
	}

	tokens = append(tokens, newToken(endOfFile, "", l.currentPosition, l.currentPosition))
	return tokens
}

func (l *lexer) nextToken(msg message) (token, message) {
	for !l.isEnd() {
		if startPos, tType, ok := l.checkSpaceToken(); ok {
			return newToken(tType, "", startPos, l.currentPosition), msg
		}
		if startPos, tType, ok := l.checkSignToken(); ok {
			return newToken(tType, "", startPos, l.currentPosition), msg
		}
		if startPos, attribute, ok := l.checkDigitToken(); ok {
			return newToken(digit, attribute, startPos, l.currentPosition), msg
		}
		if startPos, attribute, ok := l.checkChiralToken(); ok {
			return newToken(chiral, attribute, startPos, l.currentPosition), msg
		}
		if startPos, tType, attribute, ok := l.checkElementToken(); ok {
			return newToken(tType, attribute, startPos, l.currentPosition), msg
		}

		if msg == (message{}) {
			msg.pos = l.currentPosition
			msg.text = "syntax error"
		}
		l.shiftPosition()
	}

	return token{}, msg
}

func (l *lexer) checkSpaceToken() (position, tokenType, bool) {
	startPos := l.currentPosition
	var tType tokenType
	switch l.currentSymbol() {
	case '\n':
		tType = lineFeed
	case ' ':
		tType = space
	case '\t':
		tType = tab
	case '\r':
		if l.currentPosition.index+1 < len(l.text) && l.text[l.currentPosition.index+1] == '\n' {
			tType = endOfString
		} else {
			tType = carriageReturn
		}
	default:
		return position{}, -1, false
	}

	l.shiftPosition()
	return startPos, tType, true
}

func (l *lexer) checkSignToken() (position, tokenType, bool) {
	startPos := l.currentPosition
	var tType tokenType
	switch l.currentSymbol() {
	case '(':
		l.roundParensStack = append(l.roundParensStack, '(')
		tType = lRoundParen
		break
	case ')':
		if len(l.roundParensStack) == 0 {
			log.Fatalf(l.currentPosition.String() + ": close round paren ')' met before open one '('")
		}
		l.roundParensStack = l.roundParensStack[:len(l.roundParensStack)-1]
		tType = rRoundParen
		break
	case '[':
		tType = lSquareParen
		break
	case ']':
		tType = rSquareParen
		break
	case '-':
		if l.currentPosition.index+1 < len(l.text) && l.text[l.currentPosition.index+1] == '-' {
			tType = minusMinus
			l.shiftPosition()
		} else {
			tType = minus
		}
		break
	case '+':
		if l.currentPosition.index+1 < len(l.text) && l.text[l.currentPosition.index+1] == '+' {
			tType = plusPlus
			l.shiftPosition()
		} else {
			tType = plus
		}
		break
	case ':':
		tType = colon
		break
	case '=':
		tType = equal
		break
	case '#':
		tType = hash
		break
	case '$':
		tType = dollar
		break
	case '/':
		tType = slash
		break
	case '\\':
		tType = backslash
		break
	case '%':
		tType = percent
		break
	case '.':
		tType = dot
		break
	case '*':
		tType = unknown
		break
	default:
		return position{}, -1, false
	}

	l.shiftPosition()
	return startPos, tType, true
}

func (l *lexer) checkDigitToken() (position, string, bool) {
	if isDigit(l.currentSymbol()) {
		startPos := l.currentPosition
		attribute := string(l.currentSymbol())
		l.shiftPosition()
		return startPos, attribute, true
	}
	return position{}, "", false
}

func (l *lexer) checkChiralToken() (position, string, bool) {
	startPos := l.currentPosition
	maxPossibleAttributeLength := maxPossibleChiralLength
	if maxPossibleAttributeLength > len(l.text) - l.currentPosition.index {
		maxPossibleAttributeLength = len(l.text) - l.currentPosition.index
	}
	possibleAttribute := l.text[l.currentPosition.index : l.currentPosition.index + maxPossibleAttributeLength]
	attribute, ok := l.compareWithPattern(possibleAttribute, chiralPattern)
	if ok {
		l.shiftPositionNTimes(len(attribute))
	}
	return startPos, attribute, ok
}

func (l *lexer) checkElementToken() (position, tokenType, string, bool) {
	if isLetter(l.currentSymbol()) {
		startPos := l.currentPosition
		var tType tokenType
		var attribute string
		maxPossibleAttributeLength := maxPossibleElementsLength
		if maxPossibleAttributeLength > len(l.text) - l.currentPosition.index {
			maxPossibleAttributeLength = len(l.text) - l.currentPosition.index
		}
		possibleAttibute := l.text[l.currentPosition.index : l.currentPosition.index + maxPossibleAttributeLength]
		if l.currentSymbol() == 'H' {
			tType = hSymbol
			attribute = string(l.currentSymbol())
		} else if attr, ok := l.compareWithPattern(possibleAttibute, aliphaticOrganicPattern); ok {
			tType = aliphaticOrganic
			attribute = attr
		} else if attr, ok := l.compareWithPattern(possibleAttibute, elementsSymbolPattern); ok {
			tType = elementSymbol
			attribute = attr
		} else if attr, ok := l.compareWithPattern(possibleAttibute, aromaticOrganicPattern); ok {
			tType = aromaticOrganic
			attribute = attr
		} else if attr, ok := l.compareWithPattern(possibleAttibute, aromaticSymbolPattern); ok {
			tType = aromaticSymbol
			attribute = attr
		} else {
			return position{}, -1, "", false
		}
		l.shiftPositionNTimes(len(attribute))
		return startPos, tType, attribute, true
	}
	return position{}, -1, "", false
}

func (l *lexer) compareWithPattern(str string, pattern *regexp.Regexp) (string, bool) {
	location := pattern.FindIndex([]byte(str))
	if location != nil && location[0] == 0 {
		attribute := l.text[l.currentPosition.index : l.currentPosition.index + location[1]]
		return attribute, true
	}
	return "", false
}

func isDigit(symbol byte) bool {
	return symbol >= '0' && symbol <= '9'
}

func isLetter(symbol byte) bool {
	return symbol >= 'a' && symbol <= 'z' || symbol >= 'A' && symbol <= 'Z'
}

func (l *lexer) shiftPositionNTimes(n int) {
	for i := 0; i < n; i++ {
		l.shiftPosition()
	}
}

func (l *lexer) shiftPosition() {
	if l.currentSymbol() == '\n' {
		l.currentPosition.shiftNextLine()
	} else if l.currentSymbol() == '\r' {
		if l.currentPosition.index+1 < len(l.text) && l.text[l.currentPosition.index+1] == '\n' {
			l.currentPosition.index++
		}
		l.currentPosition.shiftNextLine()
	} else {
		l.currentPosition.shiftRight()
	}
	l.currentPosition.index++
}