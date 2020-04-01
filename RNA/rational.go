package rna

import (
	"fmt"
)

// Rational represends a Rational number.
type Rational struct {
	Numerator   float32
	Denominator float32
}

func (r Rational) String() string {
	return fmt.Sprintf("Rational => {%f, %f}", r.Numerator, r.Denominator)
}

// SumRats is the sum of Rationals.
func SumRats(rat1, rat2 Rational) (Rational, error) {
	if rat1.Denominator == 0 || rat2.Denominator == 0 {
		return Rational{}, fmt.Errorf("%s", "Denominator(s) is = 0")
	}
	Num := ((rat1.Numerator) * (rat2.Denominator)) + ((rat2.Numerator) * (rat1.Denominator))
	Denom := (rat1.Denominator * rat2.Denominator)

	FinalRat := Rational{
		Numerator:   Num,
		Denominator: Denom,
	}
	FinalRat = Gcd(FinalRat)

	return FinalRat, nil
}

// Gcd computes GCD on the Rational number
func Gcd(rat Rational) Rational {
	return rat
}
