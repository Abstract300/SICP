package main

import (
	"fmt"
	"log"

	rna "github.com/Abstract300/SCIP/RNA"
)

func main() {
	RationalMath()
}

// RationalMath calls Rational calculations
func RationalMath() {
	RatOne := rna.Rational{
		Numerator:   1,
		Denominator: 2,
	}
	RatTwo := rna.Rational{
		Numerator:   1,
		Denominator: 2,
	}
	RatThree := rna.Rational{
		Numerator:   2,
		Denominator: 3,
	}
	AbsOne, err := rna.SumRats(RatOne, RatTwo)
	if err != nil {
		log.Print(err)
	}
	AbsTwo, err := rna.SumRats(AbsOne, RatThree)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(AbsOne)
	fmt.Println(AbsTwo)
}
