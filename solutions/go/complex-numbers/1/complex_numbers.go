package complexnumbers

// Define the Number type here.
type Number struct {
    a float64
    b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
        a: n1.a + n2.a,
        b: n1.b + n2.b,
    }
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
        a: n1.a - n2.a,
        b: n1.b - n2.b,
    }
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
        a: n1.a*n2.a - n1.b*n2.b,
        b: n1.a*n2.b + n1.b*n2.a,
    }
}

func (n Number) Times(factor float64) Number {
	return Number{
        a: n.a*factor,
        b: n.b*factor,
    }
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
        a: (n1.a*n2.a + n1.b*n2.b) / (n2.a*n2.a + n2.b*n2.b),
        b: (n1.a*n2.b - n1.b*n2.a) / -(n2.a*n2.a + n2.b*n2.b),
    }
}

func (n Number) Conjugate() Number {
	return Number{
        a: n.a,
        b: -n.b,
    }
}

func (n Number) Abs() float64 {
    sum := n.a*n.a + n.b*n.b
    z := sum / 2
    for i := 0; i < 100; i++ {
        z = z - (z*z - sum) / (2 * z)
    }
    return z
}

func (n Number) Exp() Number {
    eToA := exp(n.a)
    cosB := cos(n.b)
    sinB := sin(n.b)
    
    return Number{
        a: eToA * cosB,
        b: eToA * sinB,
    }
}

func exp(x float64) float64 {
    result := 1.0
    term := 1.0
    
    for i := 1; i <= 100; i++ {
        term *= x / float64(i)
        result += term
    }
    
    return result
}

func sin(x float64) float64 {
    result := x
    term := x
    
    for i := 1; i <= 50; i++ {
        term *= x * x / float64((2*i)*(2*i+1))
        if i%2 == 1 {
            result -= term
        } else {
            result += term
        }
    }
    
    return result
}

func cos(x float64) float64 {
    result := 1.0
    term := 1.0
    
    for i := 1; i <= 50; i++ {
        term *= x * x / float64((2*i-1)*(2*i))
        if i%2 == 1 {
            result -= term
        } else {
            result += term
        }
    }
    
    return result
}