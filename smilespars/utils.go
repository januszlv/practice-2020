package smilespars

import "regexp"

var (
	aliphaticOrganicString = `Br?|Cl?|N|O|S|P|F|I`
	aromaticOrganicString  = `b|c|n|o|s|p`
	aromaticSymbolString   = `b|c|n|o|p|se?|as`
	elementsSymbolString   =	`H(?:e|f|g|s|o)?|L(?:i|v|a|u|r)|B(?:e|r|a|i|h|k)?|C(?:l|a|r|o|u|d|s|n|e|m|f)?|` +
		`N(?:e|a|i|b|d|p|o)?|Os?|F(?:e|r|l|m)?|M(?:g|n|o|t|d)|A(?:l|r|s|g|u|t|c|m)|S(?:i|c|e|r|n|b|g|m)?|` +
		`P(?:d|t|b|o|r|m|a|u)?|Kr?|T(?:i|c|e|a|l|b|m|h)|V|Z(?:n|r)|G(?:a|e|d)|R(?:b|u|h|e|n|a|f|g)|Yb?|` +
		`I(?:n|r)?|Xe|W|D(?:b|s|y)|E(?:u|r|s)|U`
	chiralString = `@(?:@|TH(?:1|2)|AL(?:1|2)|SP(?:1|2|3)|TB\d\d?|OH\d\d?)?`
	numberString = `\d+`
)

var maxPossibleElementsLength = 2
var maxPossibleChiralLength = 5
var valence map[string][]int

var (
	aliphaticOrganicPattern,
	aromaticOrganicPattern,
	aromaticSymbolPattern,
	elementsSymbolPattern,
	chiralPattern,
	numberPattern *regexp.Regexp
)