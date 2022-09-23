package calculator_test

import (
	"calculator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var testValue1, _ = decimal.NewFromString("1")
var testValue2, _ = decimal.NewFromString("2")
var testValue3, _ = decimal.NewFromString("3")
var testValue4, _ = decimal.NewFromString("4")
var testValue7, _ = decimal.NewFromString("7")
var testValueM7, _ = decimal.NewFromString("-7")

var _ = Describe("RPN Stack tests", func() {
	Context("Empty stack", func() {
		stack := calculator.NewStack()

		It("Should have zero in X,Y registers", func() {
			checkStackValues(stack, []decimal.Decimal{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero})

		})
		It("Should update X register", func() {

			stack.WriteX(testValue1)
			checkStackValues(stack, []decimal.Decimal{testValue1, decimal.Zero, decimal.Zero, decimal.Zero})
		})
		It("Should duplicate X to Y on Enter", func() {
			stack.Enter()
			checkStackValues(stack, []decimal.Decimal{testValue1, testValue1, decimal.Zero, decimal.Zero})
		})
		It("Should push values up stack on Enter", func() {
			stack.WriteX(testValue2).Enter().WriteX(testValue3)
			checkStackValues(stack, []decimal.Decimal{testValue3, testValue2, testValue1, decimal.Zero})
			stack.Enter()
			checkStackValues(stack, []decimal.Decimal{testValue3, testValue3, testValue2, testValue1})
		})
		It("Should drop values from stack on Drop()", func() {
			stack.Drop()
			checkStackValues(stack, []decimal.Decimal{testValue3, testValue2, testValue1, decimal.Zero})
		})

	})
	Context("Populated stack", func() {
		stack := calculator.NewStack()
		stack.WriteX(testValue1).Enter().WriteX(testValue2).Enter().WriteX(testValue3).Enter().WriteX(testValue4)
		It("Should add two numbers", func() {
			stack.Print()
			stack.BinaryOp(calculator.Add)
			checkStackValues(stack, []decimal.Decimal{testValue7, testValue2, testValue1, decimal.Zero})
		})
		It("Unary operation should change value of x-register", func() {
			stack.UnaryOp(calculator.ChangeSign)
			checkStackValues(stack, []decimal.Decimal{testValueM7, testValue2, testValue1, decimal.Zero})
		})
	})
})

func checkStackValues(stack *calculator.RpnStack, values []decimal.Decimal) {
	Expect(stack.X()).To(Equal(values[0]))
	Expect(stack.Y()).To(Equal(values[1]))
	Expect(stack.Z()).To(Equal(values[2]))
	Expect(stack.T()).To(Equal(values[3]))
}
