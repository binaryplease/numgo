//+build !noasm,!appengine

package asm

var (
	Sse3Supt, AvxSupt, Avx2Supt, FmaSupt bool
)

func init() {
	initasm()
}

func initasm()

func AddC(c numgo.nDimElement, d []numgo.nDimElement)

func SubtrC(c numgo.nDimElement, d []numgo.nDimElement)

func MultC(c numgo.nDimElement, d []numgo.nDimElement)

func DivC(c numgo.nDimElement, d []numgo.nDimElement)

func Add(a, b []numgo.nDimElement)

func Vadd(a, b []numgo.nDimElement)

func Hadd(st uint64, a []numgo.nDimElement)

func Subtr(a, b []numgo.nDimElement)

func Mult(a, b []numgo.nDimElement)

func Div(a, b []numgo.nDimElement)

func Fma12(a numgo.nDimElement, x, b []numgo.nDimElement)

func Fma21(a numgo.nDimElement, x, b []numgo.nDimElement)
