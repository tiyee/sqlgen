package v5

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type lex struct {
	input []rune
	size  int
	pos   int
	err   error
	Stmt  CreateTableStmt
	prev  *yySymType
}

func (l *lex) SetStmt(st CreateTableStmt) {
	l.Stmt = st
}
func (l *lex) Reduced(rule, state int, lval *yySymType) bool {

	fmt.Printf("rule: %v; state %v; lval: %v\n", rule, state, lval)

	return false
}
func lower(ch rune) rune { return ('a' - 'A') | ch }
func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}
func isIdentifier(ch rune) bool {
	return isLetter(ch) || isDigit(ch)
}
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *lex) Lex(lval *yySymType) int {

	var ch rune

	for ch = l.read(); isWhitespace(ch); ch = l.read() {

		if ch == rune(0) {
			return 0
		}
	}
	lval.pos = l.pos
	defer func() {
		l.prev = lval
	}()
	fmt.Println(string(ch), ch)
	switch ch {
	case '`':
		return TILDE
	case '=':
		return EQ
	case '(':
		return LPAREN
	case ')':
		return RPAREN
	case ',':
		return COMMA
	case '\'', '"':
		str := l.readStringWithQuote(ch)
		lval.s = str
		return tString
	case '0':
		return 0
	}

	if isLetter(ch) {
		str := l.readString()
		tok := findKeyword(str)
		if tok > 0 {
			return tok
		}
		lval.s = str
		return tString
	}
	if isDigit(ch) {
		str := l.readNumber()
		if n, err := strconv.ParseInt(str, 10, 64); err == nil {
			lval.i = n
		} else {
			l.Error(err.Error())
			return 0
		}

		fmt.Println("ret: ", "tString", tString)
		return tNumber
	}
	return 0

}
func (l *lex) read() rune {
	if l.pos < l.size {
		ru := l.input[l.pos]
		l.pos++
		return ru

	}
	return 0
}

func (l *lex) peek(i int) rune {
	if i < l.pos {
		return l.input[i]
	} else {
		return '0'
	}
}
func (l *lex) readStringWithQuote(quota rune) string {
	start := l.pos - 1
	for pos := l.pos; pos < l.size; pos++ {
		l.pos = pos + 1
		if l.input[pos] == quota {
			break
		}
		if l.input[pos] == '\\' {
			pos++
		}

	}
	if l.pos >= l.size {
		l.pos = l.size - 1
	}
	return string(l.input[start : l.pos+1])

}

func (l *lex) readString() string {
	var str []rune
	l.pos--
	for ch := l.read(); isIdentifier(ch); ch = l.read() {
		str = append(str, ch)

	}
	l.pos--
	return string(str)

}
func (l *lex) readNumber() string {
	var str []rune
	l.pos--
	for ch := l.read(); isDigit(ch); ch = l.read() {
		str = append(str, ch)

	}
	l.pos--
	return string(str)

}
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

// Error satisfies yyLexer.
func (l *lex) Error(s string) {
	l.err = errors.New(s)
}
