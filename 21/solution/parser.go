package solution

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	numRegExp = regexp.MustCompile("^\\d+$")
)

type parser struct {
	ctx map[string]expression
}

func newParser() *parser {
	return &parser{
		ctx: make(map[string]expression),
	}
}

func (p *parser) parseLine(s string) {
	keyAndExpr := strings.Split(s, ": ")
	p.ctx[keyAndExpr[0]] = p.parseExpression(keyAndExpr[1])
}

func (p *parser) parseExpression(s string) expression {
	parts := strings.Split(s, " ")
	if len(parts) == 1 {
		return p.parseLiteral(parts[0])
	}

	return &binOp{
		kind:  parts[1][0],
		left:  p.parseLiteral(parts[0]),
		right: p.parseLiteral(parts[2]),
	}
}

func (p *parser) parseLiteral(s string) expression {
	if numRegExp.MatchString(s) {
		return &num{parseInt(s)}
	}

	return &identifier{s}
}

func parseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}
