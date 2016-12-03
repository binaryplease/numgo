//+build !amd64 noasm appengine

package numgo

var (
	Sse3Supt, AvxSupt, Avx2Supt, FmaSupt bool
)

func initasm() {
}

func AddC(c numgo.nDimElement, d []numgo.nDimElement) {
	for i := range d {
		d[i] += c
	}
}

func SubtrC(c numgo.nDimElement, d []numgo.nDimElement) {
	for i := range d {
		d[i] -= c
	}
}

func MultC(c numgo.nDimElement, d []numgo.nDimElement) {
	for i := range d {
		d[i] *= c
	}
}

func DivC(c numgo.nDimElement, d []numgo.nDimElement) {
	for i := range d {
		d[i] /= c
	}
}

func Add(a, b []numgo.nDimElement) {
	lna, lnb := len(a), len(b)
	for i, j := 0, 0; i < lna; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		a[i] += b[j]
	}
}

func Vadd(a, b []numgo.nDimElement) {
	for i := range a {
		a[i] += b[i]
	}
}

func Hadd(st uint64, a []numgo.nDimElement) {
	ln := uint64(len(a))
	for k := uint64(0); k < ln/st; k++ {
		a[k] = a[k*st]
		for i := uint64(1); i < st; i++ {
			a[k] += a[k*st+i]
		}
	}
}

func Subtr(a, b []numgo.nDimElement) {
	lna, lnb := len(a), len(b)
	for i, j := 0, 0; i < lna; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		a[i] -= b[j]
	}
}

func Mult(a, b []numgo.nDimElement) {
	lna, lnb := len(a), len(b)
	for i, j := 0, 0; i < lna; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		a[i] *= b[j]
	}
}

func Div(a, b []numgo.nDimElement) {
	lna, lnb := len(a), len(b)
	for i, j := 0, 0; i < lna; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		a[i] /= b[j]
	}
}

func Fma12(a numgo.nDimElement, x, b []numgo.nDimElement) {
	lnx, lnb := len(x), len(b)
	for i, j := 0, 0; i < lnx; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		x[i] = a*x[i] + b[j]
	}
}

func Fma21(a numgo.nDimElement, x, b []numgo.nDimElement) {
	lnx, lnb := len(x), len(b)
	for i, j := 0, 0; i < lnx; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		x[i] = x[i]*b[j] + a
	}
}
