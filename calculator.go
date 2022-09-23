package calculator

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type RpnStack struct {
	stack []decimal.Decimal
}

func NewStack() *RpnStack {
	stack := new(RpnStack)
	stack.size(4)
	stack.Clear()
	return stack
}

func (r *RpnStack) size(s int) {
	r.stack = make([]decimal.Decimal, s)
}

func (r *RpnStack) Clear() *RpnStack {
	for i, _ := range r.stack {
		r.stack[i] = decimal.Zero
	}
	return r
}

func (r *RpnStack) X() decimal.Decimal {
	return r.stack[0]
}

func (r *RpnStack) Y() decimal.Decimal {
	return r.stack[1]
}

func (r *RpnStack) Z() decimal.Decimal {
	return r.stack[2]
}

func (r *RpnStack) T() decimal.Decimal {
	return r.stack[3]
}

func (r *RpnStack) Enter() *RpnStack {
	r.push()
	return r
}

func (r *RpnStack) WriteX(value decimal.Decimal) *RpnStack {
	r.stack[0] = value
	return r
}

func (r *RpnStack) Drop() *RpnStack {
	for i := 1; i < len(r.stack); i++ {
		j := i - 1
		r.stack[j] = r.stack[i]
	}
	r.stack[len(r.stack)-1] = decimal.Zero
	return r
}

func (r *RpnStack) push() {
	i := len(r.stack) - 1
	for i > 0 {
		j := i - 1
		r.stack[i] = r.stack[j]
		i = j
	}
}

// BinaryOp Perform binary opperation on X and Y registers in form Y op X
func (r *RpnStack) BinaryOp(f1 func(op1, op2 decimal.Decimal) decimal.Decimal) *RpnStack {
	result := f1(r.X(), r.Y())
	r.Drop().WriteX(result)
	return r
}

func Add(op1, op2 decimal.Decimal) decimal.Decimal {
	return op1.Add(op2)
}

func (r *RpnStack) Print() {
	for _, value := range r.stack {
		fmt.Println(value)
	}
}
