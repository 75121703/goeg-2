package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var britishAmerican = "./../data/americanise/british-american.txt"

func init() {
	dir, _ := filepath.Split(os.Args[0])
	britishAmerican = filepath.Join(dir, britishAmerican)
}

// gets the input and output filenames from the command line, creates corresponding file values,
// and then passes the files to the americanise() function to do the work.
func main() {
	inFilename, outFilename, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inFile, outFile := os.Stdin, os.Stdout
	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = inFile.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}

	if outFilename != "" {
		if outFile, err = os.Create(outFilename); err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = outFile.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}

	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}

// returns two strings and an error value
func filenamesFromCommandLine() (inFilename, outFilename string, err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt", filepath.Base(os.Args[0]))
		return "", "", err
	}

	if len(os.Args) > 1 {
		inFilename = os.Args[1]
		if len(os.Args) > 2 {
			outFilename = os.Args[2]
		}
	}

	if inFilename != "" && inFilename == outFilename {
		log.Fatal("won't overwrite the infile")
	}

	return inFilename, outFilename, nil
}

// buffers the inFile reader and the outFile writer. Then it reads lines from the buffered
// reader and writes each line to the buffered writer, having replaced any British English
// words with their U.S. equivalents.
func americanise(inFile io.Reader, outFile io.Writer) (err error) {
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	var replacer func(string) string
	if replacer, err = makeReplacerFunction(britishAmerican); err != nil {
		return err
	}

	wordRx := regexp.MustCompile("[A-Za-z]+")
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil  // io.EOF isn't really an error
			eof = true // this will end the loop at the next iteration
		} else if err != nil {
			return err // finish immediately for real errors
		}

		line = wordRx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}

// takes the name of a file containing original and replacement strings and returns
// a function that given an original string returns its replacement, along with an error
// value. It expects the file to be a UTF-8 encoded text file with one whitespace-separated
// original and replacement word per line.
func makeReplacerFunction(file string) (func(string) string, error) {
	rawBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	text := string(rawBytes)

	usForBritish := make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}

	return func(word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}
		return word
	}, nil
}
