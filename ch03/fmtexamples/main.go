package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jvillasante/goeg/lib/numbers"
)

func main() {
	booleans()
	integers()
	characters()
	floatingPointNumbers()
	stringExamples()
	slices()
	debugging()
}

func booleans() {
	fmt.Printf("%t %t\n", true, false)
	fmt.Printf("%d %d\n", numbers.IntForBool(true), numbers.IntForBool(false))
	fmt.Println()
}

func integers() {
	fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37, 37, 37, 37, 37)      // binary (base 2)
	fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41) // octal (base 8)

	i := 3931
	fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i) // hexadecimal (base 16)

	i = 569
	fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i, i, i, numbers.Pad(i, 6, '*'))

	fmt.Println()
}

func characters() {
	fmt.Printf("%d %#04x %U '%c'\n", 0x3A6, 934, '\u03A6', '\U000003A6')
	fmt.Println()
}

func floatingPointNumbers() {
	for _, x := range []float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
		fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, numbers.Humanize(x, 20, 5, '*', ','))
	}

	for _, x := range []complex128{2 + 3i, 172.6 - 58.3019i, -.827e2 + 9.04831e-3i} {
		fmt.Printf("|%15s|%9.3f|%.2f|%.1e|\n", fmt.Sprintf("%6.2f%+.3fi", real(x), imag(x)), x, x, x)
	}

	fmt.Println()
}

func stringExamples() {
	slogan := "End Óréttlæti♥"
	fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)

	s := "Dare to be naïve"
	fmt.Printf("|%22s|%-22s|%10s|\n", s, s, s)

	i := strings.Index(s, "n")
	fmt.Printf("|%.10s|%.*s|%-22.10s|%s|\n", s, i, s, s, s)

	fmt.Println()
}

func slices() {
	slogan := "End Óréttlæti♥"
	chars := []rune(slogan)
	bytes := []byte(slogan)

	fmt.Printf("%x\n%#x\n%#X\n", chars, chars, chars)
	fmt.Println()
	fmt.Printf("%s\n%x\n%X\n% X\n%v\n", bytes, bytes, bytes, bytes, bytes)

	fmt.Println()
}

func debugging() {
	p := polar{-83.40, 71.60}

	fmt.Printf("|%T|%v|%#v\n", p, p, p)
	fmt.Printf("|%T|%v|%t\n", false, false, false)
	fmt.Printf("|%T|%v|%d|\n", 7607, 7607, 7607)
	fmt.Printf("|%T|%v|%f|\n", math.E, math.E, math.E)
	fmt.Printf("|%T|%v|%f|\n", 5+7i, 5+7i, 5+7i)

	s := "Relativity"
	fmt.Printf("|%T|\"%v\"|\"%s\"|%q|\n", s, s, s, s)

	s = "Alias↔Synonym"
	chars := []rune(s)
	bytes := []byte(s)
	fmt.Printf("%T: %v\n%T: %v\n", chars, chars, bytes, bytes)

	i := 5
	f := -48.3124
	s = "Tomás Bretón"
	fmt.Printf("|%p → %d|%p → %f|%#p → %s|\n", &i, i, &f, f, &s, s)

	fmt.Println([]float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%#v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%.5f\n", []float64{math.E, math.Pi, math.Phi})

	fmt.Printf("%q\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%#v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%17s\n", []string{"Software patents", "kill", "innovation"})

	fmt.Printf("%v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%#v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%#v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%04b\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})

	fmt.Println()
}

type polar struct {
	x, y float64
}
