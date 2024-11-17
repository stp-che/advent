package solution

import (
	"fmt"
	"log"
)

type interpreter struct {
	ctx map[string]expression
}

func (i *interpreter) evalRoot() int {
	return i.ctx["root"].eval(i.ctx)
}

func (i *interpreter) calcParam(param string) int {
	root := i.ctx["root"].(*binOp)
	f, _ := root.left.revertWithParam(i.ctx, param)
	e := root.right
	if f == nil {
		f, _ = root.right.revertWithParam(i.ctx, param)
		e = root.left
	}

	return f(e.eval(i.ctx))
}

type expression interface {
	eval(map[string]expression) int
	revertWithParam(map[string]expression, string) (func(int) int, error)
}

type num struct {
	val int
}

func (n *num) eval(_ map[string]expression) int {
	return n.val
}

func (n *num) revertWithParam(_ map[string]expression, p string) (func(int) int, error) {
	return nil, paramNotFoundError(n, p)
}

type identifier struct {
	name string
}

func (i *identifier) eval(ctx map[string]expression) int {
	return ctx[i.name].eval(ctx)
}

func (i *identifier) revertWithParam(ctx map[string]expression, p string) (f func(int) int, err error) {
	if i.name == p {
		return func(x int) int { return x }, nil
	}

	return ctx[i.name].revertWithParam(ctx, p)
}

const (
	oPlus  = byte('+')
	oMinus = byte('-')
	oMult  = byte('*')
	oDiv   = byte('/')
)

type binOp struct {
	kind        byte
	left, right expression
}

func (o *binOp) eval(ctx map[string]expression) int {
	l := o.left.eval(ctx)
	r := o.right.eval(ctx)
	switch o.kind {
	case oPlus:
		return l + r
	case oMinus:
		return l - r
	case oMult:
		return l * r
	case oDiv:
		return l / r
	default:
		log.Fatalf("unknown operation: %v", rune(o.kind))
		return 0
	}
}

func (o *binOp) revertWithParam(ctx map[string]expression, p string) (f func(int) int, err error) {
	lf, _ := o.left.revertWithParam(ctx, p)
	rf, _ := o.right.revertWithParam(ctx, p)

	if lf == nil && rf == nil {
		return nil, paramNotFoundError(o, p)
	}

	switch o.kind {
	case oPlus:
		f, n := rf, o.left
		if rf == nil {
			f, n = lf, o.right
		}

		return func(x int) int { return f(x - n.eval(ctx)) }, nil
	case oMinus:
		if rf == nil {
			return func(x int) int { return lf(x + o.right.eval(ctx)) }, nil
		}

		return func(x int) int { return rf(o.left.eval(ctx) - x) }, nil
	case oMult:
		f, n := rf, o.left
		if rf == nil {
			f, n = lf, o.right
		}

		return func(x int) int { return f(x / n.eval(ctx)) }, nil
	case oDiv:
		if rf == nil {
			return func(x int) int { return lf(x * o.right.eval(ctx)) }, nil
		}

		return func(x int) int { return rf(o.left.eval(ctx) / x) }, nil
	default:
		log.Fatalf("unknown operation: %v", rune(o.kind))
		return func(i int) int { return 0 }, nil
	}
}

func paramNotFoundError(e expression, p string) error {
	return fmt.Errorf("%v does not contain param %s", e, p)
}
