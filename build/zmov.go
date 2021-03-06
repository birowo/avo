// Code generated by command: avogen -output zmov.go mov. DO NOT EDIT.

package build

import (
	"github.com/mmcloughlin/avo/operand"
	"go/types"
)

func (c *Context) mov(a, b operand.Op, an, bn int, t *types.Basic) {
	switch {
	case (t.Info()&types.IsInteger) != 0 && an == 1 && bn == 1:
		c.MOVB(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) == 0 && an == 1 && bn == 4:
		c.MOVBLSX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) != 0 && an == 1 && bn == 4:
		c.MOVBLZX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) == 0 && an == 1 && bn == 8:
		c.MOVBQSX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) != 0 && an == 1 && bn == 8:
		c.MOVBQZX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) == 0 && an == 1 && bn == 2:
		c.MOVBWSX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) != 0 && an == 1 && bn == 2:
		c.MOVBWZX(a, b)
	case (t.Info()&types.IsInteger) != 0 && an == 4 && bn == 4:
		c.MOVL(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) == 0 && an == 4 && bn == 8:
		c.MOVLQSX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) != 0 && an == 4 && bn == 8:
		c.MOVLQZX(a, b)
	case (t.Info()&types.IsInteger) != 0 && an == 16 && bn == 16:
		c.MOVOU(a, b)
	case (t.Info()&types.IsInteger) != 0 && an == 8 && bn == 8:
		c.MOVQ(a, b)
	case (t.Info()&types.IsInteger) != 0 && an == 8 && bn == 16:
		c.MOVQ(a, b)
	case (t.Info()&types.IsInteger) != 0 && an == 16 && bn == 8:
		c.MOVQ(a, b)
	case (t.Info()&types.IsInteger) != 0 && an == 16 && bn == 16:
		c.MOVQ(a, b)
	case (t.Info()&types.IsFloat) != 0 && an == 8 && bn == 16:
		c.MOVSD(a, b)
	case (t.Info()&types.IsFloat) != 0 && an == 16 && bn == 8:
		c.MOVSD(a, b)
	case (t.Info()&types.IsFloat) != 0 && an == 16 && bn == 16:
		c.MOVSD(a, b)
	case (t.Info()&types.IsFloat) != 0 && an == 4 && bn == 16:
		c.MOVSS(a, b)
	case (t.Info()&types.IsFloat) != 0 && an == 16 && bn == 4:
		c.MOVSS(a, b)
	case (t.Info()&types.IsFloat) != 0 && an == 16 && bn == 16:
		c.MOVSS(a, b)
	case (t.Info()&types.IsInteger) != 0 && an == 2 && bn == 2:
		c.MOVW(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) == 0 && an == 2 && bn == 4:
		c.MOVWLSX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) != 0 && an == 2 && bn == 4:
		c.MOVWLZX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) == 0 && an == 2 && bn == 8:
		c.MOVWQSX(a, b)
	case (t.Info()&types.IsInteger) != 0 && (t.Info()&types.IsUnsigned) != 0 && an == 2 && bn == 8:
		c.MOVWQZX(a, b)
	default:
		c.adderrormessage("could not deduce mov instruction")
	}
}
