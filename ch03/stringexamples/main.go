package main

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/jvillasante/goeg/lib/stringutils"
)

func main() {
	stringSplitting()
	stringReplace()
	readingData()
	strconvExamples()
	regexpExamples()
}

func stringSplitting() {
	names := "Niccolò•Noël•Geoffrey•Amélie••Turlough•José"

	fmt.Print("|")
	for _, name := range strings.Split(names, "•") {
		fmt.Printf("%s|", name)
	}
	fmt.Println()

	fmt.Print("|")
	for _, name := range strings.SplitAfter(names, "•") {
		fmt.Printf("%s|", name)
	}
	fmt.Println()

	for _, record := range []string{"László Lajtha*1892*1963",
		"Édouard Lalo\t1823\t1892", "José Ángel Lamas|1775|1814"} {
		fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
			switch char {
			case '\t', '*', '|':
				return true
			}
			return false
		}))
	}
	fmt.Println()
}

func stringReplace() {
	names := " Antônio\tAndré\tFriedrich\t\t\tJean\t\tÉlisabeth\tIsabella \t"
	names = strings.Replace(names, "\t", " ", -1)
	fmt.Printf("|%s|\n", names)

	names = " Antônio\tAndré\tFriedrich\t\t\tJean\t\tÉlisabeth\tIsabella \t"
	fmt.Printf("|%s|\n", stringutils.SimplifyWhitespace(names))

	fmt.Println(strings.Map(func(char rune) rune {
		if char > 127 {
			return '?'
		}
		return char
	}, "Jérôme Österreich"))

	fmt.Println(strings.Map(func(char rune) rune {
		if char > 127 {
			return -1
		}
		return char
	}, "Jérôme Österreich"))

	fmt.Println()
}

func readingData() {
	// In most cases readers operate on files, so here we might imagine that the reader variable was
	// created by calling bufio.NewReader() on the reader returned by an os.Open() call.
	reader := strings.NewReader("Café")

	for {
		char, size, err := reader.ReadRune()
		if err != nil { // might occur if the reader is reading a file
			if err == io.EOF { // finished without incident
				break
			}
			panic(err) // a problem ocurred
		}

		fmt.Printf("%U '%c' %d: % X\n", char, char, size, []byte(string(char)))
	}

	fmt.Println()
}

func strconvExamples() {
	for _, truth := range []string{"1", "t", "TRUE", "false", "F", "0", "5"} {
		if b, err := strconv.ParseBool(truth); err != nil {
			fmt.Printf("\n{%v}", err)
		} else {
			fmt.Print(b, " ")
		}
	}
	fmt.Println()

	x, err := strconv.ParseFloat("-99.7", 64)
	fmt.Printf("%8T %6v %v\n", x, x, err)
	y, err := strconv.ParseInt("71309", 10, 0)
	fmt.Printf("%8T %6v %v\n", y, y, err)
	z, err := strconv.Atoi("71309")
	fmt.Printf("%8T %6v %v\n", z, z, err)

	s := strconv.FormatBool(z > 100)
	fmt.Println(s)
	i, err := strconv.ParseInt("0xDEED", 0, 32)
	fmt.Println(i, err)
	j, err := strconv.ParseInt("0707", 0, 32)
	fmt.Println(j, err)
	k, err := strconv.ParseInt("10111010001", 2, 32)
	fmt.Println(k, err)

	ii := 16769023
	fmt.Println(strconv.Itoa(ii))
	fmt.Println(strconv.FormatInt(int64(ii), 10))
	fmt.Println(strconv.FormatInt(int64(ii), 2))
	fmt.Println(strconv.FormatInt(int64(ii), 16))

	s = "Alle ønsker å være fri."
	quoted := strconv.Quote(s)
	fmt.Println(quoted)
	unquoted, err := strconv.Unquote(quoted)
	fmt.Println(unquoted)

	fmt.Println()
}

func regexpExamples() {
	// names replacement
	names := []string{"László Lajtha", "Édouard Lalo", "José Ángel Lamas", "Julio C. Villasante"}
	nameRx := regexp.MustCompile(`(?P<forenames>\pL+\.?(?:\s+\pL+\.?)*)\s+(?P<surname>\pL+)`)
	fmt.Printf("%q\n", names)
	for i := 0; i < len(names); i++ {
		names[i] = nameRx.ReplaceAllString(names[i], "${surname}, ${forenames}")
	}
	fmt.Printf("%q\n", names)

	// duplicate words
	text := "This test contains duplicate words, 'cause it contains contains following contains"
	wordRx := regexp.MustCompile(`\w+`)
	if matches := wordRx.FindAllString(text, -1); matches != nil {
		previous := ""
		for _, match := range matches {
			if match == previous {
				fmt.Println("Duplicate word:", match)
			}
			previous = match
		}
	}

	// key: value lines in configuration files
	lines := `conf1: value1
conf2:value2
conf3:   value3
    conf4: value4`
	valueForKey := make(map[string]string)
	keyValueRx := regexp.MustCompile(`\s*([[:alpha:]]\w*)\s*:\s*(.+)`)
	if matches := keyValueRx.FindAllStringSubmatch(lines, -1); matches != nil {
		for _, match := range matches {
			valueForKey[match[1]] = strings.TrimRight(match[2], "\t ")
		}
	}
	fmt.Printf("%q\n", valueForKey)

	fmt.Println()
}
