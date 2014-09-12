package main

import (
	"strings"
	"unicode"
)

type Parser struct {
	Position int
	Input    string
}

func (p *Parser) nextChar() byte {
	return p.Input[p.Position]
}

func (p *Parser) startsWith(target string) bool {
	return strings.HasPrefix(p.Input[p.Position:], target)
}

func (p *Parser) EOF() bool {
	return p.Position >= len(p.Input)
}

func (p *Parser) consumeChar() byte {
	current := p.Input[p.Position]
	p.Position++
	return current
}

func (p *Parser) consumeWhile(test func(byte) bool) string {
	result := ""
	for index := 0; index < len(p.Input) && test(p.nextChar()); {
		result += string(p.consumeChar())
	}
	return result
}

func (p *Parser) consumeWhiteSpace() {
	p.consumeWhile(func(b byte) bool { return b == ' ' || b == '\n' })
}

func isAlphaNumeric(b byte) bool {
	return b == '_' || unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b))
}

func (p *Parser) parseTagName() string {
	return p.consumeWhile(isAlphaNumeric)
}
